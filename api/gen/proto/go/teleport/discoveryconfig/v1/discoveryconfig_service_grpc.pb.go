// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: teleport/discoveryconfig/v1/discoveryconfig_service.proto

package discoveryconfigv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DiscoveryConfigService_ListDiscoveryConfigs_FullMethodName        = "/teleport.discoveryconfig.v1.DiscoveryConfigService/ListDiscoveryConfigs"
	DiscoveryConfigService_GetDiscoveryConfig_FullMethodName          = "/teleport.discoveryconfig.v1.DiscoveryConfigService/GetDiscoveryConfig"
	DiscoveryConfigService_CreateDiscoveryConfig_FullMethodName       = "/teleport.discoveryconfig.v1.DiscoveryConfigService/CreateDiscoveryConfig"
	DiscoveryConfigService_UpdateDiscoveryConfig_FullMethodName       = "/teleport.discoveryconfig.v1.DiscoveryConfigService/UpdateDiscoveryConfig"
	DiscoveryConfigService_UpsertDiscoveryConfig_FullMethodName       = "/teleport.discoveryconfig.v1.DiscoveryConfigService/UpsertDiscoveryConfig"
	DiscoveryConfigService_DeleteDiscoveryConfig_FullMethodName       = "/teleport.discoveryconfig.v1.DiscoveryConfigService/DeleteDiscoveryConfig"
	DiscoveryConfigService_DeleteAllDiscoveryConfigs_FullMethodName   = "/teleport.discoveryconfig.v1.DiscoveryConfigService/DeleteAllDiscoveryConfigs"
	DiscoveryConfigService_UpdateDiscoveryConfigStatus_FullMethodName = "/teleport.discoveryconfig.v1.DiscoveryConfigService/UpdateDiscoveryConfigStatus"
)

// DiscoveryConfigServiceClient is the client API for DiscoveryConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscoveryConfigServiceClient interface {
	// ListDiscoveryConfigs returns a paginated list of Discovery Config resources.
	ListDiscoveryConfigs(ctx context.Context, in *ListDiscoveryConfigsRequest, opts ...grpc.CallOption) (*ListDiscoveryConfigsResponse, error)
	// GetDiscoveryConfig returns the specified DiscoveryConfig resource.
	GetDiscoveryConfig(ctx context.Context, in *GetDiscoveryConfigRequest, opts ...grpc.CallOption) (*DiscoveryConfig, error)
	// CreateDiscoveryConfig creates a new DiscoveryConfig resource.
	CreateDiscoveryConfig(ctx context.Context, in *CreateDiscoveryConfigRequest, opts ...grpc.CallOption) (*DiscoveryConfig, error)
	// UpdateDiscoveryConfig updates an existing DiscoveryConfig resource.
	UpdateDiscoveryConfig(ctx context.Context, in *UpdateDiscoveryConfigRequest, opts ...grpc.CallOption) (*DiscoveryConfig, error)
	// UpsertDiscoveryConfig creates or updates a DiscoveryConfig resource.
	UpsertDiscoveryConfig(ctx context.Context, in *UpsertDiscoveryConfigRequest, opts ...grpc.CallOption) (*DiscoveryConfig, error)
	// DeleteDiscoveryConfig removes the specified DiscoveryConfig resource.
	DeleteDiscoveryConfig(ctx context.Context, in *DeleteDiscoveryConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteAllDiscoveryConfigs removes all DiscoveryConfigs.
	DeleteAllDiscoveryConfigs(ctx context.Context, in *DeleteAllDiscoveryConfigsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UpdateDiscoveryConfigStatus updates an existing DiscoveryConfig resource status object.
	UpdateDiscoveryConfigStatus(ctx context.Context, in *UpdateDiscoveryConfigStatusRequest, opts ...grpc.CallOption) (*DiscoveryConfig, error)
}

type discoveryConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoveryConfigServiceClient(cc grpc.ClientConnInterface) DiscoveryConfigServiceClient {
	return &discoveryConfigServiceClient{cc}
}

func (c *discoveryConfigServiceClient) ListDiscoveryConfigs(ctx context.Context, in *ListDiscoveryConfigsRequest, opts ...grpc.CallOption) (*ListDiscoveryConfigsResponse, error) {
	out := new(ListDiscoveryConfigsResponse)
	err := c.cc.Invoke(ctx, DiscoveryConfigService_ListDiscoveryConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryConfigServiceClient) GetDiscoveryConfig(ctx context.Context, in *GetDiscoveryConfigRequest, opts ...grpc.CallOption) (*DiscoveryConfig, error) {
	out := new(DiscoveryConfig)
	err := c.cc.Invoke(ctx, DiscoveryConfigService_GetDiscoveryConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryConfigServiceClient) CreateDiscoveryConfig(ctx context.Context, in *CreateDiscoveryConfigRequest, opts ...grpc.CallOption) (*DiscoveryConfig, error) {
	out := new(DiscoveryConfig)
	err := c.cc.Invoke(ctx, DiscoveryConfigService_CreateDiscoveryConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryConfigServiceClient) UpdateDiscoveryConfig(ctx context.Context, in *UpdateDiscoveryConfigRequest, opts ...grpc.CallOption) (*DiscoveryConfig, error) {
	out := new(DiscoveryConfig)
	err := c.cc.Invoke(ctx, DiscoveryConfigService_UpdateDiscoveryConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryConfigServiceClient) UpsertDiscoveryConfig(ctx context.Context, in *UpsertDiscoveryConfigRequest, opts ...grpc.CallOption) (*DiscoveryConfig, error) {
	out := new(DiscoveryConfig)
	err := c.cc.Invoke(ctx, DiscoveryConfigService_UpsertDiscoveryConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryConfigServiceClient) DeleteDiscoveryConfig(ctx context.Context, in *DeleteDiscoveryConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DiscoveryConfigService_DeleteDiscoveryConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryConfigServiceClient) DeleteAllDiscoveryConfigs(ctx context.Context, in *DeleteAllDiscoveryConfigsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DiscoveryConfigService_DeleteAllDiscoveryConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryConfigServiceClient) UpdateDiscoveryConfigStatus(ctx context.Context, in *UpdateDiscoveryConfigStatusRequest, opts ...grpc.CallOption) (*DiscoveryConfig, error) {
	out := new(DiscoveryConfig)
	err := c.cc.Invoke(ctx, DiscoveryConfigService_UpdateDiscoveryConfigStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryConfigServiceServer is the server API for DiscoveryConfigService service.
// All implementations must embed UnimplementedDiscoveryConfigServiceServer
// for forward compatibility
type DiscoveryConfigServiceServer interface {
	// ListDiscoveryConfigs returns a paginated list of Discovery Config resources.
	ListDiscoveryConfigs(context.Context, *ListDiscoveryConfigsRequest) (*ListDiscoveryConfigsResponse, error)
	// GetDiscoveryConfig returns the specified DiscoveryConfig resource.
	GetDiscoveryConfig(context.Context, *GetDiscoveryConfigRequest) (*DiscoveryConfig, error)
	// CreateDiscoveryConfig creates a new DiscoveryConfig resource.
	CreateDiscoveryConfig(context.Context, *CreateDiscoveryConfigRequest) (*DiscoveryConfig, error)
	// UpdateDiscoveryConfig updates an existing DiscoveryConfig resource.
	UpdateDiscoveryConfig(context.Context, *UpdateDiscoveryConfigRequest) (*DiscoveryConfig, error)
	// UpsertDiscoveryConfig creates or updates a DiscoveryConfig resource.
	UpsertDiscoveryConfig(context.Context, *UpsertDiscoveryConfigRequest) (*DiscoveryConfig, error)
	// DeleteDiscoveryConfig removes the specified DiscoveryConfig resource.
	DeleteDiscoveryConfig(context.Context, *DeleteDiscoveryConfigRequest) (*emptypb.Empty, error)
	// DeleteAllDiscoveryConfigs removes all DiscoveryConfigs.
	DeleteAllDiscoveryConfigs(context.Context, *DeleteAllDiscoveryConfigsRequest) (*emptypb.Empty, error)
	// UpdateDiscoveryConfigStatus updates an existing DiscoveryConfig resource status object.
	UpdateDiscoveryConfigStatus(context.Context, *UpdateDiscoveryConfigStatusRequest) (*DiscoveryConfig, error)
	mustEmbedUnimplementedDiscoveryConfigServiceServer()
}

// UnimplementedDiscoveryConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiscoveryConfigServiceServer struct {
}

func (UnimplementedDiscoveryConfigServiceServer) ListDiscoveryConfigs(context.Context, *ListDiscoveryConfigsRequest) (*ListDiscoveryConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDiscoveryConfigs not implemented")
}
func (UnimplementedDiscoveryConfigServiceServer) GetDiscoveryConfig(context.Context, *GetDiscoveryConfigRequest) (*DiscoveryConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscoveryConfig not implemented")
}
func (UnimplementedDiscoveryConfigServiceServer) CreateDiscoveryConfig(context.Context, *CreateDiscoveryConfigRequest) (*DiscoveryConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDiscoveryConfig not implemented")
}
func (UnimplementedDiscoveryConfigServiceServer) UpdateDiscoveryConfig(context.Context, *UpdateDiscoveryConfigRequest) (*DiscoveryConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDiscoveryConfig not implemented")
}
func (UnimplementedDiscoveryConfigServiceServer) UpsertDiscoveryConfig(context.Context, *UpsertDiscoveryConfigRequest) (*DiscoveryConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertDiscoveryConfig not implemented")
}
func (UnimplementedDiscoveryConfigServiceServer) DeleteDiscoveryConfig(context.Context, *DeleteDiscoveryConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDiscoveryConfig not implemented")
}
func (UnimplementedDiscoveryConfigServiceServer) DeleteAllDiscoveryConfigs(context.Context, *DeleteAllDiscoveryConfigsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllDiscoveryConfigs not implemented")
}
func (UnimplementedDiscoveryConfigServiceServer) UpdateDiscoveryConfigStatus(context.Context, *UpdateDiscoveryConfigStatusRequest) (*DiscoveryConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDiscoveryConfigStatus not implemented")
}
func (UnimplementedDiscoveryConfigServiceServer) mustEmbedUnimplementedDiscoveryConfigServiceServer() {
}

// UnsafeDiscoveryConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscoveryConfigServiceServer will
// result in compilation errors.
type UnsafeDiscoveryConfigServiceServer interface {
	mustEmbedUnimplementedDiscoveryConfigServiceServer()
}

func RegisterDiscoveryConfigServiceServer(s grpc.ServiceRegistrar, srv DiscoveryConfigServiceServer) {
	s.RegisterService(&DiscoveryConfigService_ServiceDesc, srv)
}

func _DiscoveryConfigService_ListDiscoveryConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiscoveryConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryConfigServiceServer).ListDiscoveryConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryConfigService_ListDiscoveryConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryConfigServiceServer).ListDiscoveryConfigs(ctx, req.(*ListDiscoveryConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryConfigService_GetDiscoveryConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiscoveryConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryConfigServiceServer).GetDiscoveryConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryConfigService_GetDiscoveryConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryConfigServiceServer).GetDiscoveryConfig(ctx, req.(*GetDiscoveryConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryConfigService_CreateDiscoveryConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDiscoveryConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryConfigServiceServer).CreateDiscoveryConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryConfigService_CreateDiscoveryConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryConfigServiceServer).CreateDiscoveryConfig(ctx, req.(*CreateDiscoveryConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryConfigService_UpdateDiscoveryConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDiscoveryConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryConfigServiceServer).UpdateDiscoveryConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryConfigService_UpdateDiscoveryConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryConfigServiceServer).UpdateDiscoveryConfig(ctx, req.(*UpdateDiscoveryConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryConfigService_UpsertDiscoveryConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertDiscoveryConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryConfigServiceServer).UpsertDiscoveryConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryConfigService_UpsertDiscoveryConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryConfigServiceServer).UpsertDiscoveryConfig(ctx, req.(*UpsertDiscoveryConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryConfigService_DeleteDiscoveryConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDiscoveryConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryConfigServiceServer).DeleteDiscoveryConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryConfigService_DeleteDiscoveryConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryConfigServiceServer).DeleteDiscoveryConfig(ctx, req.(*DeleteDiscoveryConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryConfigService_DeleteAllDiscoveryConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllDiscoveryConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryConfigServiceServer).DeleteAllDiscoveryConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryConfigService_DeleteAllDiscoveryConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryConfigServiceServer).DeleteAllDiscoveryConfigs(ctx, req.(*DeleteAllDiscoveryConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryConfigService_UpdateDiscoveryConfigStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDiscoveryConfigStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryConfigServiceServer).UpdateDiscoveryConfigStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryConfigService_UpdateDiscoveryConfigStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryConfigServiceServer).UpdateDiscoveryConfigStatus(ctx, req.(*UpdateDiscoveryConfigStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscoveryConfigService_ServiceDesc is the grpc.ServiceDesc for DiscoveryConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscoveryConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.discoveryconfig.v1.DiscoveryConfigService",
	HandlerType: (*DiscoveryConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDiscoveryConfigs",
			Handler:    _DiscoveryConfigService_ListDiscoveryConfigs_Handler,
		},
		{
			MethodName: "GetDiscoveryConfig",
			Handler:    _DiscoveryConfigService_GetDiscoveryConfig_Handler,
		},
		{
			MethodName: "CreateDiscoveryConfig",
			Handler:    _DiscoveryConfigService_CreateDiscoveryConfig_Handler,
		},
		{
			MethodName: "UpdateDiscoveryConfig",
			Handler:    _DiscoveryConfigService_UpdateDiscoveryConfig_Handler,
		},
		{
			MethodName: "UpsertDiscoveryConfig",
			Handler:    _DiscoveryConfigService_UpsertDiscoveryConfig_Handler,
		},
		{
			MethodName: "DeleteDiscoveryConfig",
			Handler:    _DiscoveryConfigService_DeleteDiscoveryConfig_Handler,
		},
		{
			MethodName: "DeleteAllDiscoveryConfigs",
			Handler:    _DiscoveryConfigService_DeleteAllDiscoveryConfigs_Handler,
		},
		{
			MethodName: "UpdateDiscoveryConfigStatus",
			Handler:    _DiscoveryConfigService_UpdateDiscoveryConfigStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/discoveryconfig/v1/discoveryconfig_service.proto",
}
