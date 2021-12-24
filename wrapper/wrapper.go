package wrapper

import (
	"context"
	"github.com/ssss-top/worker-client/consulapi"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/util/log"
	"golang.org/x/xerrors"
	"net"
	"time"
)

var (
	DialTimeoutInterval    = 30 * time.Second
	RequestTimeoutInterval = 30 * time.Second
)

type SelectWrapper struct {
	client.Client
	api *consulapi.Client
}

func (m *SelectWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	var finish = func() {}
	opts = append(opts,
		client.WithDialTimeout(DialTimeoutInterval),
		client.WithRequestTimeout(RequestTimeoutInterval),
	)
	nOpts := append(opts, client.WithSelectOption(
		// create a selector strategy
		selector.WithStrategy(func(services []*registry.Service) selector.Next {
			// flatten
			var nodes []*registry.Node
			for _, service := range services {
				nodes = append(nodes, service.Nodes...)
			}

			// create the next func that always returns our node
			return func() (*registry.Node, error) {
				if len(nodes) == 0 {
					return nil, selector.ErrNoneAvailable
				}

				list, err := m.api.List(consulapi.DefaultPrefix)
				if err != nil {
					return nil, xerrors.Errorf("list consul kv: %v", err)
				}

				for _, node := range nodes {
					host, _, err := net.SplitHostPort(node.Address)
					if err != nil {
						continue
					}

					val, found := list[host]
					if !found {
						continue
					}

					if val.GPUDevicesCount <= 0 {
						continue
					}

					if err := m.api.CASGpuDevices(host, func(v *consulapi.ConsulVal) {
						v.GPUDevicesCount--
					}); err != nil {
						return nil, xerrors.Errorf("update gpu devices count: %v", err)
					}

					finish = func() {
						if err := m.api.CASGpuDevices(host, func(v *consulapi.ConsulVal) {
							v.GPUDevicesCount++
						}); err != nil {
							log.Errorf("update gpu devices count: %v", err)
						}
					}

					return node, nil
				}

				return nil, selector.ErrNoneAvailable
			}
		}),
	))

	// now do some call
	if err := m.Client.Call(ctx, req, rsp, nOpts...); err != nil {
		finish()
		return err
	}

	finish()

	return nil
}

func NewSelectWrapper(addr string) client.Wrapper {
	return func(c client.Client) client.Client {
		return &SelectWrapper{
			Client: c,
			api:    consulapi.New(addr),
		}
	}
}
