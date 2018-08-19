// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/microhq/place-srv/proto/google/google.proto

/*
Package google is a generated protocol buffer package.

It is generated from these files:
	github.com/microhq/place-srv/proto/google/google.proto

It has these top-level messages:
	Price
	Geometry
	OpeningHours
	Photo
	AltId
	NearBySearchRequest
	NearBySearchResponse
	TextSearchRequest
	TextSearchResponse
	RadarSearchRequest
	RadarSearchResponse
*/
package google

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

// Client API for Google service

type GoogleService interface {
	NearBySearch(ctx context.Context, in *NearBySearchRequest, opts ...client.CallOption) (*NearBySearchResponse, error)
	TextSearch(ctx context.Context, in *TextSearchRequest, opts ...client.CallOption) (*TextSearchResponse, error)
	RadarSearch(ctx context.Context, in *RadarSearchRequest, opts ...client.CallOption) (*RadarSearchResponse, error)
}

type googleService struct {
	c    client.Client
	name string
}

func NewGoogleService(name string, c client.Client) GoogleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "google"
	}
	return &googleService{
		c:    c,
		name: name,
	}
}

func (c *googleService) NearBySearch(ctx context.Context, in *NearBySearchRequest, opts ...client.CallOption) (*NearBySearchResponse, error) {
	req := c.c.NewRequest(c.name, "Google.NearBySearch", in)
	out := new(NearBySearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleService) TextSearch(ctx context.Context, in *TextSearchRequest, opts ...client.CallOption) (*TextSearchResponse, error) {
	req := c.c.NewRequest(c.name, "Google.TextSearch", in)
	out := new(TextSearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleService) RadarSearch(ctx context.Context, in *RadarSearchRequest, opts ...client.CallOption) (*RadarSearchResponse, error) {
	req := c.c.NewRequest(c.name, "Google.RadarSearch", in)
	out := new(RadarSearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Google service

type GoogleHandler interface {
	NearBySearch(context.Context, *NearBySearchRequest, *NearBySearchResponse) error
	TextSearch(context.Context, *TextSearchRequest, *TextSearchResponse) error
	RadarSearch(context.Context, *RadarSearchRequest, *RadarSearchResponse) error
}

func RegisterGoogleHandler(s server.Server, hdlr GoogleHandler, opts ...server.HandlerOption) error {
	type google interface {
		NearBySearch(ctx context.Context, in *NearBySearchRequest, out *NearBySearchResponse) error
		TextSearch(ctx context.Context, in *TextSearchRequest, out *TextSearchResponse) error
		RadarSearch(ctx context.Context, in *RadarSearchRequest, out *RadarSearchResponse) error
	}
	type Google struct {
		google
	}
	h := &googleHandler{hdlr}
	return s.Handle(s.NewHandler(&Google{h}, opts...))
	return s.Handle(s.NewHandler(&Google{h}, opts...))
	return s.Handle(s.NewHandler(&Google{h}, opts...))
}

type googleHandler struct {
	GoogleHandler
}

func (h *googleHandler) NearBySearch(ctx context.Context, in *NearBySearchRequest, out *NearBySearchResponse) error {
	return h.GoogleHandler.NearBySearch(ctx, in, out)
}

func (h *googleHandler) TextSearch(ctx context.Context, in *TextSearchRequest, out *TextSearchResponse) error {
	return h.GoogleHandler.TextSearch(ctx, in, out)
}

func (h *googleHandler) RadarSearch(ctx context.Context, in *RadarSearchRequest, out *RadarSearchResponse) error {
	return h.GoogleHandler.RadarSearch(ctx, in, out)
}
