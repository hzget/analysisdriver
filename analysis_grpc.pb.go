// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package analysisdriver

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

// DataAnalysisClient is the client API for DataAnalysis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataAnalysisClient interface {
	AnalyzeByAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Status, error)
	AnalyzeByPostId(ctx context.Context, in *Id, opts ...grpc.CallOption) (*PostResult, error)
	AnalyzePost(ctx context.Context, in *Text, opts ...grpc.CallOption) (*PostResult, error)
}

type dataAnalysisClient struct {
	cc grpc.ClientConnInterface
}

func NewDataAnalysisClient(cc grpc.ClientConnInterface) DataAnalysisClient {
	return &dataAnalysisClient{cc}
}

func (c *dataAnalysisClient) AnalyzeByAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/analysis.DataAnalysis/AnalyzeByAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAnalysisClient) AnalyzeByPostId(ctx context.Context, in *Id, opts ...grpc.CallOption) (*PostResult, error) {
	out := new(PostResult)
	err := c.cc.Invoke(ctx, "/analysis.DataAnalysis/AnalyzeByPostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataAnalysisClient) AnalyzePost(ctx context.Context, in *Text, opts ...grpc.CallOption) (*PostResult, error) {
	out := new(PostResult)
	err := c.cc.Invoke(ctx, "/analysis.DataAnalysis/AnalyzePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataAnalysisServer is the server API for DataAnalysis service.
// All implementations must embed UnimplementedDataAnalysisServer
// for forward compatibility
type DataAnalysisServer interface {
	AnalyzeByAuthor(context.Context, *Author) (*Status, error)
	AnalyzeByPostId(context.Context, *Id) (*PostResult, error)
	AnalyzePost(context.Context, *Text) (*PostResult, error)
	mustEmbedUnimplementedDataAnalysisServer()
}

// UnimplementedDataAnalysisServer must be embedded to have forward compatible implementations.
type UnimplementedDataAnalysisServer struct {
}

func (UnimplementedDataAnalysisServer) AnalyzeByAuthor(context.Context, *Author) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalyzeByAuthor not implemented")
}
func (UnimplementedDataAnalysisServer) AnalyzeByPostId(context.Context, *Id) (*PostResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalyzeByPostId not implemented")
}
func (UnimplementedDataAnalysisServer) AnalyzePost(context.Context, *Text) (*PostResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalyzePost not implemented")
}
func (UnimplementedDataAnalysisServer) mustEmbedUnimplementedDataAnalysisServer() {}

// UnsafeDataAnalysisServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataAnalysisServer will
// result in compilation errors.
type UnsafeDataAnalysisServer interface {
	mustEmbedUnimplementedDataAnalysisServer()
}

func RegisterDataAnalysisServer(s grpc.ServiceRegistrar, srv DataAnalysisServer) {
	s.RegisterService(&DataAnalysis_ServiceDesc, srv)
}

func _DataAnalysis_AnalyzeByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAnalysisServer).AnalyzeByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analysis.DataAnalysis/AnalyzeByAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAnalysisServer).AnalyzeByAuthor(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAnalysis_AnalyzeByPostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAnalysisServer).AnalyzeByPostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analysis.DataAnalysis/AnalyzeByPostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAnalysisServer).AnalyzeByPostId(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataAnalysis_AnalyzePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataAnalysisServer).AnalyzePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analysis.DataAnalysis/AnalyzePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataAnalysisServer).AnalyzePost(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

// DataAnalysis_ServiceDesc is the grpc.ServiceDesc for DataAnalysis service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataAnalysis_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analysis.DataAnalysis",
	HandlerType: (*DataAnalysisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnalyzeByAuthor",
			Handler:    _DataAnalysis_AnalyzeByAuthor_Handler,
		},
		{
			MethodName: "AnalyzeByPostId",
			Handler:    _DataAnalysis_AnalyzeByPostId_Handler,
		},
		{
			MethodName: "AnalyzePost",
			Handler:    _DataAnalysis_AnalyzePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analysis.proto",
}
