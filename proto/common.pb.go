// Code generated by protoc-gen-go.
// source: github.com/micro/place-srv/proto/common.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/place-srv/proto/common.proto

It has these top-level messages:
	Point
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Point struct {
	Lat float64 `protobuf:"fixed64,1,opt,name=lat" json:"lat,omitempty"`
	Lng float64 `protobuf:"fixed64,2,opt,name=lng" json:"lng,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Point)(nil), "Point")
}

var fileDescriptor0 = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x2f, 0xc8, 0x49,
	0x4c, 0x4e, 0xd5, 0x2d, 0x2e, 0x2a, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x07, 0xca, 0xe4,
	0xe6, 0xe7, 0xe9, 0x81, 0x39, 0x4a, 0x8a, 0x5c, 0xac, 0x01, 0xf9, 0x99, 0x79, 0x25, 0x42, 0xdc,
	0x5c, 0xcc, 0x39, 0x89, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x8c, 0x60, 0x4e, 0x5e, 0xba, 0x04,
	0x13, 0x88, 0x93, 0xc4, 0x06, 0x56, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x53, 0x05, 0x8e,
	0x12, 0x5a, 0x00, 0x00, 0x00,
}
