// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package powerTrainAggregator

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

// PTEsimateServiceClient is the client API for PTEsimateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PTEsimateServiceClient interface {
	EstimatePowerTrain(ctx context.Context, in *PTEstimateRequest, opts ...grpc.CallOption) (*PTEstimateResponse, error)
}

type pTEsimateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPTEsimateServiceClient(cc grpc.ClientConnInterface) PTEsimateServiceClient {
	return &pTEsimateServiceClient{cc}
}

func (c *pTEsimateServiceClient) EstimatePowerTrain(ctx context.Context, in *PTEstimateRequest, opts ...grpc.CallOption) (*PTEstimateResponse, error) {
	out := new(PTEstimateResponse)
	err := c.cc.Invoke(ctx, "/powerTrainAggregator.v1.PTEsimateService/EstimatePowerTrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PTEsimateServiceServer is the server API for PTEsimateService service.
// All implementations must embed UnimplementedPTEsimateServiceServer
// for forward compatibility
type PTEsimateServiceServer interface {
	EstimatePowerTrain(context.Context, *PTEstimateRequest) (*PTEstimateResponse, error)
	mustEmbedUnimplementedPTEsimateServiceServer()
}

// UnimplementedPTEsimateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPTEsimateServiceServer struct {
}

func (UnimplementedPTEsimateServiceServer) EstimatePowerTrain(context.Context, *PTEstimateRequest) (*PTEstimateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimatePowerTrain not implemented")
}
func (UnimplementedPTEsimateServiceServer) mustEmbedUnimplementedPTEsimateServiceServer() {}

// UnsafePTEsimateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PTEsimateServiceServer will
// result in compilation errors.
type UnsafePTEsimateServiceServer interface {
	mustEmbedUnimplementedPTEsimateServiceServer()
}

func RegisterPTEsimateServiceServer(s grpc.ServiceRegistrar, srv PTEsimateServiceServer) {
	s.RegisterService(&PTEsimateService_ServiceDesc, srv)
}

func _PTEsimateService_EstimatePowerTrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PTEstimateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PTEsimateServiceServer).EstimatePowerTrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerTrainAggregator.v1.PTEsimateService/EstimatePowerTrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PTEsimateServiceServer).EstimatePowerTrain(ctx, req.(*PTEstimateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PTEsimateService_ServiceDesc is the grpc.ServiceDesc for PTEsimateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PTEsimateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "powerTrainAggregator.v1.PTEsimateService",
	HandlerType: (*PTEsimateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EstimatePowerTrain",
			Handler:    _PTEsimateService_EstimatePowerTrain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "powerTrainAggregator/proto/v1/power_train_aggregator_api_v1.proto",
}
