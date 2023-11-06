// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: grpc/tradeapi/v1/candles.proto

package tradeapi

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
	Candles_GetDayCandles_FullMethodName      = "/grpc.tradeapi.v1.Candles/GetDayCandles"
	Candles_GetIntradayCandles_FullMethodName = "/grpc.tradeapi.v1.Candles/GetIntradayCandles"
)

// CandlesClient is the client API for Candles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CandlesClient interface {
	// Returns the list of day candles.
	// Возвращает дневные свечи.
	GetDayCandles(ctx context.Context, in *GetDayCandlesRequest, opts ...grpc.CallOption) (*GetDayCandlesResult, error)
	// Returns the list of intraday candles.
	// Возвращает внутридневные свечи.
	GetIntradayCandles(ctx context.Context, in *GetIntradayCandlesRequest, opts ...grpc.CallOption) (*GetIntradayCandlesResult, error)
}

type candlesClient struct {
	cc grpc.ClientConnInterface
}

func NewCandlesClient(cc grpc.ClientConnInterface) CandlesClient {
	return &candlesClient{cc}
}

func (c *candlesClient) GetDayCandles(ctx context.Context, in *GetDayCandlesRequest, opts ...grpc.CallOption) (*GetDayCandlesResult, error) {
	out := new(GetDayCandlesResult)
	err := c.cc.Invoke(ctx, Candles_GetDayCandles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *candlesClient) GetIntradayCandles(ctx context.Context, in *GetIntradayCandlesRequest, opts ...grpc.CallOption) (*GetIntradayCandlesResult, error) {
	out := new(GetIntradayCandlesResult)
	err := c.cc.Invoke(ctx, Candles_GetIntradayCandles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CandlesServer is the server API for Candles service.
// All implementations must embed UnimplementedCandlesServer
// for forward compatibility
type CandlesServer interface {
	// Returns the list of day candles.
	// Возвращает дневные свечи.
	GetDayCandles(context.Context, *GetDayCandlesRequest) (*GetDayCandlesResult, error)
	// Returns the list of intraday candles.
	// Возвращает внутридневные свечи.
	GetIntradayCandles(context.Context, *GetIntradayCandlesRequest) (*GetIntradayCandlesResult, error)
	mustEmbedUnimplementedCandlesServer()
}

// UnimplementedCandlesServer must be embedded to have forward compatible implementations.
type UnimplementedCandlesServer struct {
}

func (UnimplementedCandlesServer) GetDayCandles(context.Context, *GetDayCandlesRequest) (*GetDayCandlesResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDayCandles not implemented")
}
func (UnimplementedCandlesServer) GetIntradayCandles(context.Context, *GetIntradayCandlesRequest) (*GetIntradayCandlesResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntradayCandles not implemented")
}
func (UnimplementedCandlesServer) mustEmbedUnimplementedCandlesServer() {}

// UnsafeCandlesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CandlesServer will
// result in compilation errors.
type UnsafeCandlesServer interface {
	mustEmbedUnimplementedCandlesServer()
}

func RegisterCandlesServer(s grpc.ServiceRegistrar, srv CandlesServer) {
	s.RegisterService(&Candles_ServiceDesc, srv)
}

func _Candles_GetDayCandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDayCandlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandlesServer).GetDayCandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Candles_GetDayCandles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandlesServer).GetDayCandles(ctx, req.(*GetDayCandlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Candles_GetIntradayCandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntradayCandlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandlesServer).GetIntradayCandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Candles_GetIntradayCandles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandlesServer).GetIntradayCandles(ctx, req.(*GetIntradayCandlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Candles_ServiceDesc is the grpc.ServiceDesc for Candles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Candles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.tradeapi.v1.Candles",
	HandlerType: (*CandlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDayCandles",
			Handler:    _Candles_GetDayCandles_Handler,
		},
		{
			MethodName: "GetIntradayCandles",
			Handler:    _Candles_GetIntradayCandles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/tradeapi/v1/candles.proto",
}
