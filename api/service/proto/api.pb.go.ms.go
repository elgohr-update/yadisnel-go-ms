// Code generated by protoc-gen-go-ms. DO NOT EDIT.
// source: api.proto

package go_ms_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Api service

type ApiService interface {
	Register(ctx context.Context, in *Endpoint, opts ...client.CallOption) (*EmptyResponse, error)
	Deregister(ctx context.Context, in *Endpoint, opts ...client.CallOption) (*EmptyResponse, error)
}

type apiService struct {
	c    client.Client
	name string
}

func NewApiService(name string, c client.Client) ApiService {
	return &apiService{
		c:    c,
		name: name,
	}
}

func (c *apiService) Register(ctx context.Context, in *Endpoint, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "Api.Register", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) Deregister(ctx context.Context, in *Endpoint, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "Api.Deregister", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiHandler interface {
	Register(context.Context, *Endpoint, *EmptyResponse) error
	Deregister(context.Context, *Endpoint, *EmptyResponse) error
}

func RegisterApiHandler(s server.Server, hdlr ApiHandler, opts ...server.HandlerOption) error {
	type api interface {
		Register(ctx context.Context, in *Endpoint, out *EmptyResponse) error
		Deregister(ctx context.Context, in *Endpoint, out *EmptyResponse) error
	}
	type Api struct {
		api
	}
	h := &apiHandler{hdlr}
	return s.Handle(s.NewHandler(&Api{h}, opts...))
}

type apiHandler struct {
	ApiHandler
}

func (h *apiHandler) Register(ctx context.Context, in *Endpoint, out *EmptyResponse) error {
	return h.ApiHandler.Register(ctx, in, out)
}

func (h *apiHandler) Deregister(ctx context.Context, in *Endpoint, out *EmptyResponse) error {
	return h.ApiHandler.Deregister(ctx, in, out)
}