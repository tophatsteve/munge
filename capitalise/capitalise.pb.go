// Code generated by protoc-gen-go. DO NOT EDIT.
// source: capitalise.proto

/*
Package capitalise is a generated protocol buffer package.

It is generated from these files:
	capitalise.proto

It has these top-level messages:
	CapitaliseRequest
	CapitaliseResponse
*/
package capitalise

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type CapitaliseRequest struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *CapitaliseRequest) Reset()                    { *m = CapitaliseRequest{} }
func (m *CapitaliseRequest) String() string            { return proto.CompactTextString(m) }
func (*CapitaliseRequest) ProtoMessage()               {}
func (*CapitaliseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CapitaliseRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type CapitaliseResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *CapitaliseResponse) Reset()                    { *m = CapitaliseResponse{} }
func (m *CapitaliseResponse) String() string            { return proto.CompactTextString(m) }
func (*CapitaliseResponse) ProtoMessage()               {}
func (*CapitaliseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CapitaliseResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*CapitaliseRequest)(nil), "capitalise.CapitaliseRequest")
	proto.RegisterType((*CapitaliseResponse)(nil), "capitalise.CapitaliseResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Capitalise service

type CapitaliseClient interface {
	Capitalise(ctx context.Context, in *CapitaliseRequest, opts ...grpc.CallOption) (*CapitaliseResponse, error)
}

type capitaliseClient struct {
	cc *grpc.ClientConn
}

func NewCapitaliseClient(cc *grpc.ClientConn) CapitaliseClient {
	return &capitaliseClient{cc}
}

func (c *capitaliseClient) Capitalise(ctx context.Context, in *CapitaliseRequest, opts ...grpc.CallOption) (*CapitaliseResponse, error) {
	out := new(CapitaliseResponse)
	err := grpc.Invoke(ctx, "/capitalise.Capitalise/capitalise", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Capitalise service

type CapitaliseServer interface {
	Capitalise(context.Context, *CapitaliseRequest) (*CapitaliseResponse, error)
}

func RegisterCapitaliseServer(s *grpc.Server, srv CapitaliseServer) {
	s.RegisterService(&_Capitalise_serviceDesc, srv)
}

func _Capitalise_Capitalise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapitaliseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapitaliseServer).Capitalise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capitalise.Capitalise/Capitalise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapitaliseServer).Capitalise(ctx, req.(*CapitaliseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Capitalise_serviceDesc = grpc.ServiceDesc{
	ServiceName: "capitalise.Capitalise",
	HandlerType: (*CapitaliseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "capitalise",
			Handler:    _Capitalise_Capitalise_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "capitalise.proto",
}

func init() { proto.RegisterFile("capitalise.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4e, 0x2c, 0xc8,
	0x2c, 0x49, 0xcc, 0xc9, 0x2c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0xa9, 0x73, 0x09, 0x3a, 0xc3, 0x79, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x25, 0xa9, 0x15, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92,
	0x06, 0x97, 0x10, 0xb2, 0xc2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x6c, 0x2a, 0x8d, 0x22, 0xb9,
	0xb8, 0x10, 0x2a, 0x85, 0xbc, 0xb9, 0x90, 0xac, 0x13, 0x92, 0xd5, 0x43, 0x72, 0x0d, 0x86, 0xc5,
	0x52, 0x72, 0xb8, 0xa4, 0x21, 0xd6, 0x25, 0xb1, 0x81, 0x3d, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0x4b, 0x75, 0xbb, 0xd4, 0x00, 0x00, 0x00,
}
