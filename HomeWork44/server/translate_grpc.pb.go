// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: translate.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	TranslateService_GiveTranslation_FullMethodName = "/translateService.TranslateService/GiveTranslation"
)

// TranslateServiceClient is the client API for TranslateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranslateServiceClient interface {
	GiveTranslation(ctx context.Context, in *Massage, opts ...grpc.CallOption) (*Answer, error)
}

type translateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslateServiceClient(cc grpc.ClientConnInterface) TranslateServiceClient {
	return &translateServiceClient{cc}
}

func (c *translateServiceClient) GiveTranslation(ctx context.Context, in *Massage, opts ...grpc.CallOption) (*Answer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Answer)
	err := c.cc.Invoke(ctx, TranslateService_GiveTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslateServiceServer is the server API for TranslateService service.
// All implementations must embed UnimplementedTranslateServiceServer
// for forward compatibility
type TranslateServiceServer interface {
	GiveTranslation(context.Context, *Massage) (*Answer, error)
	mustEmbedUnimplementedTranslateServiceServer()
}

// UnimplementedTranslateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTranslateServiceServer struct {
}

func (UnimplementedTranslateServiceServer) GiveTranslation(context.Context, *Massage) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GiveTranslation not implemented")
}
func (UnimplementedTranslateServiceServer) mustEmbedUnimplementedTranslateServiceServer() {}

// UnsafeTranslateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranslateServiceServer will
// result in compilation errors.
type UnsafeTranslateServiceServer interface {
	mustEmbedUnimplementedTranslateServiceServer()
}

func RegisterTranslateServiceServer(s grpc.ServiceRegistrar, srv TranslateServiceServer) {
	s.RegisterService(&TranslateService_ServiceDesc, srv)
}

func _TranslateService_GiveTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Massage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServiceServer).GiveTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslateService_GiveTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServiceServer).GiveTranslation(ctx, req.(*Massage))
	}
	return interceptor(ctx, in, info, handler)
}

// TranslateService_ServiceDesc is the grpc.ServiceDesc for TranslateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TranslateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "translateService.TranslateService",
	HandlerType: (*TranslateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GiveTranslation",
			Handler:    _TranslateService_GiveTranslation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "translate.proto",
}
