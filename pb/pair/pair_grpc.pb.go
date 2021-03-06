// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pair.proto

package pair

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

// PairClient is the client API for Pair service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PairClient interface {
	ElementSave(ctx context.Context, in *EleSaveReq, opts ...grpc.CallOption) (*Response, error)
	ElementView(ctx context.Context, in *EleViewReq, opts ...grpc.CallOption) (*EleViewRsp, error)
	ElementList(ctx context.Context, in *EleListReq, opts ...grpc.CallOption) (*ELeListRsp, error)
	ElementPair(ctx context.Context, in *ElePairReq, opts ...grpc.CallOption) (*ELeListRsp, error)
}

type pairClient struct {
	cc grpc.ClientConnInterface
}

func NewPairClient(cc grpc.ClientConnInterface) PairClient {
	return &pairClient{cc}
}

func (c *pairClient) ElementSave(ctx context.Context, in *EleSaveReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pair.Pair/ElementSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairClient) ElementView(ctx context.Context, in *EleViewReq, opts ...grpc.CallOption) (*EleViewRsp, error) {
	out := new(EleViewRsp)
	err := c.cc.Invoke(ctx, "/pair.Pair/ElementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairClient) ElementList(ctx context.Context, in *EleListReq, opts ...grpc.CallOption) (*ELeListRsp, error) {
	out := new(ELeListRsp)
	err := c.cc.Invoke(ctx, "/pair.Pair/ElementList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairClient) ElementPair(ctx context.Context, in *ElePairReq, opts ...grpc.CallOption) (*ELeListRsp, error) {
	out := new(ELeListRsp)
	err := c.cc.Invoke(ctx, "/pair.Pair/ElementPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PairServer is the server API for Pair service.
// All implementations must embed UnimplementedPairServer
// for forward compatibility
type PairServer interface {
	ElementSave(context.Context, *EleSaveReq) (*Response, error)
	ElementView(context.Context, *EleViewReq) (*EleViewRsp, error)
	ElementList(context.Context, *EleListReq) (*ELeListRsp, error)
	ElementPair(context.Context, *ElePairReq) (*ELeListRsp, error)
	mustEmbedUnimplementedPairServer()
}

// UnimplementedPairServer must be embedded to have forward compatible implementations.
type UnimplementedPairServer struct {
}

func (UnimplementedPairServer) ElementSave(context.Context, *EleSaveReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ElementSave not implemented")
}
func (UnimplementedPairServer) ElementView(context.Context, *EleViewReq) (*EleViewRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ElementView not implemented")
}
func (UnimplementedPairServer) ElementList(context.Context, *EleListReq) (*ELeListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ElementList not implemented")
}
func (UnimplementedPairServer) ElementPair(context.Context, *ElePairReq) (*ELeListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ElementPair not implemented")
}
func (UnimplementedPairServer) mustEmbedUnimplementedPairServer() {}

// UnsafePairServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PairServer will
// result in compilation errors.
type UnsafePairServer interface {
	mustEmbedUnimplementedPairServer()
}

func RegisterPairServer(s grpc.ServiceRegistrar, srv PairServer) {
	s.RegisterService(&Pair_ServiceDesc, srv)
}

func _Pair_ElementSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EleSaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairServer).ElementSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pair.Pair/ElementSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairServer).ElementSave(ctx, req.(*EleSaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pair_ElementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EleViewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairServer).ElementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pair.Pair/ElementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairServer).ElementView(ctx, req.(*EleViewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pair_ElementList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairServer).ElementList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pair.Pair/ElementList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairServer).ElementList(ctx, req.(*EleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pair_ElementPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElePairReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairServer).ElementPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pair.Pair/ElementPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairServer).ElementPair(ctx, req.(*ElePairReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Pair_ServiceDesc is the grpc.ServiceDesc for Pair service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pair_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pair.Pair",
	HandlerType: (*PairServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ElementSave",
			Handler:    _Pair_ElementSave_Handler,
		},
		{
			MethodName: "ElementView",
			Handler:    _Pair_ElementView_Handler,
		},
		{
			MethodName: "ElementList",
			Handler:    _Pair_ElementList_Handler,
		},
		{
			MethodName: "ElementPair",
			Handler:    _Pair_ElementPair_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pair.proto",
}
