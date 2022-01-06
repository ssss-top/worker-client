package client

import (
	"context"
	"github.com/filecoin-project/go-state-types/abi"
)

type Client struct {
	opts *Options
	sched *scheduler
}

type Options struct {
	Name string
	RegistryAddr string
}

type Option func(*Options)

func ServiceName(name string) Option {
	return func(opts *Options) {
		opts.Name = name
	}
}

func RegistryAddr(addr string) Option {
	return func(opts *Options) {
		opts.RegistryAddr = addr
	}
}

func New(opts ...Option) *Client {
	options := new(Options)
	for _, opt := range opts {
		opt(options)
	}

	return &Client{
		sched: newScheduler(options),
	}
}


func (c *Client) SealCommit2(ctx context.Context, phase1Output []byte, number abi.SectorNumber, minerId abi.ActorID) ([]byte, error) {
	defer func() {c.sched.free <- struct{}{}}()
	return c.sched.Schedule(ctx, phase1Output, number, minerId)
}