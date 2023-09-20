// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: api/verifyCode/verifyCode.proto

package verifyCode

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
	VerifyCode_GetVerifyCode_FullMethodName = "/api.verifyCode.VerifyCode/GetVerifyCode"
)

// VerifyCodeClient is the client API for VerifyCode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerifyCodeClient interface {
	GetVerifyCode(ctx context.Context, in *GetVerifyCodeRequest, opts ...grpc.CallOption) (*GetVerifyCodeReply, error)
}

type verifyCodeClient struct {
	cc grpc.ClientConnInterface
}

func NewVerifyCodeClient(cc grpc.ClientConnInterface) VerifyCodeClient {
	return &verifyCodeClient{cc}
}

func (c *verifyCodeClient) GetVerifyCode(ctx context.Context, in *GetVerifyCodeRequest, opts ...grpc.CallOption) (*GetVerifyCodeReply, error) {
	out := new(GetVerifyCodeReply)
	err := c.cc.Invoke(ctx, VerifyCode_GetVerifyCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifyCodeServer is the server API for VerifyCode service.
// All implementations must embed UnimplementedVerifyCodeServer
// for forward compatibility
type VerifyCodeServer interface {
	GetVerifyCode(context.Context, *GetVerifyCodeRequest) (*GetVerifyCodeReply, error)
	mustEmbedUnimplementedVerifyCodeServer()
}

// UnimplementedVerifyCodeServer must be embedded to have forward compatible implementations.
type UnimplementedVerifyCodeServer struct {
}

func (UnimplementedVerifyCodeServer) GetVerifyCode(context.Context, *GetVerifyCodeRequest) (*GetVerifyCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVerifyCode not implemented")
}
func (UnimplementedVerifyCodeServer) mustEmbedUnimplementedVerifyCodeServer() {}

// UnsafeVerifyCodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerifyCodeServer will
// result in compilation errors.
type UnsafeVerifyCodeServer interface {
	mustEmbedUnimplementedVerifyCodeServer()
}

func RegisterVerifyCodeServer(s grpc.ServiceRegistrar, srv VerifyCodeServer) {
	s.RegisterService(&VerifyCode_ServiceDesc, srv)
}

func _VerifyCode_GetVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyCodeServer).GetVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerifyCode_GetVerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyCodeServer).GetVerifyCode(ctx, req.(*GetVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VerifyCode_ServiceDesc is the grpc.ServiceDesc for VerifyCode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VerifyCode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.verifyCode.VerifyCode",
	HandlerType: (*VerifyCodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVerifyCode",
			Handler:    _VerifyCode_GetVerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/verifyCode/verifyCode.proto",
}
