// Code generated by protoc-gen-go.
// source: github.com/micro/place-srv/proto/location/location.proto
// DO NOT EDIT!

/*
Package location is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/place-srv/proto/location/location.proto

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
import common "github.com/micro/place-srv/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Place struct {
	Id          string        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Location    *common.Point `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	Iata        string        `protobuf:"bytes,4,opt,name=iata" json:"iata,omitempty"`
	Address     string        `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Phone       string        `protobuf:"bytes,6,opt,name=phone" json:"phone,omitempty"`
	Description string        `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	Tags        []string      `protobuf:"bytes,8,rep,name=tags" json:"tags,omitempty"`
}

func (m *Place) Reset()                    { *m = Place{} }
func (m *Place) String() string            { return proto.CompactTextString(m) }
func (*Place) ProtoMessage()               {}
func (*Place) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Place) GetLocation() *common.Point {
	if m != nil {
		return m.Location
	}
	return nil
}

type CreateRequest struct {
	Place *Place `protobuf:"bytes,1,opt,name=place" json:"place,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateRequest) GetPlace() *Place {
	if m != nil {
		return m.Place
	}
	return nil
}

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ReadRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ReadResponse struct {
	Place *Place `protobuf:"bytes,1,opt,name=place" json:"place,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadResponse) GetPlace() *Place {
	if m != nil {
		return m.Place
	}
	return nil
}

type UpdateRequest struct {
	Place *Place `protobuf:"bytes,1,opt,name=place" json:"place,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateRequest) GetPlace() *Place {
	if m != nil {
		return m.Place
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type SearchRequest struct {
	Query  string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Limit  uint64 `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type SearchResponse struct {
	Places []*Place `protobuf:"bytes,1,rep,name=places" json:"places,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SearchResponse) GetPlaces() []*Place {
	if m != nil {
		return m.Places
	}
	return nil
}

func init() {
	proto.RegisterType((*Place)(nil), "Place")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Location service

type LocationClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type locationClient struct {
	c           client.Client
	serviceName string
}

func NewLocationClient(serviceName string, c client.Client) LocationClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "location"
	}
	return &locationClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *locationClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Location.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Location.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Location.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Location.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Location.Search", in)
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

func RegisterLocationHandler(s server.Server, hdlr LocationHandler) {
	s.Handle(s.NewHandler(&Location{hdlr}))
}

type Location struct {
	LocationHandler
}

func (h *Location) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.LocationHandler.Create(ctx, in, out)
}

func (h *Location) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.LocationHandler.Read(ctx, in, out)
}

func (h *Location) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.LocationHandler.Update(ctx, in, out)
}

func (h *Location) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.LocationHandler.Delete(ctx, in, out)
}

func (h *Location) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.LocationHandler.Search(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0xce, 0xd2, 0x30,
	0x14, 0xc0, 0xbf, 0x7d, 0x8c, 0x7d, 0x70, 0x60, 0xc3, 0xd4, 0x68, 0x2a, 0xde, 0x90, 0x25, 0x18,
	0x12, 0xc3, 0x96, 0xe0, 0x8d, 0x37, 0x5e, 0xe9, 0xa5, 0x17, 0x04, 0xe3, 0x03, 0x94, 0xad, 0x40,
	0x93, 0x6d, 0x9d, 0x6d, 0x31, 0xf1, 0x1d, 0x7c, 0x4d, 0xdf, 0xc3, 0xb3, 0x6e, 0x95, 0x8d, 0xc4,
	0xe8, 0x5d, 0x4f, 0xfb, 0x3b, 0xff, 0x7e, 0x1b, 0xbc, 0x3f, 0x0b, 0x73, 0xb9, 0x1e, 0x93, 0x4c,
	0x96, 0x69, 0x29, 0x32, 0x25, 0xd3, 0xba, 0x60, 0x19, 0xdf, 0x6a, 0xf5, 0x3d, 0xad, 0x95, 0x34,
	0x32, 0x2d, 0x64, 0xc6, 0x8c, 0x90, 0xd5, 0x9f, 0x43, 0x62, 0xef, 0x97, 0xdb, 0x7f, 0x66, 0xe2,
	0x4b, 0xe9, 0xf0, 0xf8, 0xa7, 0x07, 0xe3, 0x7d, 0x03, 0x10, 0x80, 0x47, 0x91, 0x53, 0x6f, 0xe5,
	0x6d, 0xa6, 0x64, 0x0e, 0x7e, 0xc5, 0x4a, 0x4e, 0x1f, 0x6d, 0x44, 0x61, 0xe2, 0x9a, 0xd0, 0x11,
	0xde, 0xcc, 0x76, 0x41, 0xb2, 0x97, 0xa2, 0x32, 0x0d, 0x27, 0x98, 0x61, 0xd4, 0xb7, 0xdc, 0x02,
	0x9e, 0x58, 0x9e, 0x2b, 0xae, 0x35, 0x1d, 0xdb, 0x8b, 0x10, 0xc6, 0xf5, 0x45, 0x56, 0x9c, 0x06,
	0x36, 0x7c, 0x0e, 0xb3, 0x9c, 0xeb, 0x4c, 0x89, 0xda, 0x96, 0x7a, 0x72, 0xad, 0x0c, 0x3b, 0x6b,
	0x3a, 0x59, 0x8d, 0x36, 0xd3, 0xf8, 0x0d, 0x84, 0x1f, 0x15, 0x67, 0x86, 0x1f, 0xf8, 0xb7, 0x2b,
	0xd7, 0x86, 0xbc, 0xc0, 0x12, 0xcd, 0x78, 0x76, 0x30, 0xdb, 0xb8, 0x89, 0xe2, 0x67, 0x10, 0x39,
	0x4e, 0xd7, 0xb2, 0xd2, 0x3c, 0x7e, 0x05, 0xb3, 0x03, 0x67, 0xb9, 0xcb, 0xeb, 0x6d, 0x13, 0xaf,
	0x61, 0xde, 0x3e, 0xb5, 0xe8, 0xdf, 0x6a, 0x62, 0xef, 0xaf, 0x75, 0xfe, 0x5f, 0xbd, 0x1d, 0xd7,
	0xf5, 0x7e, 0x0d, 0xe1, 0x27, 0x5e, 0xf0, 0x5b, 0x66, 0xbf, 0x3b, 0xe2, 0xee, 0xb1, 0xc3, 0x3f,
	0x40, 0xf8, 0x85, 0x33, 0x95, 0x5d, 0x1c, 0x8e, 0x9e, 0xf0, 0xa0, 0x7e, 0x74, 0xf6, 0x31, 0x2c,
	0x44, 0x29, 0x8c, 0xd5, 0xea, 0x93, 0x08, 0x02, 0x79, 0x3a, 0x69, 0x6e, 0xac, 0x55, 0x3f, 0xde,
	0x40, 0xe4, 0xd2, 0xbb, 0x85, 0x5e, 0x42, 0x60, 0x07, 0xd5, 0x58, 0x60, 0x74, 0x9b, 0x74, 0xf7,
	0xcb, 0x83, 0xc9, 0xe7, 0xee, 0xcb, 0x91, 0xb7, 0x10, 0xb4, 0xca, 0x48, 0x94, 0x0c, 0x1c, 0x2f,
	0x17, 0xc9, 0x9d, 0xcb, 0x07, 0xb2, 0x06, 0xbf, 0x51, 0x46, 0xe6, 0x49, 0x4f, 0xea, 0x32, 0x4c,
	0xfa, 0x1e, 0x11, 0xc3, 0x9a, 0xad, 0x0a, 0xac, 0x39, 0x70, 0x87, 0x35, 0xef, 0x1c, 0x59, 0xb8,
	0x15, 0x81, 0xf0, 0x40, 0x17, 0xc2, 0x77, 0x86, 0x2c, 0xdc, 0x2e, 0x89, 0xf0, 0x40, 0x16, 0xc2,
	0xc3, 0xed, 0xe3, 0x87, 0x63, 0x60, 0xff, 0xe5, 0x77, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x20,
	0xac, 0x42, 0x16, 0x36, 0x03, 0x00, 0x00,
}