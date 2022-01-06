package client

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ssss-top/worker-client/proto"
	"github.com/ssss-top/worker-client/wrapper"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/util/log"
	"golang.org/x/xerrors"
	"strings"
	"sync"
	"time"
)

var scheduleInterval = 1 *time.Second

type workerRequest struct {
	ctx context.Context
	ret chan <- workerResponse
	index int

	Phase1Output []byte
	Number       abi.SectorNumber
	MinerId      abi.ActorID
}

type workerResponse struct {
	err error
	proof []byte
}

type scheduler struct {
	lk   	sync.Mutex
	schedule chan *workerRequest
	schedQueue  *requestQueue
	free chan struct{}
	closing chan struct{}
	closed chan struct{}
	service proto.WorkerService
}

func newScheduler(opts *Options) *scheduler {
	service := proto.NewWorkerService(
		opts.Name,
		client.NewClient(
			client.Wrap(wrapper.NewSelectWrapper(opts.RegistryAddr)),
			client.Registry(consul.NewRegistry(registry.Addrs(opts.RegistryAddr))),
		))

	s := &scheduler{
		schedule: make(chan *workerRequest),
		schedQueue: &requestQueue{},
		free :    make(chan struct{}),
		closing:  make(chan struct{}),
		closed:   make(chan struct{}),
		service:  service,
	}

	go s.runSched()

	return s
}

func (s *scheduler) Schedule(ctx context.Context, phase1Output []byte, number abi.SectorNumber, minerId abi.ActorID) ([]byte, error) {
	ret := make(chan workerResponse)

	select {
	case s.schedule <- &workerRequest{
		Phase1Output: phase1Output,
		Number:       number,
		MinerId:      minerId,
		ctx:          ctx,
		ret:          ret,
	}:
	case <-s.closing:
		return nil, xerrors.New("closing")
	case <- ctx.Done():
		return nil, ctx.Err()
	}

	select {
	case resp := <-ret:
		return resp.proof, resp.err
	case <- s.closing:
		return nil, xerrors.Errorf("closing")
	case <- ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *scheduler) runSched() {
	defer close(s.closed)

	for {
		select {
		 	case req := <- s.schedule:
				s.schedQueue.Push(req)
			case <- s.free:
			case <- time.After(scheduleInterval):
			case <- s.closing:
				return
		}

		s.trySched()
	}
}

func (s *scheduler) trySched() {
	queueLen := s.schedQueue.Len()

	log.Debugf("SCHED %d queued", queueLen)

	if queueLen == 0 {
		return
	}

	s.lk.Lock()
	defer s.lk.Unlock()

	for i := 0; i < queueLen; i++ {
		task := (*s.schedQueue)[i]
		s.schedQueue.Remove(i)

		go func(task *workerRequest) {
			var (
				err  error
				resp  []byte
			)

			resp, err = s.call(task)
			if err != nil && strings.Contains(err.Error(), selector.ErrNoneAvailable.Error()) {
				log.Warn(selector.ErrNoneAvailable)
				s.schedQueue.Push(task)
				return
			}

			select {
			case task.ret <- workerResponse{proof: resp, err: err}:
			case <-task.ctx.Done():
				log.Warnf("request got cancelled before we could respond")
			case <-s.closing:
				log.Warnf("scheduler closed while sending response")
			}
		}(task)
	}
}


func (s *scheduler) call(req *workerRequest) ([]byte, error ){
	 resp, err :=  s.service.SealCommit2(req.ctx, &proto.SealCommit2Request{
		Sector:               &proto.SectorID{
			Miner:                uint64(req.MinerId),
			Number:               uint64(req.Number),
		},
		Commit1Out:           req.Phase1Output,
	})


	if err != nil || resp == nil {
		return nil, err
	}

	return resp.Proof, err
}


func (sh *scheduler) Close(ctx context.Context) error {
	close(sh.closing)
	select {
	case <-sh.closed:
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}
