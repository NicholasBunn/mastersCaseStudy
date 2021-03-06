// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package comfortService

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

// ComfortServiceClient is the client API for ComfortService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComfortServiceClient interface {
	// The 'Comfort Rating' call provides foresight for tactical decision-making by providing a comfort rating for a proposed route, based on estimated vibrations on board.
	ComfortRating(ctx context.Context, in *ComfortRequest, opts ...grpc.CallOption) (*ComfortResponse, error)
}

type comfortServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComfortServiceClient(cc grpc.ClientConnInterface) ComfortServiceClient {
	return &comfortServiceClient{cc}
}

func (c *comfortServiceClient) ComfortRating(ctx context.Context, in *ComfortRequest, opts ...grpc.CallOption) (*ComfortResponse, error) {
	out := new(ComfortResponse)
	err := c.cc.Invoke(ctx, "/comfortServiceAPI.v1.ComfortService/ComfortRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComfortServiceServer is the server API for ComfortService service.
// All implementations must embed UnimplementedComfortServiceServer
// for forward compatibility
type ComfortServiceServer interface {
	// The 'Comfort Rating' call provides foresight for tactical decision-making by providing a comfort rating for a proposed route, based on estimated vibrations on board.
	ComfortRating(context.Context, *ComfortRequest) (*ComfortResponse, error)
	mustEmbedUnimplementedComfortServiceServer()
}

// UnimplementedComfortServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComfortServiceServer struct {
}

func (UnimplementedComfortServiceServer) ComfortRating(context.Context, *ComfortRequest) (*ComfortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComfortRating not implemented")
}
func (UnimplementedComfortServiceServer) mustEmbedUnimplementedComfortServiceServer() {}

// UnsafeComfortServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComfortServiceServer will
// result in compilation errors.
type UnsafeComfortServiceServer interface {
	mustEmbedUnimplementedComfortServiceServer()
}

func RegisterComfortServiceServer(s grpc.ServiceRegistrar, srv ComfortServiceServer) {
	s.RegisterService(&ComfortService_ServiceDesc, srv)
}

func _ComfortService_ComfortRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComfortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComfortServiceServer).ComfortRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comfortServiceAPI.v1.ComfortService/ComfortRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComfortServiceServer).ComfortRating(ctx, req.(*ComfortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComfortService_ServiceDesc is the grpc.ServiceDesc for ComfortService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComfortService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comfortServiceAPI.v1.ComfortService",
	HandlerType: (*ComfortServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComfortRating",
			Handler:    _ComfortService_ComfortRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comfortService/proto/v1/comfort_service_api_v1.proto",
}
