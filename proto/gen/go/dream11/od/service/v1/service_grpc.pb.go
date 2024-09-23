// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: dream11/od/service/v1/service.proto

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
	ServiceService_DeployService_FullMethodName         = "/dream11.od.service.v1.ServiceService/DeployService"
	ServiceService_ReleaseService_FullMethodName        = "/dream11.od.service.v1.ServiceService/ReleaseService"
	ServiceService_DeployReleasedService_FullMethodName = "/dream11.od.service.v1.ServiceService/DeployReleasedService"
	ServiceService_DeployServiceSet_FullMethodName      = "/dream11.od.service.v1.ServiceService/DeployServiceSet"
	ServiceService_OperateService_FullMethodName        = "/dream11.od.service.v1.ServiceService/OperateService"
	ServiceService_UndeployService_FullMethodName       = "/dream11.od.service.v1.ServiceService/UndeployService"
	ServiceService_ListService_FullMethodName           = "/dream11.od.service.v1.ServiceService/ListService"
	ServiceService_DescribeService_FullMethodName       = "/dream11.od.service.v1.ServiceService/DescribeService"
)

// ServiceServiceClient is the client API for ServiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceServiceClient interface {
	DeployService(ctx context.Context, in *DeployServiceRequest, opts ...grpc.CallOption) (ServiceService_DeployServiceClient, error)
	ReleaseService(ctx context.Context, in *ReleaseServiceRequest, opts ...grpc.CallOption) (ServiceService_ReleaseServiceClient, error)
	DeployReleasedService(ctx context.Context, in *DeployReleasedServiceRequest, opts ...grpc.CallOption) (ServiceService_DeployReleasedServiceClient, error)
	DeployServiceSet(ctx context.Context, in *DeployServiceSetRequest, opts ...grpc.CallOption) (ServiceService_DeployServiceSetClient, error)
	OperateService(ctx context.Context, in *OperateServiceRequest, opts ...grpc.CallOption) (ServiceService_OperateServiceClient, error)
	UndeployService(ctx context.Context, in *UndeployServiceRequest, opts ...grpc.CallOption) (ServiceService_UndeployServiceClient, error)
	ListService(ctx context.Context, in *ListServiceRequest, opts ...grpc.CallOption) (*ListServiceResponse, error)
	DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*DescribeServiceResponse, error)
}

type serviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceServiceClient(cc grpc.ClientConnInterface) ServiceServiceClient {
	return &serviceServiceClient{cc}
}

func (c *serviceServiceClient) DeployService(ctx context.Context, in *DeployServiceRequest, opts ...grpc.CallOption) (ServiceService_DeployServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceService_ServiceDesc.Streams[0], ServiceService_DeployService_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceServiceDeployServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceService_DeployServiceClient interface {
	Recv() (*DeployServiceResponse, error)
	grpc.ClientStream
}

type serviceServiceDeployServiceClient struct {
	grpc.ClientStream
}

func (x *serviceServiceDeployServiceClient) Recv() (*DeployServiceResponse, error) {
	m := new(DeployServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceServiceClient) ReleaseService(ctx context.Context, in *ReleaseServiceRequest, opts ...grpc.CallOption) (ServiceService_ReleaseServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceService_ServiceDesc.Streams[1], ServiceService_ReleaseService_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceServiceReleaseServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceService_ReleaseServiceClient interface {
	Recv() (*ReleaseServiceResponse, error)
	grpc.ClientStream
}

type serviceServiceReleaseServiceClient struct {
	grpc.ClientStream
}

func (x *serviceServiceReleaseServiceClient) Recv() (*ReleaseServiceResponse, error) {
	m := new(ReleaseServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceServiceClient) DeployReleasedService(ctx context.Context, in *DeployReleasedServiceRequest, opts ...grpc.CallOption) (ServiceService_DeployReleasedServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceService_ServiceDesc.Streams[2], ServiceService_DeployReleasedService_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceServiceDeployReleasedServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceService_DeployReleasedServiceClient interface {
	Recv() (*DeployReleasedServiceResponse, error)
	grpc.ClientStream
}

type serviceServiceDeployReleasedServiceClient struct {
	grpc.ClientStream
}

func (x *serviceServiceDeployReleasedServiceClient) Recv() (*DeployReleasedServiceResponse, error) {
	m := new(DeployReleasedServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceServiceClient) DeployServiceSet(ctx context.Context, in *DeployServiceSetRequest, opts ...grpc.CallOption) (ServiceService_DeployServiceSetClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceService_ServiceDesc.Streams[3], ServiceService_DeployServiceSet_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceServiceDeployServiceSetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceService_DeployServiceSetClient interface {
	Recv() (*DeployServiceSetResponse, error)
	grpc.ClientStream
}

type serviceServiceDeployServiceSetClient struct {
	grpc.ClientStream
}

func (x *serviceServiceDeployServiceSetClient) Recv() (*DeployServiceSetResponse, error) {
	m := new(DeployServiceSetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceServiceClient) OperateService(ctx context.Context, in *OperateServiceRequest, opts ...grpc.CallOption) (ServiceService_OperateServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceService_ServiceDesc.Streams[4], ServiceService_OperateService_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceServiceOperateServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceService_OperateServiceClient interface {
	Recv() (*OperateServiceResponse, error)
	grpc.ClientStream
}

type serviceServiceOperateServiceClient struct {
	grpc.ClientStream
}

func (x *serviceServiceOperateServiceClient) Recv() (*OperateServiceResponse, error) {
	m := new(OperateServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceServiceClient) UndeployService(ctx context.Context, in *UndeployServiceRequest, opts ...grpc.CallOption) (ServiceService_UndeployServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceService_ServiceDesc.Streams[5], ServiceService_UndeployService_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceServiceUndeployServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceService_UndeployServiceClient interface {
	Recv() (*UndeployServiceResponse, error)
	grpc.ClientStream
}

type serviceServiceUndeployServiceClient struct {
	grpc.ClientStream
}

func (x *serviceServiceUndeployServiceClient) Recv() (*UndeployServiceResponse, error) {
	m := new(UndeployServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceServiceClient) ListService(ctx context.Context, in *ListServiceRequest, opts ...grpc.CallOption) (*ListServiceResponse, error) {
	out := new(ListServiceResponse)
	err := c.cc.Invoke(ctx, ServiceService_ListService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*DescribeServiceResponse, error) {
	out := new(DescribeServiceResponse)
	err := c.cc.Invoke(ctx, ServiceService_DescribeService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServiceServer is the server API for ServiceService service.
// All implementations must embed UnimplementedServiceServiceServer
// for forward compatibility
type ServiceServiceServer interface {
	DeployService(*DeployServiceRequest, ServiceService_DeployServiceServer) error
	ReleaseService(*ReleaseServiceRequest, ServiceService_ReleaseServiceServer) error
	DeployReleasedService(*DeployReleasedServiceRequest, ServiceService_DeployReleasedServiceServer) error
	DeployServiceSet(*DeployServiceSetRequest, ServiceService_DeployServiceSetServer) error
	OperateService(*OperateServiceRequest, ServiceService_OperateServiceServer) error
	UndeployService(*UndeployServiceRequest, ServiceService_UndeployServiceServer) error
	ListService(context.Context, *ListServiceRequest) (*ListServiceResponse, error)
	DescribeService(context.Context, *DescribeServiceRequest) (*DescribeServiceResponse, error)
	mustEmbedUnimplementedServiceServiceServer()
}

// UnimplementedServiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServiceServer struct {
}

func (UnimplementedServiceServiceServer) DeployService(*DeployServiceRequest, ServiceService_DeployServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method DeployService not implemented")
}
func (UnimplementedServiceServiceServer) ReleaseService(*ReleaseServiceRequest, ServiceService_ReleaseServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method ReleaseService not implemented")
}
func (UnimplementedServiceServiceServer) DeployReleasedService(*DeployReleasedServiceRequest, ServiceService_DeployReleasedServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method DeployReleasedService not implemented")
}
func (UnimplementedServiceServiceServer) DeployServiceSet(*DeployServiceSetRequest, ServiceService_DeployServiceSetServer) error {
	return status.Errorf(codes.Unimplemented, "method DeployServiceSet not implemented")
}
func (UnimplementedServiceServiceServer) OperateService(*OperateServiceRequest, ServiceService_OperateServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method OperateService not implemented")
}
func (UnimplementedServiceServiceServer) UndeployService(*UndeployServiceRequest, ServiceService_UndeployServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method UndeployService not implemented")
}
func (UnimplementedServiceServiceServer) ListService(context.Context, *ListServiceRequest) (*ListServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListService not implemented")
}
func (UnimplementedServiceServiceServer) DescribeService(context.Context, *DescribeServiceRequest) (*DescribeServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeService not implemented")
}
func (UnimplementedServiceServiceServer) mustEmbedUnimplementedServiceServiceServer() {}

// UnsafeServiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServiceServer will
// result in compilation errors.
type UnsafeServiceServiceServer interface {
	mustEmbedUnimplementedServiceServiceServer()
}

func RegisterServiceServiceServer(s grpc.ServiceRegistrar, srv ServiceServiceServer) {
	s.RegisterService(&ServiceService_ServiceDesc, srv)
}

func _ServiceService_DeployService_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DeployServiceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServiceServer).DeployService(m, &serviceServiceDeployServiceServer{stream})
}

type ServiceService_DeployServiceServer interface {
	Send(*DeployServiceResponse) error
	grpc.ServerStream
}

type serviceServiceDeployServiceServer struct {
	grpc.ServerStream
}

func (x *serviceServiceDeployServiceServer) Send(m *DeployServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ServiceService_ReleaseService_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReleaseServiceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServiceServer).ReleaseService(m, &serviceServiceReleaseServiceServer{stream})
}

type ServiceService_ReleaseServiceServer interface {
	Send(*ReleaseServiceResponse) error
	grpc.ServerStream
}

type serviceServiceReleaseServiceServer struct {
	grpc.ServerStream
}

func (x *serviceServiceReleaseServiceServer) Send(m *ReleaseServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ServiceService_DeployReleasedService_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DeployReleasedServiceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServiceServer).DeployReleasedService(m, &serviceServiceDeployReleasedServiceServer{stream})
}

type ServiceService_DeployReleasedServiceServer interface {
	Send(*DeployReleasedServiceResponse) error
	grpc.ServerStream
}

type serviceServiceDeployReleasedServiceServer struct {
	grpc.ServerStream
}

func (x *serviceServiceDeployReleasedServiceServer) Send(m *DeployReleasedServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ServiceService_DeployServiceSet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DeployServiceSetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServiceServer).DeployServiceSet(m, &serviceServiceDeployServiceSetServer{stream})
}

type ServiceService_DeployServiceSetServer interface {
	Send(*DeployServiceSetResponse) error
	grpc.ServerStream
}

type serviceServiceDeployServiceSetServer struct {
	grpc.ServerStream
}

func (x *serviceServiceDeployServiceSetServer) Send(m *DeployServiceSetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ServiceService_OperateService_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OperateServiceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServiceServer).OperateService(m, &serviceServiceOperateServiceServer{stream})
}

type ServiceService_OperateServiceServer interface {
	Send(*OperateServiceResponse) error
	grpc.ServerStream
}

type serviceServiceOperateServiceServer struct {
	grpc.ServerStream
}

func (x *serviceServiceOperateServiceServer) Send(m *OperateServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ServiceService_UndeployService_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UndeployServiceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServiceServer).UndeployService(m, &serviceServiceUndeployServiceServer{stream})
}

type ServiceService_UndeployServiceServer interface {
	Send(*UndeployServiceResponse) error
	grpc.ServerStream
}

type serviceServiceUndeployServiceServer struct {
	grpc.ServerStream
}

func (x *serviceServiceUndeployServiceServer) Send(m *UndeployServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ServiceService_ListService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).ListService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceService_ListService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).ListService(ctx, req.(*ListServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_DescribeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).DescribeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceService_DescribeService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).DescribeService(ctx, req.(*DescribeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceService_ServiceDesc is the grpc.ServiceDesc for ServiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dream11.od.service.v1.ServiceService",
	HandlerType: (*ServiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListService",
			Handler:    _ServiceService_ListService_Handler,
		},
		{
			MethodName: "DescribeService",
			Handler:    _ServiceService_DescribeService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeployService",
			Handler:       _ServiceService_DeployService_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReleaseService",
			Handler:       _ServiceService_ReleaseService_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DeployReleasedService",
			Handler:       _ServiceService_DeployReleasedService_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DeployServiceSet",
			Handler:       _ServiceService_DeployServiceSet_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OperateService",
			Handler:       _ServiceService_OperateService_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UndeployService",
			Handler:       _ServiceService_UndeployService_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dream11/od/service/v1/service.proto",
}
