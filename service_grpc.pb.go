// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package billing

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

// BillingClient is the client API for Billing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillingClient interface {
	AddOrUpdateProductsPrices(ctx context.Context, in *AddOrUpdatePrices, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteProductsPrices(ctx context.Context, in *DeletePrices, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPrices(ctx context.Context, in *QueryPrices, opts ...grpc.CallOption) (*Prices, error)
	CalculateTotalPrice(ctx context.Context, in *CalculatePrice, opts ...grpc.CallOption) (*TotalPrice, error)
}

type billingClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingClient(cc grpc.ClientConnInterface) BillingClient {
	return &billingClient{cc}
}

func (c *billingClient) AddOrUpdateProductsPrices(ctx context.Context, in *AddOrUpdatePrices, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/billing.Billing/AddOrUpdateProductsPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) DeleteProductsPrices(ctx context.Context, in *DeletePrices, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/billing.Billing/DeleteProductsPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) GetPrices(ctx context.Context, in *QueryPrices, opts ...grpc.CallOption) (*Prices, error) {
	out := new(Prices)
	err := c.cc.Invoke(ctx, "/billing.Billing/GetPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) CalculateTotalPrice(ctx context.Context, in *CalculatePrice, opts ...grpc.CallOption) (*TotalPrice, error) {
	out := new(TotalPrice)
	err := c.cc.Invoke(ctx, "/billing.Billing/CalculateTotalPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingServer is the server API for Billing service.
// All implementations must embed UnimplementedBillingServer
// for forward compatibility
type BillingServer interface {
	AddOrUpdateProductsPrices(context.Context, *AddOrUpdatePrices) (*emptypb.Empty, error)
	DeleteProductsPrices(context.Context, *DeletePrices) (*emptypb.Empty, error)
	GetPrices(context.Context, *QueryPrices) (*Prices, error)
	CalculateTotalPrice(context.Context, *CalculatePrice) (*TotalPrice, error)
	mustEmbedUnimplementedBillingServer()
}

// UnimplementedBillingServer must be embedded to have forward compatible implementations.
type UnimplementedBillingServer struct {
}

func (UnimplementedBillingServer) AddOrUpdateProductsPrices(context.Context, *AddOrUpdatePrices) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateProductsPrices not implemented")
}
func (UnimplementedBillingServer) DeleteProductsPrices(context.Context, *DeletePrices) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductsPrices not implemented")
}
func (UnimplementedBillingServer) GetPrices(context.Context, *QueryPrices) (*Prices, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrices not implemented")
}
func (UnimplementedBillingServer) CalculateTotalPrice(context.Context, *CalculatePrice) (*TotalPrice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateTotalPrice not implemented")
}
func (UnimplementedBillingServer) mustEmbedUnimplementedBillingServer() {}

// UnsafeBillingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillingServer will
// result in compilation errors.
type UnsafeBillingServer interface {
	mustEmbedUnimplementedBillingServer()
}

func RegisterBillingServer(s grpc.ServiceRegistrar, srv BillingServer) {
	s.RegisterService(&Billing_ServiceDesc, srv)
}

func _Billing_AddOrUpdateProductsPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdatePrices)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).AddOrUpdateProductsPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Billing/AddOrUpdateProductsPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).AddOrUpdateProductsPrices(ctx, req.(*AddOrUpdatePrices))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_DeleteProductsPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePrices)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).DeleteProductsPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Billing/DeleteProductsPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).DeleteProductsPrices(ctx, req.(*DeletePrices))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_GetPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPrices)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).GetPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Billing/GetPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).GetPrices(ctx, req.(*QueryPrices))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_CalculateTotalPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatePrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).CalculateTotalPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Billing/CalculateTotalPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).CalculateTotalPrice(ctx, req.(*CalculatePrice))
	}
	return interceptor(ctx, in, info, handler)
}

// Billing_ServiceDesc is the grpc.ServiceDesc for Billing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Billing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "billing.Billing",
	HandlerType: (*BillingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrUpdateProductsPrices",
			Handler:    _Billing_AddOrUpdateProductsPrices_Handler,
		},
		{
			MethodName: "DeleteProductsPrices",
			Handler:    _Billing_DeleteProductsPrices_Handler,
		},
		{
			MethodName: "GetPrices",
			Handler:    _Billing_GetPrices_Handler,
		},
		{
			MethodName: "CalculateTotalPrice",
			Handler:    _Billing_CalculateTotalPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}