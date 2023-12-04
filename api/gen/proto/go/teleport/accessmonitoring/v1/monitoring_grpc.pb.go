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
// source: teleport/accessmonitoring/v1/monitoring.proto

package accessmonitoring

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
	MonitoringService_UpsertRule_FullMethodName = "/teleport.accessmonitoring.v1.MonitoringService/UpsertRule"
	MonitoringService_GetRule_FullMethodName    = "/teleport.accessmonitoring.v1.MonitoringService/GetRule"
	MonitoringService_ListRules_FullMethodName  = "/teleport.accessmonitoring.v1.MonitoringService/ListRules"
	MonitoringService_DeleteRule_FullMethodName = "/teleport.accessmonitoring.v1.MonitoringService/DeleteRule"
)

// MonitoringServiceClient is the client API for MonitoringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitoringServiceClient interface {
	// UpsertAuditQuery upsets an audit query.
	UpsertRule(ctx context.Context, in *UpsertRuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetAuditQuery returns an audit query.
	GetRule(ctx context.Context, in *GetRuleRequest, opts ...grpc.CallOption) (*Rule, error)
	// ListAuditQueries returns a paginated list of all Okta import rule resources.
	ListRules(ctx context.Context, in *ListRulesRequest, opts ...grpc.CallOption) (*ListRulesResponse, error)
	// DeleteAuditQuery deletes an audit query.
	DeleteRule(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type monitoringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitoringServiceClient(cc grpc.ClientConnInterface) MonitoringServiceClient {
	return &monitoringServiceClient{cc}
}

func (c *monitoringServiceClient) UpsertRule(ctx context.Context, in *UpsertRuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MonitoringService_UpsertRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringServiceClient) GetRule(ctx context.Context, in *GetRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.cc.Invoke(ctx, MonitoringService_GetRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringServiceClient) ListRules(ctx context.Context, in *ListRulesRequest, opts ...grpc.CallOption) (*ListRulesResponse, error) {
	out := new(ListRulesResponse)
	err := c.cc.Invoke(ctx, MonitoringService_ListRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringServiceClient) DeleteRule(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MonitoringService_DeleteRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitoringServiceServer is the server API for MonitoringService service.
// All implementations must embed UnimplementedMonitoringServiceServer
// for forward compatibility
type MonitoringServiceServer interface {
	// UpsertAuditQuery upsets an audit query.
	UpsertRule(context.Context, *UpsertRuleRequest) (*emptypb.Empty, error)
	// GetAuditQuery returns an audit query.
	GetRule(context.Context, *GetRuleRequest) (*Rule, error)
	// ListAuditQueries returns a paginated list of all Okta import rule resources.
	ListRules(context.Context, *ListRulesRequest) (*ListRulesResponse, error)
	// DeleteAuditQuery deletes an audit query.
	DeleteRule(context.Context, *DeleteRuleRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMonitoringServiceServer()
}

// UnimplementedMonitoringServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMonitoringServiceServer struct {
}

func (UnimplementedMonitoringServiceServer) UpsertRule(context.Context, *UpsertRuleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertRule not implemented")
}
func (UnimplementedMonitoringServiceServer) GetRule(context.Context, *GetRuleRequest) (*Rule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRule not implemented")
}
func (UnimplementedMonitoringServiceServer) ListRules(context.Context, *ListRulesRequest) (*ListRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRules not implemented")
}
func (UnimplementedMonitoringServiceServer) DeleteRule(context.Context, *DeleteRuleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRule not implemented")
}
func (UnimplementedMonitoringServiceServer) mustEmbedUnimplementedMonitoringServiceServer() {}

// UnsafeMonitoringServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitoringServiceServer will
// result in compilation errors.
type UnsafeMonitoringServiceServer interface {
	mustEmbedUnimplementedMonitoringServiceServer()
}

func RegisterMonitoringServiceServer(s grpc.ServiceRegistrar, srv MonitoringServiceServer) {
	s.RegisterService(&MonitoringService_ServiceDesc, srv)
}

func _MonitoringService_UpsertRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServiceServer).UpsertRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitoringService_UpsertRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServiceServer).UpsertRule(ctx, req.(*UpsertRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitoringService_GetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServiceServer).GetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitoringService_GetRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServiceServer).GetRule(ctx, req.(*GetRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitoringService_ListRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServiceServer).ListRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitoringService_ListRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServiceServer).ListRules(ctx, req.(*ListRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitoringService_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServiceServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitoringService_DeleteRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServiceServer).DeleteRule(ctx, req.(*DeleteRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MonitoringService_ServiceDesc is the grpc.ServiceDesc for MonitoringService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MonitoringService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.accessmonitoring.v1.MonitoringService",
	HandlerType: (*MonitoringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertRule",
			Handler:    _MonitoringService_UpsertRule_Handler,
		},
		{
			MethodName: "GetRule",
			Handler:    _MonitoringService_GetRule_Handler,
		},
		{
			MethodName: "ListRules",
			Handler:    _MonitoringService_ListRules_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _MonitoringService_DeleteRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/accessmonitoring/v1/monitoring.proto",
}
