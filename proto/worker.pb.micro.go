// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: worker.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for WorkerService service

func NewWorkerServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for WorkerService service

type WorkerService interface {
	SealCommit2(ctx context.Context, in *SealCommit2Request, opts ...client.CallOption) (*SealCommit2Response, error)
}

type workerService struct {
	c    client.Client
	name string
}

func NewWorkerService(name string, c client.Client) WorkerService {
	return &workerService{
		c:    c,
		name: name,
	}
}

func (c *workerService) SealCommit2(ctx context.Context, in *SealCommit2Request, opts ...client.CallOption) (*SealCommit2Response, error) {
	req := c.c.NewRequest(c.name, "WorkerService.SealCommit2", in)
	out := new(SealCommit2Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkerService service

type WorkerServiceHandler interface {
	SealCommit2(context.Context, *SealCommit2Request, *SealCommit2Response) error
}

func RegisterWorkerServiceHandler(s server.Server, hdlr WorkerServiceHandler, opts ...server.HandlerOption) error {
	type workerService interface {
		SealCommit2(ctx context.Context, in *SealCommit2Request, out *SealCommit2Response) error
	}
	type WorkerService struct {
		workerService
	}
	h := &workerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WorkerService{h}, opts...))
}

type workerServiceHandler struct {
	WorkerServiceHandler
}

func (h *workerServiceHandler) SealCommit2(ctx context.Context, in *SealCommit2Request, out *SealCommit2Response) error {
	return h.WorkerServiceHandler.SealCommit2(ctx, in, out)
}