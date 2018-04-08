// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/place-srv/proto/location/location.proto

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
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

func (m *Place) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Place) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Place) GetLocation() *common.Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Place) GetIata() string {
	if m != nil {
		return m.Iata
	}
	return ""
}

func (m *Place) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Place) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Place) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Place) GetTags() []string {
	if m != nil {
		return m.Tags
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

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

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

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

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

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

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
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Location service

type LocationClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type locationClient struct {
	cc *grpc.ClientConn
}

func NewLocationClient(cc *grpc.ClientConn) LocationClient {
	return &locationClient{cc}
}

func (c *locationClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/Location/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := grpc.Invoke(ctx, "/Location/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/Location/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/Location/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/Location/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Location service

type LocationServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterLocationServer(s *grpc.Server, srv LocationServer) {
	s.RegisterService(&_Location_serviceDesc, srv)
}

func _Location_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Location/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Location_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Location/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Location_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Location/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Location_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Location/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Location_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Location/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Location_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Location",
	HandlerType: (*LocationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Location_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Location_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Location_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Location_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Location_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/micro/place-srv/proto/location/location.proto",
}

func init() {
	proto.RegisterFile("github.com/micro/place-srv/proto/location/location.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcb, 0x6a, 0xdc, 0x30,
	0x14, 0x1d, 0xcf, 0xd8, 0xce, 0xf4, 0x4e, 0xec, 0x09, 0x22, 0x14, 0x61, 0xfa, 0x30, 0x82, 0x42,
	0xa0, 0x1d, 0x4d, 0x49, 0x37, 0xdd, 0xb7, 0xcb, 0x2e, 0x82, 0x43, 0x3f, 0x40, 0xb1, 0x95, 0x8c,
	0x60, 0x6c, 0x39, 0x92, 0xa6, 0xd0, 0x4f, 0xec, 0x87, 0xf4, 0x3f, 0x8a, 0x24, 0x2b, 0xb5, 0xbd,
	0x09, 0xd9, 0xdd, 0x7b, 0x74, 0xce, 0x7d, 0x9c, 0x6b, 0xc3, 0xd7, 0x07, 0x61, 0x0e, 0xa7, 0x3b,
	0x5a, 0xcb, 0x76, 0xdf, 0x8a, 0x5a, 0xc9, 0x7d, 0x7f, 0x64, 0x35, 0xdf, 0x69, 0xf5, 0x6b, 0xdf,
	0x2b, 0x69, 0xe4, 0xfe, 0x28, 0x6b, 0x66, 0x84, 0xec, 0x9e, 0x02, 0xea, 0xf0, 0x62, 0xf7, 0xac,
	0xb2, 0x96, 0x6d, 0x1b, 0xe8, 0xe4, 0x4f, 0x04, 0xc9, 0x8d, 0x25, 0xa0, 0x1c, 0x96, 0xa2, 0xc1,
	0x51, 0x19, 0x5d, 0xbd, 0xaa, 0x96, 0xa2, 0x41, 0x08, 0xe2, 0x8e, 0xb5, 0x1c, 0x2f, 0x1d, 0xe2,
	0x62, 0x44, 0x60, 0x1d, 0xda, 0xe1, 0x55, 0x19, 0x5d, 0x6d, 0xae, 0x53, 0x7a, 0x23, 0x45, 0x67,
	0xaa, 0x27, 0xdc, 0xea, 0x04, 0x33, 0x0c, 0xc7, 0x5e, 0x67, 0x63, 0x84, 0xe1, 0x8c, 0x35, 0x8d,
	0xe2, 0x5a, 0xe3, 0xc4, 0xc1, 0x21, 0x45, 0x97, 0x90, 0xf4, 0x07, 0xd9, 0x71, 0x9c, 0x3a, 0xdc,
	0x27, 0xa8, 0x84, 0x4d, 0xc3, 0x75, 0xad, 0x44, 0xef, 0x5a, 0x9d, 0xb9, 0xb7, 0x31, 0x64, 0xbb,
	0x18, 0xf6, 0xa0, 0xf1, 0xba, 0x5c, 0xd9, 0x2e, 0x36, 0x26, 0x3b, 0xc8, 0xbe, 0x29, 0xce, 0x0c,
	0xaf, 0xf8, 0xe3, 0x89, 0x6b, 0x83, 0xde, 0x40, 0xe2, 0x96, 0x77, 0x5b, 0xb9, 0x59, 0x6d, 0x56,
	0x79, 0x90, 0x5c, 0x40, 0x1e, 0xe8, 0xba, 0x97, 0x9d, 0xe6, 0xe4, 0x2d, 0x6c, 0x2a, 0xce, 0x9a,
	0x20, 0x9f, 0x39, 0x42, 0x3e, 0xc1, 0xb9, 0x7f, 0xf6, 0xf4, 0x67, 0xca, 0xef, 0x20, 0xfb, 0xd9,
	0x37, 0x2f, 0x99, 0x26, 0xd0, 0x87, 0x69, 0xde, 0x43, 0xf6, 0x9d, 0x1f, 0xf9, 0xff, 0x02, 0xf3,
	0x79, 0x2e, 0x20, 0x0f, 0x84, 0x41, 0x72, 0x0b, 0xd9, 0x2d, 0x67, 0xaa, 0x3e, 0x04, 0xc9, 0x25,
	0x24, 0x8f, 0x27, 0xae, 0x7e, 0x0f, 0x2a, 0x9f, 0x58, 0xf4, 0x28, 0x5a, 0x61, 0xdc, 0x8d, 0xe2,
	0xca, 0x27, 0xe8, 0x35, 0xa4, 0xf2, 0xfe, 0x5e, 0x73, 0xe3, 0x6e, 0x14, 0x57, 0x43, 0x46, 0x3e,
	0x43, 0x1e, 0x8a, 0x0e, 0x8b, 0xbf, 0x83, 0xd4, 0x0d, 0xad, 0x71, 0x54, 0xae, 0x46, 0xab, 0x0c,
	0xe8, 0xf5, 0xdf, 0x08, 0xd6, 0x3f, 0xc2, 0xf7, 0xf0, 0x11, 0x52, 0x6f, 0x33, 0xca, 0xe9, 0xe4,
	0x3c, 0xc5, 0x96, 0xce, 0xfc, 0x5f, 0xa0, 0x0f, 0x10, 0x5b, 0x8b, 0xd1, 0x39, 0x1d, 0x1d, 0xa2,
	0xc8, 0xe8, 0xd8, 0x77, 0xb2, 0xb0, 0x35, 0xbd, 0x59, 0x28, 0xa7, 0x13, 0x93, 0x8b, 0x2d, 0x9d,
	0xb9, 0xe8, 0xc8, 0xde, 0x26, 0x94, 0xd3, 0x89, 0xa1, 0xc5, 0x96, 0xce, 0xfc, 0x73, 0x64, 0xbf,
	0x2c, 0xca, 0xe9, 0xc4, 0xca, 0x62, 0x4b, 0xa7, 0x2e, 0x90, 0xc5, 0x5d, 0xea, 0xfe, 0xa1, 0x2f,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x55, 0x58, 0x0d, 0x47, 0xae, 0x03, 0x00, 0x00,
}
