// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbgo

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

// ClusterLcmServiceClient is the client API for ClusterLcmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterLcmServiceClient interface {
	// CreateCluster creates a Kubernetes cluster and returns cluster id
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*IDsResponse, error)
	// ScaleCluster scales the Kubernetes cluster
	ScaleCluster(ctx context.Context, in *ScaleClusterRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// InstallAppGroups install app groups, return a array of app group id
	InstallAppGroups(ctx context.Context, in *InstallAppGroupsRequest, opts ...grpc.CallOption) (*IDsResponse, error)
	// UninstallAppGroups uninstall app groups.
	UninstallAppGroups(ctx context.Context, in *UninstallAppGroupsRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
}

type clusterLcmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterLcmServiceClient(cc grpc.ClientConnInterface) ClusterLcmServiceClient {
	return &clusterLcmServiceClient{cc}
}

func (c *clusterLcmServiceClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*IDsResponse, error) {
	out := new(IDsResponse)
	err := c.cc.Invoke(ctx, "/pbgo.ClusterLcmService/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterLcmServiceClient) ScaleCluster(ctx context.Context, in *ScaleClusterRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/pbgo.ClusterLcmService/ScaleCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterLcmServiceClient) InstallAppGroups(ctx context.Context, in *InstallAppGroupsRequest, opts ...grpc.CallOption) (*IDsResponse, error) {
	out := new(IDsResponse)
	err := c.cc.Invoke(ctx, "/pbgo.ClusterLcmService/InstallAppGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterLcmServiceClient) UninstallAppGroups(ctx context.Context, in *UninstallAppGroupsRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/pbgo.ClusterLcmService/UninstallAppGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterLcmServiceServer is the server API for ClusterLcmService service.
// All implementations must embed UnimplementedClusterLcmServiceServer
// for forward compatibility
type ClusterLcmServiceServer interface {
	// CreateCluster creates a Kubernetes cluster and returns cluster id
	CreateCluster(context.Context, *CreateClusterRequest) (*IDsResponse, error)
	// ScaleCluster scales the Kubernetes cluster
	ScaleCluster(context.Context, *ScaleClusterRequest) (*SimpleResponse, error)
	// InstallAppGroups install app groups, return a array of app group id
	InstallAppGroups(context.Context, *InstallAppGroupsRequest) (*IDsResponse, error)
	// UninstallAppGroups uninstall app groups.
	UninstallAppGroups(context.Context, *UninstallAppGroupsRequest) (*SimpleResponse, error)
	mustEmbedUnimplementedClusterLcmServiceServer()
}

// UnimplementedClusterLcmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClusterLcmServiceServer struct {
}

func (UnimplementedClusterLcmServiceServer) CreateCluster(context.Context, *CreateClusterRequest) (*IDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (UnimplementedClusterLcmServiceServer) ScaleCluster(context.Context, *ScaleClusterRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScaleCluster not implemented")
}
func (UnimplementedClusterLcmServiceServer) InstallAppGroups(context.Context, *InstallAppGroupsRequest) (*IDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallAppGroups not implemented")
}
func (UnimplementedClusterLcmServiceServer) UninstallAppGroups(context.Context, *UninstallAppGroupsRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UninstallAppGroups not implemented")
}
func (UnimplementedClusterLcmServiceServer) mustEmbedUnimplementedClusterLcmServiceServer() {}

// UnsafeClusterLcmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterLcmServiceServer will
// result in compilation errors.
type UnsafeClusterLcmServiceServer interface {
	mustEmbedUnimplementedClusterLcmServiceServer()
}

func RegisterClusterLcmServiceServer(s grpc.ServiceRegistrar, srv ClusterLcmServiceServer) {
	s.RegisterService(&ClusterLcmService_ServiceDesc, srv)
}

func _ClusterLcmService_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterLcmServiceServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.ClusterLcmService/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterLcmServiceServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterLcmService_ScaleCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterLcmServiceServer).ScaleCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.ClusterLcmService/ScaleCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterLcmServiceServer).ScaleCluster(ctx, req.(*ScaleClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterLcmService_InstallAppGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallAppGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterLcmServiceServer).InstallAppGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.ClusterLcmService/InstallAppGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterLcmServiceServer).InstallAppGroups(ctx, req.(*InstallAppGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterLcmService_UninstallAppGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UninstallAppGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterLcmServiceServer).UninstallAppGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.ClusterLcmService/UninstallAppGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterLcmServiceServer).UninstallAppGroups(ctx, req.(*UninstallAppGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterLcmService_ServiceDesc is the grpc.ServiceDesc for ClusterLcmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterLcmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbgo.ClusterLcmService",
	HandlerType: (*ClusterLcmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCluster",
			Handler:    _ClusterLcmService_CreateCluster_Handler,
		},
		{
			MethodName: "ScaleCluster",
			Handler:    _ClusterLcmService_ScaleCluster_Handler,
		},
		{
			MethodName: "InstallAppGroups",
			Handler:    _ClusterLcmService_InstallAppGroups_Handler,
		},
		{
			MethodName: "UninstallAppGroups",
			Handler:    _ClusterLcmService_UninstallAppGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster_lcm.proto",
}
