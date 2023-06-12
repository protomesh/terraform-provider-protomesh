// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/services/v1/resource_store.proto

package servicesv1

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

// ResourceStoreClient is the client API for ResourceStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceStoreClient interface {
	// Put create/update a resource.
	Put(ctx context.Context, in *PutResourceRequest, opts ...grpc.CallOption) (*PutResourceResponse, error)
	// Drop drop multiple resources within a namespace.
	Drop(ctx context.Context, in *DropResourcesRequest, opts ...grpc.CallOption) (*DropResourcesResponse, error)
	// Get return a single resource within a namespace.
	Get(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*GetResourceResponse, error)
	// Watch streams changes in the resources states whithin a namespace.
	Watch(ctx context.Context, in *WatchResourcesRequest, opts ...grpc.CallOption) (ResourceStore_WatchClient, error)
}

type resourceStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceStoreClient(cc grpc.ClientConnInterface) ResourceStoreClient {
	return &resourceStoreClient{cc}
}

func (c *resourceStoreClient) Put(ctx context.Context, in *PutResourceRequest, opts ...grpc.CallOption) (*PutResourceResponse, error) {
	out := new(PutResourceResponse)
	err := c.cc.Invoke(ctx, "/protomesh.services.v1.ResourceStore/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Drop(ctx context.Context, in *DropResourcesRequest, opts ...grpc.CallOption) (*DropResourcesResponse, error) {
	out := new(DropResourcesResponse)
	err := c.cc.Invoke(ctx, "/protomesh.services.v1.ResourceStore/Drop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Get(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*GetResourceResponse, error) {
	out := new(GetResourceResponse)
	err := c.cc.Invoke(ctx, "/protomesh.services.v1.ResourceStore/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Watch(ctx context.Context, in *WatchResourcesRequest, opts ...grpc.CallOption) (ResourceStore_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &ResourceStore_ServiceDesc.Streams[0], "/protomesh.services.v1.ResourceStore/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceStoreWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ResourceStore_WatchClient interface {
	Recv() (*WatchResourcesResponse, error)
	grpc.ClientStream
}

type resourceStoreWatchClient struct {
	grpc.ClientStream
}

func (x *resourceStoreWatchClient) Recv() (*WatchResourcesResponse, error) {
	m := new(WatchResourcesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ResourceStoreServer is the server API for ResourceStore service.
// All implementations must embed UnimplementedResourceStoreServer
// for forward compatibility
type ResourceStoreServer interface {
	// Put create/update a resource.
	Put(context.Context, *PutResourceRequest) (*PutResourceResponse, error)
	// Drop drop multiple resources within a namespace.
	Drop(context.Context, *DropResourcesRequest) (*DropResourcesResponse, error)
	// Get return a single resource within a namespace.
	Get(context.Context, *GetResourceRequest) (*GetResourceResponse, error)
	// Watch streams changes in the resources states whithin a namespace.
	Watch(*WatchResourcesRequest, ResourceStore_WatchServer) error
	mustEmbedUnimplementedResourceStoreServer()
}

// UnimplementedResourceStoreServer must be embedded to have forward compatible implementations.
type UnimplementedResourceStoreServer struct {
}

func (UnimplementedResourceStoreServer) Put(context.Context, *PutResourceRequest) (*PutResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedResourceStoreServer) Drop(context.Context, *DropResourcesRequest) (*DropResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Drop not implemented")
}
func (UnimplementedResourceStoreServer) Get(context.Context, *GetResourceRequest) (*GetResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedResourceStoreServer) Watch(*WatchResourcesRequest, ResourceStore_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedResourceStoreServer) mustEmbedUnimplementedResourceStoreServer() {}

// UnsafeResourceStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceStoreServer will
// result in compilation errors.
type UnsafeResourceStoreServer interface {
	mustEmbedUnimplementedResourceStoreServer()
}

func RegisterResourceStoreServer(s grpc.ServiceRegistrar, srv ResourceStoreServer) {
	s.RegisterService(&ResourceStore_ServiceDesc, srv)
}

func _ResourceStore_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protomesh.services.v1.ResourceStore/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).Put(ctx, req.(*PutResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Drop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).Drop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protomesh.services.v1.ResourceStore/Drop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).Drop(ctx, req.(*DropResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protomesh.services.v1.ResourceStore/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).Get(ctx, req.(*GetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchResourcesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourceStoreServer).Watch(m, &resourceStoreWatchServer{stream})
}

type ResourceStore_WatchServer interface {
	Send(*WatchResourcesResponse) error
	grpc.ServerStream
}

type resourceStoreWatchServer struct {
	grpc.ServerStream
}

func (x *resourceStoreWatchServer) Send(m *WatchResourcesResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ResourceStore_ServiceDesc is the grpc.ServiceDesc for ResourceStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protomesh.services.v1.ResourceStore",
	HandlerType: (*ResourceStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _ResourceStore_Put_Handler,
		},
		{
			MethodName: "Drop",
			Handler:    _ResourceStore_Drop_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ResourceStore_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _ResourceStore_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/services/v1/resource_store.proto",
}
