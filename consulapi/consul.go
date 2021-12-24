package consulapi

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"golang.org/x/xerrors"
	"log"
	"strings"
)

var DefaultPrefix = "w/"

type Client struct {
	client *api.Client
}

type ConsulVal struct {
	GPUDevicesCount int64 `json:"gpu_devices_count"`
}

func New(registryAddr string) *Client {
	config := api.DefaultConfig()
	config.Address = registryAddr
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{client: client}
}

func (c *Client) put(k string, v []byte) error {
	kv := c.client.KV()
	_, err := kv.Put(&api.KVPair{Key: k, Value: v}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) get(k string) ([]byte, error) {
	kv := c.client.KV()
	pair, _, err := kv.Get(k, nil)
	if err != nil {
		return nil, err
	}
	return pair.Value, nil
}

func withPrefix(k string) string {
	return fmt.Sprintf("%s%s", DefaultPrefix, k)
}

func trimPrefix(k string) string {
	return strings.TrimPrefix(k, DefaultPrefix)
}

func (c *Client) PutGpuDevices(key string, count int64) error {
	val, err := json.Marshal(&ConsulVal{GPUDevicesCount: count})
	if err != nil {
		return err
	}

	if err := c.put(withPrefix(key), val); err != nil {
		return err
	}

	return nil
}

func (c *Client) CASGpuDevices(key string, fn func(v *ConsulVal)) error {
	kv := c.client.KV()
	pair, _, err := kv.Get(withPrefix(key), nil)
	if err != nil {
		return err
	}

	var before ConsulVal
	if err := json.Unmarshal(pair.Value, &before); err != nil {
		return err
	}

	fn(&before)

	bytes, err := json.Marshal(before)
	if err != nil {
		return err
	}

	pair.Value = bytes

	ok, _, err := kv.CAS(pair, nil)
	if err != nil {
		return err
	}

	if !ok {
		return xerrors.New("update kv pair failed")
	}

	return nil
}

func (c *Client) GetGpuDevices(key string) (*ConsulVal, error) {
	bytes, err := c.get(withPrefix(key))
	if err != nil {
		return nil, err
	}

	var val ConsulVal
	if err := json.Unmarshal(bytes, &val); err != nil {
		return nil, err
	}

	return &val, nil
}

func (c *Client) List(prefix string) (map[string]*ConsulVal, error) {
	kv := c.client.KV()
	pairs, _, err := kv.List(prefix, nil)
	if err != nil {
		return nil, err
	}

	out := make(map[string]*ConsulVal)
	for _, pair := range pairs {
		var val ConsulVal
		if err := json.Unmarshal(pair.Value, &val); err != nil {
			return nil, err
		}

		fmt.Println("trimPrefix(pair.Key)", trimPrefix(pair.Key))
		out[trimPrefix(pair.Key)] = &val
	}

	return out, nil
}
