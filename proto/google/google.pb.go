// Code generated by protoc-gen-go.
// source: github.com/micro/place-srv/proto/google/google.proto
// DO NOT EDIT!

/*
Package google is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/place-srv/proto/google/google.proto

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

type Price struct {
	MinPrice uint64 `protobuf:"varint,1,opt,name=min_price" json:"min_price,omitempty"`
	MaxPrice uint64 `protobuf:"varint,2,opt,name=max_price" json:"max_price,omitempty"`
}

func (m *Price) Reset()                    { *m = Price{} }
func (m *Price) String() string            { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()               {}
func (*Price) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Geometry struct {
	Location *common.Point `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
}

func (m *Geometry) Reset()                    { *m = Geometry{} }
func (m *Geometry) String() string            { return proto.CompactTextString(m) }
func (*Geometry) ProtoMessage()               {}
func (*Geometry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Geometry) GetLocation() *common.Point {
	if m != nil {
		return m.Location
	}
	return nil
}

type OpeningHours struct {
	OpenNow bool `protobuf:"varint,1,opt,name=open_now" json:"open_now,omitempty"`
}

func (m *OpeningHours) Reset()                    { *m = OpeningHours{} }
func (m *OpeningHours) String() string            { return proto.CompactTextString(m) }
func (*OpeningHours) ProtoMessage()               {}
func (*OpeningHours) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Photo struct {
	Height           uint64   `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	HtmlAttributions []string `protobuf:"bytes,2,rep,name=html_attributions" json:"html_attributions,omitempty"`
	PhotoReference   string   `protobuf:"bytes,3,opt,name=photo_reference" json:"photo_reference,omitempty"`
	Width            uint64   `protobuf:"varint,4,opt,name=width" json:"width,omitempty"`
}

func (m *Photo) Reset()                    { *m = Photo{} }
func (m *Photo) String() string            { return proto.CompactTextString(m) }
func (*Photo) ProtoMessage()               {}
func (*Photo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type AltId struct {
	PlaceId string `protobuf:"bytes,1,opt,name=place_id" json:"place_id,omitempty"`
	Scope   string `protobuf:"bytes,2,opt,name=scope" json:"scope,omitempty"`
}

func (m *AltId) Reset()                    { *m = AltId{} }
func (m *AltId) String() string            { return proto.CompactTextString(m) }
func (*AltId) ProtoMessage()               {}
func (*AltId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type NearBySearchRequest struct {
	Location      *common.Point `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Radius        uint64        `protobuf:"varint,2,opt,name=radius" json:"radius,omitempty"`
	RankBy        string        `protobuf:"bytes,3,opt,name=rank_by" json:"rank_by,omitempty"`
	Keyword       string        `protobuf:"bytes,4,opt,name=keyword" json:"keyword,omitempty"`
	Language      string        `protobuf:"bytes,5,opt,name=language" json:"language,omitempty"`
	Price         *Price        `protobuf:"bytes,6,opt,name=price" json:"price,omitempty"`
	Name          string        `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	OpenNow       bool          `protobuf:"varint,8,opt,name=open_now" json:"open_now,omitempty"`
	Types         []string      `protobuf:"bytes,9,rep,name=types" json:"types,omitempty"`
	PageToken     string        `protobuf:"bytes,10,opt,name=page_token" json:"page_token,omitempty"`
	ZagatSelected bool          `protobuf:"varint,11,opt,name=zagat_selected" json:"zagat_selected,omitempty"`
}

func (m *NearBySearchRequest) Reset()                    { *m = NearBySearchRequest{} }
func (m *NearBySearchRequest) String() string            { return proto.CompactTextString(m) }
func (*NearBySearchRequest) ProtoMessage()               {}
func (*NearBySearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *NearBySearchRequest) GetLocation() *common.Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *NearBySearchRequest) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

type NearBySearchResponse struct {
	HtmlAttributions []string                                   `protobuf:"bytes,1,rep,name=html_attributions" json:"html_attributions,omitempty"`
	Results          []*NearBySearchResponse_NearBySearchResult `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	Status           string                                     `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *NearBySearchResponse) Reset()                    { *m = NearBySearchResponse{} }
func (m *NearBySearchResponse) String() string            { return proto.CompactTextString(m) }
func (*NearBySearchResponse) ProtoMessage()               {}
func (*NearBySearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *NearBySearchResponse) GetResults() []*NearBySearchResponse_NearBySearchResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type NearBySearchResponse_NearBySearchResult struct {
	Geometry     *Geometry     `protobuf:"bytes,1,opt,name=geometry" json:"geometry,omitempty"`
	Icon         string        `protobuf:"bytes,2,opt,name=icon" json:"icon,omitempty"`
	Id           string        `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Name         string        `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	OpeningHours *OpeningHours `protobuf:"bytes,5,opt,name=opening_hours" json:"opening_hours,omitempty"`
	Photos       []*Photo      `protobuf:"bytes,6,rep,name=photos" json:"photos,omitempty"`
	PlaceId      string        `protobuf:"bytes,7,opt,name=place_id" json:"place_id,omitempty"`
	AltIds       []*AltId      `protobuf:"bytes,8,rep,name=alt_ids" json:"alt_ids,omitempty"`
	Reference    string        `protobuf:"bytes,9,opt,name=reference" json:"reference,omitempty"`
	Types        []string      `protobuf:"bytes,10,rep,name=types" json:"types,omitempty"`
	Vicinity     string        `protobuf:"bytes,11,opt,name=vicinity" json:"vicinity,omitempty"`
}

func (m *NearBySearchResponse_NearBySearchResult) Reset() {
	*m = NearBySearchResponse_NearBySearchResult{}
}
func (m *NearBySearchResponse_NearBySearchResult) String() string { return proto.CompactTextString(m) }
func (*NearBySearchResponse_NearBySearchResult) ProtoMessage()    {}
func (*NearBySearchResponse_NearBySearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

func (m *NearBySearchResponse_NearBySearchResult) GetGeometry() *Geometry {
	if m != nil {
		return m.Geometry
	}
	return nil
}

func (m *NearBySearchResponse_NearBySearchResult) GetOpeningHours() *OpeningHours {
	if m != nil {
		return m.OpeningHours
	}
	return nil
}

func (m *NearBySearchResponse_NearBySearchResult) GetPhotos() []*Photo {
	if m != nil {
		return m.Photos
	}
	return nil
}

func (m *NearBySearchResponse_NearBySearchResult) GetAltIds() []*AltId {
	if m != nil {
		return m.AltIds
	}
	return nil
}

type TextSearchRequest struct {
	Location      *common.Point `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Radius        uint64        `protobuf:"varint,2,opt,name=radius" json:"radius,omitempty"`
	Query         string        `protobuf:"bytes,3,opt,name=query" json:"query,omitempty"`
	Language      string        `protobuf:"bytes,4,opt,name=language" json:"language,omitempty"`
	Price         *Price        `protobuf:"bytes,5,opt,name=price" json:"price,omitempty"`
	Name          string        `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	OpenNow       bool          `protobuf:"varint,7,opt,name=open_now" json:"open_now,omitempty"`
	Types         []string      `protobuf:"bytes,8,rep,name=types" json:"types,omitempty"`
	ZagatSelected bool          `protobuf:"varint,9,opt,name=zagat_selected" json:"zagat_selected,omitempty"`
	PageToken     string        `protobuf:"bytes,10,opt,name=page_token" json:"page_token,omitempty"`
}

func (m *TextSearchRequest) Reset()                    { *m = TextSearchRequest{} }
func (m *TextSearchRequest) String() string            { return proto.CompactTextString(m) }
func (*TextSearchRequest) ProtoMessage()               {}
func (*TextSearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TextSearchRequest) GetLocation() *common.Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *TextSearchRequest) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

type TextSearchResponse struct {
	HtmlAttributions []string                               `protobuf:"bytes,1,rep,name=html_attributions" json:"html_attributions,omitempty"`
	Results          []*TextSearchResponse_TextSearchResult `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	Status           string                                 `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *TextSearchResponse) Reset()                    { *m = TextSearchResponse{} }
func (m *TextSearchResponse) String() string            { return proto.CompactTextString(m) }
func (*TextSearchResponse) ProtoMessage()               {}
func (*TextSearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TextSearchResponse) GetResults() []*TextSearchResponse_TextSearchResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type TextSearchResponse_TextSearchResult struct {
	Geometry         *Geometry     `protobuf:"bytes,1,opt,name=geometry" json:"geometry,omitempty"`
	Icon             string        `protobuf:"bytes,2,opt,name=icon" json:"icon,omitempty"`
	Id               string        `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Name             string        `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	OpeningHours     *OpeningHours `protobuf:"bytes,5,opt,name=opening_hours" json:"opening_hours,omitempty"`
	Photos           []*Photo      `protobuf:"bytes,6,rep,name=photos" json:"photos,omitempty"`
	PlaceId          string        `protobuf:"bytes,7,opt,name=place_id" json:"place_id,omitempty"`
	AltIds           []*AltId      `protobuf:"bytes,8,rep,name=alt_ids" json:"alt_ids,omitempty"`
	Reference        string        `protobuf:"bytes,9,opt,name=reference" json:"reference,omitempty"`
	Types            []string      `protobuf:"bytes,10,rep,name=types" json:"types,omitempty"`
	FormattedAddress string        `protobuf:"bytes,11,opt,name=formatted_address" json:"formatted_address,omitempty"`
}

func (m *TextSearchResponse_TextSearchResult) Reset()         { *m = TextSearchResponse_TextSearchResult{} }
func (m *TextSearchResponse_TextSearchResult) String() string { return proto.CompactTextString(m) }
func (*TextSearchResponse_TextSearchResult) ProtoMessage()    {}
func (*TextSearchResponse_TextSearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

func (m *TextSearchResponse_TextSearchResult) GetGeometry() *Geometry {
	if m != nil {
		return m.Geometry
	}
	return nil
}

func (m *TextSearchResponse_TextSearchResult) GetOpeningHours() *OpeningHours {
	if m != nil {
		return m.OpeningHours
	}
	return nil
}

func (m *TextSearchResponse_TextSearchResult) GetPhotos() []*Photo {
	if m != nil {
		return m.Photos
	}
	return nil
}

func (m *TextSearchResponse_TextSearchResult) GetAltIds() []*AltId {
	if m != nil {
		return m.AltIds
	}
	return nil
}

type RadarSearchRequest struct {
	Location      *common.Point `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Radius        uint64        `protobuf:"varint,2,opt,name=radius" json:"radius,omitempty"`
	Keyword       string        `protobuf:"bytes,3,opt,name=keyword" json:"keyword,omitempty"`
	Price         *Price        `protobuf:"bytes,4,opt,name=price" json:"price,omitempty"`
	Name          string        `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	OpenNow       bool          `protobuf:"varint,6,opt,name=open_now" json:"open_now,omitempty"`
	Types         []string      `protobuf:"bytes,7,rep,name=types" json:"types,omitempty"`
	ZagatSelected bool          `protobuf:"varint,8,opt,name=zagat_selected" json:"zagat_selected,omitempty"`
}

func (m *RadarSearchRequest) Reset()                    { *m = RadarSearchRequest{} }
func (m *RadarSearchRequest) String() string            { return proto.CompactTextString(m) }
func (*RadarSearchRequest) ProtoMessage()               {}
func (*RadarSearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RadarSearchRequest) GetLocation() *common.Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RadarSearchRequest) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

type RadarSearchResponse struct {
	HtmlAttributions []string                                 `protobuf:"bytes,1,rep,name=html_attributions" json:"html_attributions,omitempty"`
	Results          []*RadarSearchResponse_RadarSearchResult `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	Status           string                                   `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *RadarSearchResponse) Reset()                    { *m = RadarSearchResponse{} }
func (m *RadarSearchResponse) String() string            { return proto.CompactTextString(m) }
func (*RadarSearchResponse) ProtoMessage()               {}
func (*RadarSearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RadarSearchResponse) GetResults() []*RadarSearchResponse_RadarSearchResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type RadarSearchResponse_RadarSearchResult struct {
	Geometry *Geometry `protobuf:"bytes,1,opt,name=geometry" json:"geometry,omitempty"`
	PlaceId  string    `protobuf:"bytes,2,opt,name=place_id" json:"place_id,omitempty"`
}

func (m *RadarSearchResponse_RadarSearchResult) Reset()         { *m = RadarSearchResponse_RadarSearchResult{} }
func (m *RadarSearchResponse_RadarSearchResult) String() string { return proto.CompactTextString(m) }
func (*RadarSearchResponse_RadarSearchResult) ProtoMessage()    {}
func (*RadarSearchResponse_RadarSearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

func (m *RadarSearchResponse_RadarSearchResult) GetGeometry() *Geometry {
	if m != nil {
		return m.Geometry
	}
	return nil
}

func init() {
	proto.RegisterType((*Price)(nil), "Price")
	proto.RegisterType((*Geometry)(nil), "Geometry")
	proto.RegisterType((*OpeningHours)(nil), "OpeningHours")
	proto.RegisterType((*Photo)(nil), "Photo")
	proto.RegisterType((*AltId)(nil), "AltId")
	proto.RegisterType((*NearBySearchRequest)(nil), "NearBySearchRequest")
	proto.RegisterType((*NearBySearchResponse)(nil), "NearBySearchResponse")
	proto.RegisterType((*NearBySearchResponse_NearBySearchResult)(nil), "NearBySearchResponse.NearBySearchResult")
	proto.RegisterType((*TextSearchRequest)(nil), "TextSearchRequest")
	proto.RegisterType((*TextSearchResponse)(nil), "TextSearchResponse")
	proto.RegisterType((*TextSearchResponse_TextSearchResult)(nil), "TextSearchResponse.TextSearchResult")
	proto.RegisterType((*RadarSearchRequest)(nil), "RadarSearchRequest")
	proto.RegisterType((*RadarSearchResponse)(nil), "RadarSearchResponse")
	proto.RegisterType((*RadarSearchResponse_RadarSearchResult)(nil), "RadarSearchResponse.RadarSearchResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Google service

type GoogleClient interface {
	NearBySearch(ctx context.Context, in *NearBySearchRequest, opts ...client.CallOption) (*NearBySearchResponse, error)
	TextSearch(ctx context.Context, in *TextSearchRequest, opts ...client.CallOption) (*TextSearchResponse, error)
	RadarSearch(ctx context.Context, in *RadarSearchRequest, opts ...client.CallOption) (*RadarSearchResponse, error)
}

type googleClient struct {
	c           client.Client
	serviceName string
}

func NewGoogleClient(serviceName string, c client.Client) GoogleClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "google"
	}
	return &googleClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *googleClient) NearBySearch(ctx context.Context, in *NearBySearchRequest, opts ...client.CallOption) (*NearBySearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Google.NearBySearch", in)
	out := new(NearBySearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleClient) TextSearch(ctx context.Context, in *TextSearchRequest, opts ...client.CallOption) (*TextSearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Google.TextSearch", in)
	out := new(TextSearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleClient) RadarSearch(ctx context.Context, in *RadarSearchRequest, opts ...client.CallOption) (*RadarSearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Google.RadarSearch", in)
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

func RegisterGoogleHandler(s server.Server, hdlr GoogleHandler) {
	s.Handle(s.NewHandler(&Google{hdlr}))
}

type Google struct {
	GoogleHandler
}

func (h *Google) NearBySearch(ctx context.Context, in *NearBySearchRequest, out *NearBySearchResponse) error {
	return h.GoogleHandler.NearBySearch(ctx, in, out)
}

func (h *Google) TextSearch(ctx context.Context, in *TextSearchRequest, out *TextSearchResponse) error {
	return h.GoogleHandler.TextSearch(ctx, in, out)
}

func (h *Google) RadarSearch(ctx context.Context, in *RadarSearchRequest, out *RadarSearchResponse) error {
	return h.GoogleHandler.RadarSearch(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xdc, 0x95, 0xcd, 0x6e, 0xdb, 0x38,
	0x10, 0x80, 0x23, 0xff, 0xc8, 0xd6, 0x38, 0x7f, 0x56, 0xfe, 0xb4, 0xde, 0x4b, 0x20, 0x04, 0x0b,
	0x5f, 0x22, 0x03, 0xde, 0x5d, 0x2c, 0xb6, 0x40, 0x0f, 0xcd, 0x25, 0xed, 0xa5, 0x0d, 0xd2, 0x9e,
	0x7a, 0x11, 0x68, 0x89, 0x91, 0x88, 0x48, 0xa2, 0x4a, 0xd1, 0x49, 0xdc, 0x77, 0xea, 0x1b, 0xb4,
	0xef, 0xd0, 0x97, 0xe8, 0x23, 0x34, 0xb7, 0x02, 0x1d, 0x51, 0x76, 0x62, 0x59, 0x4a, 0x03, 0xb4,
	0xb7, 0x5e, 0x12, 0x68, 0xcc, 0x19, 0xce, 0x7c, 0xfc, 0x44, 0xc1, 0x3f, 0x01, 0x93, 0xe1, 0x74,
	0xe2, 0x78, 0x3c, 0x1e, 0xc5, 0xcc, 0x13, 0x7c, 0x94, 0x46, 0xc4, 0xa3, 0xc7, 0x99, 0xb8, 0x1a,
	0xa5, 0x82, 0x4b, 0x3e, 0x0a, 0x38, 0x0f, 0x22, 0x3a, 0xff, 0xe7, 0xa8, 0xd8, 0xe0, 0xf8, 0xd1,
	0x2c, 0xfc, 0x25, 0xe6, 0x49, 0xb1, 0xdc, 0x3e, 0x86, 0xf6, 0x99, 0x60, 0x1e, 0x35, 0xfb, 0x60,
	0xc4, 0x2c, 0x71, 0xd3, 0xfc, 0xc1, 0xd2, 0x0e, 0xb5, 0x61, 0x4b, 0x85, 0xc8, 0xcd, 0x3c, 0xd4,
	0xc8, 0x43, 0xf6, 0x11, 0x74, 0x4f, 0x29, 0x8f, 0xa9, 0x14, 0x33, 0xd3, 0x82, 0x6e, 0xc4, 0x3d,
	0x22, 0x19, 0x4f, 0x54, 0x42, 0x6f, 0xac, 0x3b, 0x67, 0x9c, 0x25, 0xd2, 0x3e, 0x84, 0xf5, 0x57,
	0x29, 0x4d, 0x58, 0x12, 0x3c, 0xe7, 0x53, 0x91, 0x99, 0xdb, 0xd0, 0xe5, 0xf8, 0xec, 0x26, 0xfc,
	0x5a, 0xad, 0xec, 0xda, 0x6f, 0x71, 0xdb, 0x10, 0xf7, 0x37, 0x37, 0x41, 0x0f, 0x29, 0x0b, 0x42,
	0x39, 0xdf, 0xf3, 0x0f, 0xe8, 0x87, 0x32, 0x8e, 0x5c, 0x22, 0xa5, 0x60, 0x93, 0x69, 0x5e, 0x3c,
	0xc3, 0xbd, 0x9b, 0x43, 0xc3, 0x3c, 0x80, 0xad, 0x34, 0xcf, 0x71, 0x05, 0xbd, 0xa0, 0x82, 0x26,
	0xd8, 0x54, 0x13, 0x73, 0x0c, 0x73, 0x03, 0xda, 0xd7, 0xcc, 0x97, 0xa1, 0xd5, 0x52, 0x3d, 0x0e,
	0xa1, 0xfd, 0x2c, 0x92, 0x2f, 0xfc, 0x7c, 0x5b, 0x35, 0xbb, 0xcb, 0x7c, 0x55, 0x5d, 0xad, 0xcc,
	0x3c, 0x6c, 0x45, 0x4d, 0x63, 0xd8, 0x5f, 0x34, 0xd8, 0x79, 0x49, 0x89, 0x38, 0x99, 0xbd, 0xc6,
	0xbf, 0x5e, 0x78, 0x4e, 0xdf, 0x4d, 0x69, 0x26, 0x1f, 0x9e, 0x2c, 0x6f, 0x57, 0x10, 0x9f, 0x4d,
	0xb3, 0x82, 0x87, 0xb9, 0x05, 0x1d, 0x41, 0x92, 0x4b, 0x77, 0x32, 0x9b, 0xf7, 0x82, 0x81, 0x4b,
	0x3a, 0xbb, 0xe6, 0xc2, 0x57, 0xdd, 0x18, 0x79, 0x13, 0x11, 0x49, 0x82, 0x29, 0x09, 0xa8, 0xd5,
	0x56, 0x91, 0x3d, 0x68, 0x17, 0x48, 0xf5, 0x45, 0x69, 0x75, 0x00, 0xeb, 0xd0, 0x4a, 0x48, 0x4c,
	0xad, 0xce, 0x22, 0xed, 0x0e, 0x59, 0x37, 0x47, 0x96, 0xf7, 0x2e, 0x67, 0x29, 0xcd, 0x2c, 0x43,
	0xd1, 0x30, 0x01, 0x52, 0xac, 0xe9, 0x4a, 0x7e, 0x49, 0x13, 0x0b, 0x54, 0xd2, 0x3e, 0x6c, 0xbe,
	0x27, 0x01, 0x91, 0x6e, 0x46, 0x23, 0xea, 0x49, 0xea, 0x5b, 0x3d, 0x45, 0xfb, 0x6b, 0x03, 0x76,
	0xcb, 0x73, 0x66, 0x29, 0x72, 0xa5, 0xf5, 0xb4, 0x35, 0x55, 0xff, 0x7f, 0x9c, 0x8c, 0x66, 0xd3,
	0x48, 0x16, 0xf8, 0x7b, 0xe3, 0xa1, 0x53, 0x57, 0x62, 0x35, 0x88, 0x09, 0x39, 0xa4, 0x4c, 0x12,
	0x89, 0x90, 0x14, 0x93, 0xc1, 0xad, 0x06, 0x66, 0xcd, 0xb2, 0x3f, 0xa1, 0x1b, 0xcc, 0x5d, 0x9a,
	0x53, 0x36, 0x9c, 0x3b, 0xb9, 0x90, 0x06, 0xf3, 0x10, 0xbf, 0x3a, 0x28, 0x9c, 0xb5, 0x81, 0x67,
	0x58, 0x10, 0x5e, 0x70, 0x2a, 0xf0, 0x1e, 0xc1, 0x06, 0x2f, 0x54, 0x73, 0xc3, 0xdc, 0x35, 0xc5,
	0xb8, 0x37, 0xde, 0x70, 0x4a, 0x02, 0xee, 0x83, 0xae, 0xd4, 0xc9, 0x90, 0x79, 0xb3, 0x60, 0xae,
	0xec, 0x5b, 0x36, 0xa4, 0xe0, 0x7e, 0x00, 0x1d, 0x12, 0x49, 0x7c, 0xce, 0x10, 0x7b, 0xb1, 0xb4,
	0x90, 0x09, 0x5f, 0x86, 0x7b, 0xef, 0x8c, 0x85, 0x4d, 0xc5, 0x89, 0x80, 0x22, 0x86, 0xc5, 0xae,
	0x98, 0xc7, 0x12, 0x26, 0x67, 0x8a, 0xbb, 0x61, 0x7f, 0xd6, 0xa0, 0xff, 0x86, 0xde, 0xc8, 0x9f,
	0xb5, 0x0b, 0x37, 0xc0, 0x14, 0xb1, 0x70, 0x6b, 0x59, 0xa5, 0x56, 0x59, 0xa5, 0x76, 0xad, 0x4a,
	0x7a, 0x45, 0xa5, 0x4e, 0x59, 0xa5, 0xae, 0x6a, 0xbc, 0xaa, 0x8d, 0xa1, 0x96, 0xd5, 0x28, 0x66,
	0xdf, 0x36, 0xc0, 0x5c, 0x1e, 0xe9, 0x71, 0x91, 0xfe, 0x5d, 0x15, 0xe9, 0xc8, 0xa9, 0x16, 0x28,
	0x87, 0xea, 0x24, 0xfa, 0xa6, 0xc1, 0x76, 0x65, 0xd1, 0xef, 0xa6, 0x10, 0x62, 0xbc, 0xe0, 0x22,
	0x46, 0x8a, 0xd4, 0x77, 0x89, 0xef, 0x23, 0xb8, 0x6c, 0xee, 0xd2, 0x07, 0x7c, 0x89, 0xce, 0x89,
	0x4f, 0xc4, 0x2f, 0x5c, 0x55, 0x8b, 0x9b, 0xa9, 0x59, 0x96, 0xa7, 0x55, 0x2b, 0x4f, 0xbb, 0x22,
	0x8f, 0x5e, 0x96, 0xa7, 0xf3, 0x80, 0x3c, 0xea, 0xba, 0xb2, 0x3f, 0xe1, 0xdd, 0x5a, 0xea, 0xf7,
	0x71, 0x53, 0xfe, 0x5b, 0x35, 0xe5, 0x2f, 0xa7, 0xa6, 0xc2, 0x4a, 0xac, 0xce, 0x95, 0x13, 0xe8,
	0x57, 0x17, 0xfd, 0xd0, 0x95, 0xe5, 0x53, 0x54, 0xbe, 0x8c, 0x3f, 0x6a, 0xa0, 0x9f, 0xaa, 0x0f,
	0xab, 0xf9, 0x14, 0xd6, 0x97, 0xaf, 0x2f, 0x73, 0xd7, 0xa9, 0xf9, 0x68, 0x0c, 0xf6, 0x6a, 0xef,
	0x47, 0x7b, 0x0d, 0xc7, 0x82, 0x7b, 0x71, 0x4d, 0xd3, 0xa9, 0xdc, 0x08, 0x83, 0x9d, 0x9a, 0x37,
	0x02, 0x13, 0x9f, 0x40, 0x6f, 0x69, 0x0c, 0x73, 0xc7, 0xa9, 0x9e, 0xff, 0x60, 0xb7, 0x0e, 0x91,
	0xbd, 0x36, 0xd1, 0xd5, 0xe7, 0xfd, 0xef, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x66, 0xa7, 0xaa,
	0x4b, 0x45, 0x08, 0x00, 0x00,
}
