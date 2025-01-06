// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	EnvironmentService_ListEnvironment_FullMethodName     = "/dream11.od.environment.v1.EnvironmentService/ListEnvironment"
	EnvironmentService_DescribeEnvironment_FullMethodName = "/dream11.od.environment.v1.EnvironmentService/DescribeEnvironment"
	EnvironmentService_UpdateEnvironment_FullMethodName   = "/dream11.od.environment.v1.EnvironmentService/UpdateEnvironment"
	EnvironmentService_CreateEnvironment_FullMethodName   = "/dream11.od.environment.v1.EnvironmentService/CreateEnvironment"
	EnvironmentService_DeleteEnvironment_FullMethodName   = "/dream11.od.environment.v1.EnvironmentService/DeleteEnvironment"
	EnvironmentService_StatusEnvironment_FullMethodName   = "/dream11.od.environment.v1.EnvironmentService/StatusEnvironment"
)

// EnvironmentServiceClient is the client API for EnvironmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentServiceClient interface {
	ListEnvironment(ctx context.Context, in *ListEnvironmentRequest, opts ...grpc.CallOption) (*ListEnvironmentResponse, error)
	DescribeEnvironment(ctx context.Context, in *DescribeEnvironmentRequest, opts ...grpc.CallOption) (*DescribeEnvironmentResponse, error)
	UpdateEnvironment(ctx context.Context, in *UpdateEnvironmentRequest, opts ...grpc.CallOption) (*UpdateEnvironmentResponse, error)
	CreateEnvironment(ctx context.Context, in *CreateEnvironmentRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CreateEnvironmentResponse], error)
	DeleteEnvironment(ctx context.Context, in *DeleteEnvironmentRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DeleteEnvironmentResponse], error)
	StatusEnvironment(ctx context.Context, in *StatusEnvironmentRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StatusEnvironmentResponse], error)
}

type environmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentServiceClient(cc grpc.ClientConnInterface) EnvironmentServiceClient {
	return &environmentServiceClient{cc}
}

func (c *environmentServiceClient) ListEnvironment(ctx context.Context, in *ListEnvironmentRequest, opts ...grpc.CallOption) (*ListEnvironmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListEnvironmentResponse)
	err := c.cc.Invoke(ctx, EnvironmentService_ListEnvironment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentServiceClient) DescribeEnvironment(ctx context.Context, in *DescribeEnvironmentRequest, opts ...grpc.CallOption) (*DescribeEnvironmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeEnvironmentResponse)
	err := c.cc.Invoke(ctx, EnvironmentService_DescribeEnvironment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentServiceClient) UpdateEnvironment(ctx context.Context, in *UpdateEnvironmentRequest, opts ...grpc.CallOption) (*UpdateEnvironmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEnvironmentResponse)
	err := c.cc.Invoke(ctx, EnvironmentService_UpdateEnvironment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentServiceClient) CreateEnvironment(ctx context.Context, in *CreateEnvironmentRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CreateEnvironmentResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &EnvironmentService_ServiceDesc.Streams[0], EnvironmentService_CreateEnvironment_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CreateEnvironmentRequest, CreateEnvironmentResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EnvironmentService_CreateEnvironmentClient = grpc.ServerStreamingClient[CreateEnvironmentResponse]

func (c *environmentServiceClient) DeleteEnvironment(ctx context.Context, in *DeleteEnvironmentRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DeleteEnvironmentResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &EnvironmentService_ServiceDesc.Streams[1], EnvironmentService_DeleteEnvironment_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DeleteEnvironmentRequest, DeleteEnvironmentResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EnvironmentService_DeleteEnvironmentClient = grpc.ServerStreamingClient[DeleteEnvironmentResponse]

func (c *environmentServiceClient) StatusEnvironment(ctx context.Context, in *StatusEnvironmentRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StatusEnvironmentResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &EnvironmentService_ServiceDesc.Streams[2], EnvironmentService_StatusEnvironment_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StatusEnvironmentRequest, StatusEnvironmentResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EnvironmentService_StatusEnvironmentClient = grpc.ServerStreamingClient[StatusEnvironmentResponse]

// EnvironmentServiceServer is the server API for EnvironmentService service.
// All implementations must embed UnimplementedEnvironmentServiceServer
// for forward compatibility.
type EnvironmentServiceServer interface {
	ListEnvironment(context.Context, *ListEnvironmentRequest) (*ListEnvironmentResponse, error)
	DescribeEnvironment(context.Context, *DescribeEnvironmentRequest) (*DescribeEnvironmentResponse, error)
	UpdateEnvironment(context.Context, *UpdateEnvironmentRequest) (*UpdateEnvironmentResponse, error)
	CreateEnvironment(*CreateEnvironmentRequest, grpc.ServerStreamingServer[CreateEnvironmentResponse]) error
	DeleteEnvironment(*DeleteEnvironmentRequest, grpc.ServerStreamingServer[DeleteEnvironmentResponse]) error
	StatusEnvironment(*StatusEnvironmentRequest, grpc.ServerStreamingServer[StatusEnvironmentResponse]) error
	mustEmbedUnimplementedEnvironmentServiceServer()
}

// UnimplementedEnvironmentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEnvironmentServiceServer struct{}

func (UnimplementedEnvironmentServiceServer) ListEnvironment(context.Context, *ListEnvironmentRequest) (*ListEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnvironment not implemented")
}
func (UnimplementedEnvironmentServiceServer) DescribeEnvironment(context.Context, *DescribeEnvironmentRequest) (*DescribeEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeEnvironment not implemented")
}
func (UnimplementedEnvironmentServiceServer) UpdateEnvironment(context.Context, *UpdateEnvironmentRequest) (*UpdateEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEnvironment not implemented")
}
func (UnimplementedEnvironmentServiceServer) CreateEnvironment(*CreateEnvironmentRequest, grpc.ServerStreamingServer[CreateEnvironmentResponse]) error {
	return status.Errorf(codes.Unimplemented, "method CreateEnvironment not implemented")
}
func (UnimplementedEnvironmentServiceServer) DeleteEnvironment(*DeleteEnvironmentRequest, grpc.ServerStreamingServer[DeleteEnvironmentResponse]) error {
	return status.Errorf(codes.Unimplemented, "method DeleteEnvironment not implemented")
}
func (UnimplementedEnvironmentServiceServer) StatusEnvironment(*StatusEnvironmentRequest, grpc.ServerStreamingServer[StatusEnvironmentResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StatusEnvironment not implemented")
}
func (UnimplementedEnvironmentServiceServer) mustEmbedUnimplementedEnvironmentServiceServer() {}
func (UnimplementedEnvironmentServiceServer) testEmbeddedByValue()                            {}

// UnsafeEnvironmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentServiceServer will
// result in compilation errors.
type UnsafeEnvironmentServiceServer interface {
	mustEmbedUnimplementedEnvironmentServiceServer()
}

func RegisterEnvironmentServiceServer(s grpc.ServiceRegistrar, srv EnvironmentServiceServer) {
	// If the following call pancis, it indicates UnimplementedEnvironmentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
	return srv.(EnvironmentServiceServer).CreateEnvironment(m, &grpc.GenericServerStream[CreateEnvironmentRequest, CreateEnvironmentResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EnvironmentService_CreateEnvironmentServer = grpc.ServerStreamingServer[CreateEnvironmentResponse]

func _EnvironmentService_DeleteEnvironment_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DeleteEnvironmentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EnvironmentServiceServer).DeleteEnvironment(m, &grpc.GenericServerStream[DeleteEnvironmentRequest, DeleteEnvironmentResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EnvironmentService_DeleteEnvironmentServer = grpc.ServerStreamingServer[DeleteEnvironmentResponse]

func _EnvironmentService_StatusEnvironment_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StatusEnvironmentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EnvironmentServiceServer).StatusEnvironment(m, &grpc.GenericServerStream[StatusEnvironmentRequest, StatusEnvironmentResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EnvironmentService_StatusEnvironmentServer = grpc.ServerStreamingServer[StatusEnvironmentResponse]

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
		{
			StreamName:    "StatusEnvironment",
			Handler:       _EnvironmentService_StatusEnvironment_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dream11/od/environment/v1/environment.proto",
}
