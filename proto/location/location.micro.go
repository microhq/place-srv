// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/microhq/place-srv/proto/location/location.proto

/*
Package location is a generated protocol buffer package.

It is generated from these files:
	github.com/microhq/place-srv/proto/location/location.proto

It has these top-level messages:
	Place
	CreateRequest
	CreateResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	SearchRequest
	SearchResponse
*/
package location

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/microhq/place-srv/proto"

import (
	context "context"
	api "github.com/micro/go-api"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Location service

type LocationService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type locationService struct {
	c    client.Client
	name string
}

func NewLocationService(name string, c client.Client) LocationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "location"
	}
	return &locationService{
		c:    c,
		name: name,
	}
}

func (c *locationService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Location.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Location.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Location.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Location.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.name, "Location.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Location service

type LocationHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
}

func RegisterLocationHandler(s server.Server, hdlr LocationHandler, opts ...server.HandlerOption) error {
	type location interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error
	}
	type Location struct {
		location
	}
	h := &locationHandler{hdlr}
	return s.Handle(s.NewHandler(&Location{h}, opts...))
	return s.Handle(s.NewHandler(&Location{h}, opts...))
	return s.Handle(s.NewHandler(&Location{h}, opts...))
	return s.Handle(s.NewHandler(&Location{h}, opts...))
	return s.Handle(s.NewHandler(&Location{h}, opts...))
}

type locationHandler struct {
	LocationHandler
}

func (h *locationHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.LocationHandler.Create(ctx, in, out)
}

func (h *locationHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.LocationHandler.Read(ctx, in, out)
}

func (h *locationHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.LocationHandler.Update(ctx, in, out)
}

func (h *locationHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.LocationHandler.Delete(ctx, in, out)
}

func (h *locationHandler) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.LocationHandler.Search(ctx, in, out)
}
