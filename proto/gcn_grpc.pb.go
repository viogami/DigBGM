// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: gcn.proto

package proto

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
	GCNService_ProcessGraph_FullMethodName = "/GCNService/ProcessGraph"
)

// GCNServiceClient is the client API for GCNService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GCNServiceClient interface {
	ProcessGraph(ctx context.Context, in *GraphData, opts ...grpc.CallOption) (*GCNResult, error)
}

type gCNServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGCNServiceClient(cc grpc.ClientConnInterface) GCNServiceClient {
	return &gCNServiceClient{cc}
}

func (c *gCNServiceClient) ProcessGraph(ctx context.Context, in *GraphData, opts ...grpc.CallOption) (*GCNResult, error) {
	out := new(GCNResult)
	err := c.cc.Invoke(ctx, GCNService_ProcessGraph_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GCNServiceServer is the server API for GCNService service.
// All implementations must embed UnimplementedGCNServiceServer
// for forward compatibility
type GCNServiceServer interface {
	ProcessGraph(context.Context, *GraphData) (*GCNResult, error)
	mustEmbedUnimplementedGCNServiceServer()
}

// UnimplementedGCNServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGCNServiceServer struct {
}

func (UnimplementedGCNServiceServer) ProcessGraph(context.Context, *GraphData) (*GCNResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessGraph not implemented")
}
func (UnimplementedGCNServiceServer) mustEmbedUnimplementedGCNServiceServer() {}

// UnsafeGCNServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GCNServiceServer will
// result in compilation errors.
type UnsafeGCNServiceServer interface {
	mustEmbedUnimplementedGCNServiceServer()
}

func RegisterGCNServiceServer(s grpc.ServiceRegistrar, srv GCNServiceServer) {
	s.RegisterService(&GCNService_ServiceDesc, srv)
}

func _GCNService_ProcessGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GCNServiceServer).ProcessGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GCNService_ProcessGraph_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GCNServiceServer).ProcessGraph(ctx, req.(*GraphData))
	}
	return interceptor(ctx, in, info, handler)
}

// GCNService_ServiceDesc is the grpc.ServiceDesc for GCNService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GCNService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GCNService",
	HandlerType: (*GCNServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessGraph",
			Handler:    _GCNService_ProcessGraph_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gcn.proto",
}
