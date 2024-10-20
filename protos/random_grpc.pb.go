// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: protos/random.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Random_GetRandNum_FullMethodName  = "/protoRand.Random/GetRandNum"
	Random_GetRandPass_FullMethodName = "/protoRand.Random/GetRandPass"
)

// RandomClient is the client API for Random service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RandomClient interface {
	GetRandNum(ctx context.Context, in *RandNumRequest, opts ...grpc.CallOption) (*RandNumResponse, error)
	GetRandPass(ctx context.Context, in *RandPassRequest, opts ...grpc.CallOption) (*RandPassResponse, error)
}

type randomClient struct {
	cc grpc.ClientConnInterface
}

func NewRandomClient(cc grpc.ClientConnInterface) RandomClient {
	return &randomClient{cc}
}

func (c *randomClient) GetRandNum(ctx context.Context, in *RandNumRequest, opts ...grpc.CallOption) (*RandNumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RandNumResponse)
	err := c.cc.Invoke(ctx, Random_GetRandNum_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomClient) GetRandPass(ctx context.Context, in *RandPassRequest, opts ...grpc.CallOption) (*RandPassResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RandPassResponse)
	err := c.cc.Invoke(ctx, Random_GetRandPass_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandomServer is the server API for Random service.
// All implementations must embed UnimplementedRandomServer
// for forward compatibility.
type RandomServer interface {
	GetRandNum(context.Context, *RandNumRequest) (*RandNumResponse, error)
	GetRandPass(context.Context, *RandPassRequest) (*RandPassResponse, error)
	mustEmbedUnimplementedRandomServer()
}

// UnimplementedRandomServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRandomServer struct{}

func (UnimplementedRandomServer) GetRandNum(context.Context, *RandNumRequest) (*RandNumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandNum not implemented")
}
func (UnimplementedRandomServer) GetRandPass(context.Context, *RandPassRequest) (*RandPassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandPass not implemented")
}
func (UnimplementedRandomServer) mustEmbedUnimplementedRandomServer() {}
func (UnimplementedRandomServer) testEmbeddedByValue()                {}

// UnsafeRandomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RandomServer will
// result in compilation errors.
type UnsafeRandomServer interface {
	mustEmbedUnimplementedRandomServer()
}

func RegisterRandomServer(s grpc.ServiceRegistrar, srv RandomServer) {
	// If the following call pancis, it indicates UnimplementedRandomServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Random_ServiceDesc, srv)
}

func _Random_GetRandNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandNumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomServer).GetRandNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Random_GetRandNum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomServer).GetRandNum(ctx, req.(*RandNumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Random_GetRandPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomServer).GetRandPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Random_GetRandPass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomServer).GetRandPass(ctx, req.(*RandPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Random_ServiceDesc is the grpc.ServiceDesc for Random service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Random_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoRand.Random",
	HandlerType: (*RandomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRandNum",
			Handler:    _Random_GetRandNum_Handler,
		},
		{
			MethodName: "GetRandPass",
			Handler:    _Random_GetRandPass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/random.proto",
}
