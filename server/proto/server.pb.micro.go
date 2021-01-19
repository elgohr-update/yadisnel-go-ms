// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: server.proto

package go_micro_server

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/yadisnel/go-ms/v2/api"
	client "github.com/yadisnel/go-ms/v2/client"
	server "github.com/yadisnel/go-ms/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Server service

func NewServerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Server service

type ServerService interface {
	Handle(ctx context.Context, in *HandleRequest, opts ...client.CallOption) (*HandleResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...client.CallOption) (*SubscribeResponse, error)
}

type serverService struct {
	c    client.Client
	name string
}

func NewServerService(name string, c client.Client) ServerService {
	return &serverService{
		c:    c,
		name: name,
	}
}

func (c *serverService) Handle(ctx context.Context, in *HandleRequest, opts ...client.CallOption) (*HandleResponse, error) {
	req := c.c.NewRequest(c.name, "Server.Handle", in)
	out := new(HandleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverService) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...client.CallOption) (*SubscribeResponse, error) {
	req := c.c.NewRequest(c.name, "Server.Subscribe", in)
	out := new(SubscribeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Server service

type ServerHandler interface {
	Handle(context.Context, *HandleRequest, *HandleResponse) error
	Subscribe(context.Context, *SubscribeRequest, *SubscribeResponse) error
}

func RegisterServerHandler(s server.Server, hdlr ServerHandler, opts ...server.HandlerOption) error {
	type server interface {
		Handle(ctx context.Context, in *HandleRequest, out *HandleResponse) error
		Subscribe(ctx context.Context, in *SubscribeRequest, out *SubscribeResponse) error
	}
	type Server struct {
		server
	}
	h := &serverHandler{hdlr}
	return s.Handle(s.NewHandler(&Server{h}, opts...))
}

type serverHandler struct {
	ServerHandler
}

func (h *serverHandler) Handle(ctx context.Context, in *HandleRequest, out *HandleResponse) error {
	return h.ServerHandler.Handle(ctx, in, out)
}

func (h *serverHandler) Subscribe(ctx context.Context, in *SubscribeRequest, out *SubscribeResponse) error {
	return h.ServerHandler.Subscribe(ctx, in, out)
}
