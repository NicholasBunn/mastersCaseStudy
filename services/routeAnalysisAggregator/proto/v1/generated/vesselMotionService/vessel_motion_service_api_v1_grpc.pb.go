// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vesselMotionService

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

// VesselMotionServiceClient is the client API for VesselMotionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VesselMotionServiceClient interface {
	// The 'Motion Estimate' call provides foresight for tactical decision-making by providing high-frequency acceleration estimates for a requested sailing conditions at a requested location on the ship
	MotionEstimate(ctx context.Context, in *MotionEstimateRequest, opts ...grpc.CallOption) (*MotionEstimateResponse, error)
	// The 'Motion Tracking' call provides insight for tactical decision-making by providing real-time, high-frequency acceleration readings for a requested location on the ship
	MotionTracking(ctx context.Context, in *MotionTrackingRequest, opts ...grpc.CallOption) (*MotionTrackingResponse, error)
	// The 'Motion Estimation Evaluation' call provides hindsight for strategic decision-making by evaluating the accuracy of the models predictions
	MotionEstimateEvaluation(ctx context.Context, in *MotionEstimateRequest, opts ...grpc.CallOption) (*MotionEvaluationResponse, error)
}

type vesselMotionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVesselMotionServiceClient(cc grpc.ClientConnInterface) VesselMotionServiceClient {
	return &vesselMotionServiceClient{cc}
}

func (c *vesselMotionServiceClient) MotionEstimate(ctx context.Context, in *MotionEstimateRequest, opts ...grpc.CallOption) (*MotionEstimateResponse, error) {
	out := new(MotionEstimateResponse)
	err := c.cc.Invoke(ctx, "/vesselMotionService.v1.vesselMotionService/MotionEstimate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselMotionServiceClient) MotionTracking(ctx context.Context, in *MotionTrackingRequest, opts ...grpc.CallOption) (*MotionTrackingResponse, error) {
	out := new(MotionTrackingResponse)
	err := c.cc.Invoke(ctx, "/vesselMotionService.v1.vesselMotionService/MotionTracking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselMotionServiceClient) MotionEstimateEvaluation(ctx context.Context, in *MotionEstimateRequest, opts ...grpc.CallOption) (*MotionEvaluationResponse, error) {
	out := new(MotionEvaluationResponse)
	err := c.cc.Invoke(ctx, "/vesselMotionService.v1.vesselMotionService/MotionEstimateEvaluation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VesselMotionServiceServer is the server API for VesselMotionService service.
// All implementations must embed UnimplementedVesselMotionServiceServer
// for forward compatibility
type VesselMotionServiceServer interface {
	// The 'Motion Estimate' call provides foresight for tactical decision-making by providing high-frequency acceleration estimates for a requested sailing conditions at a requested location on the ship
	MotionEstimate(context.Context, *MotionEstimateRequest) (*MotionEstimateResponse, error)
	// The 'Motion Tracking' call provides insight for tactical decision-making by providing real-time, high-frequency acceleration readings for a requested location on the ship
	MotionTracking(context.Context, *MotionTrackingRequest) (*MotionTrackingResponse, error)
	// The 'Motion Estimation Evaluation' call provides hindsight for strategic decision-making by evaluating the accuracy of the models predictions
	MotionEstimateEvaluation(context.Context, *MotionEstimateRequest) (*MotionEvaluationResponse, error)
	mustEmbedUnimplementedVesselMotionServiceServer()
}

// UnimplementedVesselMotionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVesselMotionServiceServer struct {
}

func (UnimplementedVesselMotionServiceServer) MotionEstimate(context.Context, *MotionEstimateRequest) (*MotionEstimateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MotionEstimate not implemented")
}
func (UnimplementedVesselMotionServiceServer) MotionTracking(context.Context, *MotionTrackingRequest) (*MotionTrackingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MotionTracking not implemented")
}
func (UnimplementedVesselMotionServiceServer) MotionEstimateEvaluation(context.Context, *MotionEstimateRequest) (*MotionEvaluationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MotionEstimateEvaluation not implemented")
}
func (UnimplementedVesselMotionServiceServer) mustEmbedUnimplementedVesselMotionServiceServer() {}

// UnsafeVesselMotionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VesselMotionServiceServer will
// result in compilation errors.
type UnsafeVesselMotionServiceServer interface {
	mustEmbedUnimplementedVesselMotionServiceServer()
}

func RegisterVesselMotionServiceServer(s grpc.ServiceRegistrar, srv VesselMotionServiceServer) {
	s.RegisterService(&VesselMotionService_ServiceDesc, srv)
}

func _VesselMotionService_MotionEstimate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotionEstimateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VesselMotionServiceServer).MotionEstimate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vesselMotionService.v1.vesselMotionService/MotionEstimate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VesselMotionServiceServer).MotionEstimate(ctx, req.(*MotionEstimateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VesselMotionService_MotionTracking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotionTrackingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VesselMotionServiceServer).MotionTracking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vesselMotionService.v1.vesselMotionService/MotionTracking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VesselMotionServiceServer).MotionTracking(ctx, req.(*MotionTrackingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VesselMotionService_MotionEstimateEvaluation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotionEstimateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VesselMotionServiceServer).MotionEstimateEvaluation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vesselMotionService.v1.vesselMotionService/MotionEstimateEvaluation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VesselMotionServiceServer).MotionEstimateEvaluation(ctx, req.(*MotionEstimateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VesselMotionService_ServiceDesc is the grpc.ServiceDesc for VesselMotionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VesselMotionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vesselMotionService.v1.vesselMotionService",
	HandlerType: (*VesselMotionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MotionEstimate",
			Handler:    _VesselMotionService_MotionEstimate_Handler,
		},
		{
			MethodName: "MotionTracking",
			Handler:    _VesselMotionService_MotionTracking_Handler,
		},
		{
			MethodName: "MotionEstimateEvaluation",
			Handler:    _VesselMotionService_MotionEstimateEvaluation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vesselMotionService/proto/v1/vessel_motion_service_api_v1.proto",
}
