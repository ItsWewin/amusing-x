// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package riskservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RiskServiceClient is the client API for RiskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RiskServiceClient interface {
	LoginRiskControl(ctx context.Context, in *LoginRiskRequest, opts ...grpc.CallOption) (*LoginRiskReply, error)
	Pong(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*PongReply, error)
}

type riskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRiskServiceClient(cc grpc.ClientConnInterface) RiskServiceClient {
	return &riskServiceClient{cc}
}

func (c *riskServiceClient) LoginRiskControl(ctx context.Context, in *LoginRiskRequest, opts ...grpc.CallOption) (*LoginRiskReply, error) {
	out := new(LoginRiskReply)
	err := c.cc.Invoke(ctx, "/riskservice.RiskService/LoginRiskControl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riskServiceClient) Pong(ctx context.Context, in *BlankParams, opts ...grpc.CallOption) (*PongReply, error) {
	out := new(PongReply)
	err := c.cc.Invoke(ctx, "/riskservice.RiskService/Pong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RiskServiceServer is the server API for RiskService service.
// All implementations must embed UnimplementedRiskServiceServer
// for forward compatibility
type RiskServiceServer interface {
	LoginRiskControl(context.Context, *LoginRiskRequest) (*LoginRiskReply, error)
	Pong(context.Context, *BlankParams) (*PongReply, error)
	mustEmbedUnimplementedRiskServiceServer()
}

// UnimplementedRiskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRiskServiceServer struct {
}

func (*UnimplementedRiskServiceServer) LoginRiskControl(context.Context, *LoginRiskRequest) (*LoginRiskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginRiskControl not implemented")
}
func (*UnimplementedRiskServiceServer) Pong(context.Context, *BlankParams) (*PongReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pong not implemented")
}
func (*UnimplementedRiskServiceServer) mustEmbedUnimplementedRiskServiceServer() {}

func RegisterRiskServiceServer(s *grpc.Server, srv RiskServiceServer) {
	s.RegisterService(&_RiskService_serviceDesc, srv)
}

func _RiskService_LoginRiskControl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskServiceServer).LoginRiskControl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/riskservice.RiskService/LoginRiskControl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskServiceServer).LoginRiskControl(ctx, req.(*LoginRiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiskService_Pong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlankParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskServiceServer).Pong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/riskservice.RiskService/Pong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskServiceServer).Pong(ctx, req.(*BlankParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _RiskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "riskservice.RiskService",
	HandlerType: (*RiskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginRiskControl",
			Handler:    _RiskService_LoginRiskControl_Handler,
		},
		{
			MethodName: "Pong",
			Handler:    _RiskService_Pong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}