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

// AppLcmServiceClient is the client API for AppLcmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppLcmServiceClient interface {
	// InstallApps install apps, return a array of application id
	InstallApps(ctx context.Context, in *InstallAppsRequest, opts ...grpc.CallOption) (*IDsResponse, error)
	// UninstallApps uninstall apps
	UninstallApps(ctx context.Context, in *UninstallAppsRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
}

type appLcmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppLcmServiceClient(cc grpc.ClientConnInterface) AppLcmServiceClient {
	return &appLcmServiceClient{cc}
}

func (c *appLcmServiceClient) InstallApps(ctx context.Context, in *InstallAppsRequest, opts ...grpc.CallOption) (*IDsResponse, error) {
	out := new(IDsResponse)
	err := c.cc.Invoke(ctx, "/pbgo.AppLcmService/InstallApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appLcmServiceClient) UninstallApps(ctx context.Context, in *UninstallAppsRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/pbgo.AppLcmService/UninstallApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppLcmServiceServer is the server API for AppLcmService service.
// All implementations must embed UnimplementedAppLcmServiceServer
// for forward compatibility
type AppLcmServiceServer interface {
	// InstallApps install apps, return a array of application id
	InstallApps(context.Context, *InstallAppsRequest) (*IDsResponse, error)
	// UninstallApps uninstall apps
	UninstallApps(context.Context, *UninstallAppsRequest) (*SimpleResponse, error)
	mustEmbedUnimplementedAppLcmServiceServer()
}

// UnimplementedAppLcmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppLcmServiceServer struct {
}

func (UnimplementedAppLcmServiceServer) InstallApps(context.Context, *InstallAppsRequest) (*IDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallApps not implemented")
}
func (UnimplementedAppLcmServiceServer) UninstallApps(context.Context, *UninstallAppsRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UninstallApps not implemented")
}
func (UnimplementedAppLcmServiceServer) mustEmbedUnimplementedAppLcmServiceServer() {}

// UnsafeAppLcmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppLcmServiceServer will
// result in compilation errors.
type UnsafeAppLcmServiceServer interface {
	mustEmbedUnimplementedAppLcmServiceServer()
}

func RegisterAppLcmServiceServer(s grpc.ServiceRegistrar, srv AppLcmServiceServer) {
	s.RegisterService(&AppLcmService_ServiceDesc, srv)
}

func _AppLcmService_InstallApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppLcmServiceServer).InstallApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.AppLcmService/InstallApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppLcmServiceServer).InstallApps(ctx, req.(*InstallAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppLcmService_UninstallApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UninstallAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppLcmServiceServer).UninstallApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgo.AppLcmService/UninstallApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppLcmServiceServer).UninstallApps(ctx, req.(*UninstallAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppLcmService_ServiceDesc is the grpc.ServiceDesc for AppLcmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppLcmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbgo.AppLcmService",
	HandlerType: (*AppLcmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InstallApps",
			Handler:    _AppLcmService_InstallApps_Handler,
		},
		{
			MethodName: "UninstallApps",
			Handler:    _AppLcmService_UninstallApps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application_lcm.proto",
}