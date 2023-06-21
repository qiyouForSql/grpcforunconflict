// Copyright 2015 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The canonical version of this proto can be found at
// https://github.com/grpc/grpc-proto/blob/master/grpc/health/v1/health.proto

// Code generated by protoc-gen-go-grpcforunconflict. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: grpc/health/v1/health.proto

package grpc_health_v1

import (
	context "context"
	grpcforunconflict "github.com/qiyouForSql/grpcforunconflict"
	codes "github.com/qiyouForSql/grpcforunconflict/codes"
	status "github.com/qiyouForSql/grpcforunconflict/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ =grpcforunconflict.SupportPackageIsVersion7

const (
	Health_Check_FullMethodName = "/grpcforunconflict.health.v1.Health/Check"
	Health_Watch_FullMethodName = "/grpcforunconflict.health.v1.Health/Watch"
)

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthClient interface {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpcforunconflict.CallOption) (*HealthCheckResponse, error)
	// Performs a watch for the serving status of the requested service.
	// The server will immediately send back a message indicating the current
	// serving status.  It will then subsequently send a new message whenever
	// the service's serving status changes.
	//
	// If the requested service is unknown when the call is received, the
	// server will send a message setting the serving status to
	// SERVICE_UNKNOWN but will *not* terminate the call.  If at some
	// future point, the serving status of the service becomes known, the
	// server will send a new message with the service's serving status.
	//
	// If the call terminates with status UNIMPLEMENTED, then clients
	// should assume this method is not supported and should not retry the
	// call.  If the call terminates with any other status (including OK),
	// clients should retry the call with appropriate exponential backoff.
	Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpcforunconflict.CallOption) (Health_WatchClient, error)
}

type healthClient struct {
	ccgrpcforunconflict.ClientConnInterface
}

func NewHealthClient(ccgrpcforunconflict.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpcforunconflict.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, Health_Check_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthClient) Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpcforunconflict.CallOption) (Health_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Health_ServiceDesc.Streams[0], Health_Watch_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &healthWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Health_WatchClient interface {
	Recv() (*HealthCheckResponse, error)
	grpcforunconflict.ClientStream
}

type healthWatchClient struct {
	grpcforunconflict.ClientStream
}

func (x *healthWatchClient) Recv() (*HealthCheckResponse, error) {
	m := new(HealthCheckResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HealthServer is the server API for Health service.
// All implementations should embed UnimplementedHealthServer
// for forward compatibility
type HealthServer interface {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	// Performs a watch for the serving status of the requested service.
	// The server will immediately send back a message indicating the current
	// serving status.  It will then subsequently send a new message whenever
	// the service's serving status changes.
	//
	// If the requested service is unknown when the call is received, the
	// server will send a message setting the serving status to
	// SERVICE_UNKNOWN but will *not* terminate the call.  If at some
	// future point, the serving status of the service becomes known, the
	// server will send a new message with the service's serving status.
	//
	// If the call terminates with status UNIMPLEMENTED, then clients
	// should assume this method is not supported and should not retry the
	// call.  If the call terminates with any other status (including OK),
	// clients should retry the call with appropriate exponential backoff.
	Watch(*HealthCheckRequest, Health_WatchServer) error
}

// UnimplementedHealthServer should be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (UnimplementedHealthServer) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedHealthServer) Watch(*HealthCheckRequest, Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

// UnsafeHealthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServer will
// result in compilation errors.
type UnsafeHealthServer interface {
	mustEmbedUnimplementedHealthServer()
}

func RegisterHealthServer(sgrpcforunconflict.ServiceRegistrar, srv HealthServer) {
	s.RegisterService(&Health_ServiceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptorgrpcforunconflict.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpcforunconflict.UnaryServerInfo{
		Server:     srv,
		FullMethod: Health_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Health_Watch_Handler(srv interface{}, streamgrpcforunconflict.ServerStream) error {
	m := new(HealthCheckRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HealthServer).Watch(m, &healthWatchServer{stream})
}

type Health_WatchServer interface {
	Send(*HealthCheckResponse) error
	grpcforunconflict.ServerStream
}

type healthWatchServer struct {
	grpcforunconflict.ServerStream
}

func (x *healthWatchServer) Send(m *HealthCheckResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Health_ServiceDesc is thegrpcforunconflict.ServiceDesc for Health service.
// It's only intended for direct use withgrpcforunconflict.RegisterService,
// and not to be introspected or modified (even as a copy)
var Health_ServiceDesc =grpcforunconflict.ServiceDesc{
	ServiceName: "grpcforunconflict.health.v1.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpcforunconflict.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams: []grpcforunconflict.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Health_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/health/v1/health.proto",
}
