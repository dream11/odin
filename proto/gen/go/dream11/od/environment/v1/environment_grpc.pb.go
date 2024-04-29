// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: dream11/od/environment/v1/environment.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EnvironmentService_ListEnvironment_FullMethodName     = "/dream11.od.environment.v1.EnvironmentService/ListEnvironment"
	EnvironmentService_DescribeEnvironment_FullMethodName = "/dream11.od.environment.v1.EnvironmentService/DescribeEnvironment"
	EnvironmentService_UpdateEnvironment_FullMethodName   = "/dream11.od.environment.v1.EnvironmentService/UpdateEnvironment"
	EnvironmentService_CreateEnvironment_FullMethodName   = "/dream11.od.environment.v1.EnvironmentService/CreateEnvironment"
	EnvironmentService_DeleteEnvironment_FullMethodName   = "/dream11.od.environment.v1.EnvironmentService/DeleteEnvironment"
)

// EnvironmentServiceClient is the client API for EnvironmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentServiceClient interface {
	ListEnvironment(ctx context.Context, in *ListEnvironmentRequest, opts ...grpc.CallOption) (*ListEnvironmentResponse, error)
	DescribeEnvironment(ctx context.Context, in *DescribeEnvironmentRequest, opts ...grpc.CallOption) (*DescribeEnvironmentResponse, error)
	UpdateEnvironment(ctx context.Context, in *UpdateEnvironmentRequest, opts ...grpc.CallOption) (*UpdateEnvironmentResponse, error)
	CreateEnvironment(ctx context.Context, in *CreateEnvironmentRequest, opts ...grpc.CallOption) (EnvironmentService_CreateEnvironmentClient, error)
	DeleteEnvironment(ctx context.Context, in *DeleteEnvironmentRequest, opts ...grpc.CallOption) (EnvironmentService_DeleteEnvironmentClient, error)
}

type environmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentServiceClient(cc grpc.ClientConnInterface) EnvironmentServiceClient {
	return &environmentServiceClient{cc}
}

func (c *environmentServiceClient) ListEnvironment(ctx context.Context, in *ListEnvironmentRequest, opts ...grpc.CallOption) (*ListEnvironmentResponse, error) {
	out := new(ListEnvironmentResponse)
	err := c.cc.Invoke(ctx, EnvironmentService_ListEnvironment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentServiceClient) DescribeEnvironment(ctx context.Context, in *DescribeEnvironmentRequest, opts ...grpc.CallOption) (*DescribeEnvironmentResponse, error) {
	out := new(DescribeEnvironmentResponse)
	err := c.cc.Invoke(ctx, EnvironmentService_DescribeEnvironment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentServiceClient) UpdateEnvironment(ctx context.Context, in *UpdateEnvironmentRequest, opts ...grpc.CallOption) (*UpdateEnvironmentResponse, error) {
	out := new(UpdateEnvironmentResponse)
	err := c.cc.Invoke(ctx, EnvironmentService_UpdateEnvironment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentServiceClient) CreateEnvironment(ctx context.Context, in *CreateEnvironmentRequest, opts ...grpc.CallOption) (EnvironmentService_CreateEnvironmentClient, error) {
	stream, err := c.cc.NewStream(ctx, &EnvironmentService_ServiceDesc.Streams[0], EnvironmentService_CreateEnvironment_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &environmentServiceCreateEnvironmentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EnvironmentService_CreateEnvironmentClient interface {
	Recv() (*CreateEnvironmentResponse, error)
	grpc.ClientStream
}

type environmentServiceCreateEnvironmentClient struct {
	grpc.ClientStream
}

func (x *environmentServiceCreateEnvironmentClient) Recv() (*CreateEnvironmentResponse, error) {
	m := new(CreateEnvironmentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *environmentServiceClient) DeleteEnvironment(ctx context.Context, in *DeleteEnvironmentRequest, opts ...grpc.CallOption) (EnvironmentService_DeleteEnvironmentClient, error) {
	stream, err := c.cc.NewStream(ctx, &EnvironmentService_ServiceDesc.Streams[1], EnvironmentService_DeleteEnvironment_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &environmentServiceDeleteEnvironmentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EnvironmentService_DeleteEnvironmentClient interface {
	Recv() (*DeleteEnvironmentResponse, error)
	grpc.ClientStream
}

type environmentServiceDeleteEnvironmentClient struct {
	grpc.ClientStream
}

func (x *environmentServiceDeleteEnvironmentClient) Recv() (*DeleteEnvironmentResponse, error) {
	m := new(DeleteEnvironmentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EnvironmentServiceServer is the server API for EnvironmentService service.
// All implementations must embed UnimplementedEnvironmentServiceServer
// for forward compatibility
type EnvironmentServiceServer interface {
	ListEnvironment(context.Context, *ListEnvironmentRequest) (*ListEnvironmentResponse, error)
	DescribeEnvironment(context.Context, *DescribeEnvironmentRequest) (*DescribeEnvironmentResponse, error)
	UpdateEnvironment(context.Context, *UpdateEnvironmentRequest) (*UpdateEnvironmentResponse, error)
	CreateEnvironment(*CreateEnvironmentRequest, EnvironmentService_CreateEnvironmentServer) error
	DeleteEnvironment(*DeleteEnvironmentRequest, EnvironmentService_DeleteEnvironmentServer) error
	mustEmbedUnimplementedEnvironmentServiceServer()
}

// UnimplementedEnvironmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEnvironmentServiceServer struct {
}

func (UnimplementedEnvironmentServiceServer) ListEnvironment(context.Context, *ListEnvironmentRequest) (*ListEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnvironment not implemented")
}
func (UnimplementedEnvironmentServiceServer) DescribeEnvironment(context.Context, *DescribeEnvironmentRequest) (*DescribeEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeEnvironment not implemented")
}
func (UnimplementedEnvironmentServiceServer) UpdateEnvironment(context.Context, *UpdateEnvironmentRequest) (*UpdateEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEnvironment not implemented")
}
func (UnimplementedEnvironmentServiceServer) CreateEnvironment(*CreateEnvironmentRequest, EnvironmentService_CreateEnvironmentServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateEnvironment not implemented")
}
func (UnimplementedEnvironmentServiceServer) DeleteEnvironment(*DeleteEnvironmentRequest, EnvironmentService_DeleteEnvironmentServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteEnvironment not implemented")
}
func (UnimplementedEnvironmentServiceServer) mustEmbedUnimplementedEnvironmentServiceServer() {}

// UnsafeEnvironmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentServiceServer will
// result in compilation errors.
type UnsafeEnvironmentServiceServer interface {
	mustEmbedUnimplementedEnvironmentServiceServer()
}

func RegisterEnvironmentServiceServer(s grpc.ServiceRegistrar, srv EnvironmentServiceServer) {
	s.RegisterService(&EnvironmentService_ServiceDesc, srv)
}

func _EnvironmentService_ListEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServiceServer).ListEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentService_ListEnvironment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServiceServer).ListEnvironment(ctx, req.(*ListEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentService_DescribeEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServiceServer).DescribeEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentService_DescribeEnvironment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServiceServer).DescribeEnvironment(ctx, req.(*DescribeEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentService_UpdateEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServiceServer).UpdateEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentService_UpdateEnvironment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServiceServer).UpdateEnvironment(ctx, req.(*UpdateEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentService_CreateEnvironment_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateEnvironmentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EnvironmentServiceServer).CreateEnvironment(m, &environmentServiceCreateEnvironmentServer{stream})
}

type EnvironmentService_CreateEnvironmentServer interface {
	Send(*CreateEnvironmentResponse) error
	grpc.ServerStream
}

type environmentServiceCreateEnvironmentServer struct {
	grpc.ServerStream
}

func (x *environmentServiceCreateEnvironmentServer) Send(m *CreateEnvironmentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EnvironmentService_DeleteEnvironment_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DeleteEnvironmentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EnvironmentServiceServer).DeleteEnvironment(m, &environmentServiceDeleteEnvironmentServer{stream})
}

type EnvironmentService_DeleteEnvironmentServer interface {
	Send(*DeleteEnvironmentResponse) error
	grpc.ServerStream
}

type environmentServiceDeleteEnvironmentServer struct {
	grpc.ServerStream
}

func (x *environmentServiceDeleteEnvironmentServer) Send(m *DeleteEnvironmentResponse) error {
	return x.ServerStream.SendMsg(m)
}

// EnvironmentService_ServiceDesc is the grpc.ServiceDesc for EnvironmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvironmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dream11.od.environment.v1.EnvironmentService",
	HandlerType: (*EnvironmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEnvironment",
			Handler:    _EnvironmentService_ListEnvironment_Handler,
		},
		{
			MethodName: "DescribeEnvironment",
			Handler:    _EnvironmentService_DescribeEnvironment_Handler,
		},
		{
			MethodName: "UpdateEnvironment",
			Handler:    _EnvironmentService_UpdateEnvironment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateEnvironment",
			Handler:       _EnvironmentService_CreateEnvironment_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DeleteEnvironment",
			Handler:       _EnvironmentService_DeleteEnvironment_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dream11/od/environment/v1/environment.proto",
}
