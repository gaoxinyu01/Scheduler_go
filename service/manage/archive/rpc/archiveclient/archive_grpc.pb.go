// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: archive.proto

package archiveclient

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
	Archive_AppLoggerAdd_FullMethodName      = "/archiveclient.Archive/AppLoggerAdd"
	Archive_AppLoggerFindList_FullMethodName = "/archiveclient.Archive/AppLoggerFindList"
)

// ArchiveClient is the client API for Archive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArchiveClient interface {
	// 用户日志
	AppLoggerAdd(ctx context.Context, in *AppLoggerAddReq, opts ...grpc.CallOption) (*CommonResp, error)
	AppLoggerFindList(ctx context.Context, in *AppLoggerFindListReq, opts ...grpc.CallOption) (*AppLoggerFindListResp, error)
}

type archiveClient struct {
	cc grpc.ClientConnInterface
}

func NewArchiveClient(cc grpc.ClientConnInterface) ArchiveClient {
	return &archiveClient{cc}
}

func (c *archiveClient) AppLoggerAdd(ctx context.Context, in *AppLoggerAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Archive_AppLoggerAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archiveClient) AppLoggerFindList(ctx context.Context, in *AppLoggerFindListReq, opts ...grpc.CallOption) (*AppLoggerFindListResp, error) {
	out := new(AppLoggerFindListResp)
	err := c.cc.Invoke(ctx, Archive_AppLoggerFindList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchiveServer is the server API for Archive service.
// All implementations must embed UnimplementedArchiveServer
// for forward compatibility
type ArchiveServer interface {
	// 用户日志
	AppLoggerAdd(context.Context, *AppLoggerAddReq) (*CommonResp, error)
	AppLoggerFindList(context.Context, *AppLoggerFindListReq) (*AppLoggerFindListResp, error)
	mustEmbedUnimplementedArchiveServer()
}

// UnimplementedArchiveServer must be embedded to have forward compatible implementations.
type UnimplementedArchiveServer struct {
}

func (UnimplementedArchiveServer) AppLoggerAdd(context.Context, *AppLoggerAddReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppLoggerAdd not implemented")
}
func (UnimplementedArchiveServer) AppLoggerFindList(context.Context, *AppLoggerFindListReq) (*AppLoggerFindListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppLoggerFindList not implemented")
}
func (UnimplementedArchiveServer) mustEmbedUnimplementedArchiveServer() {}

// UnsafeArchiveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArchiveServer will
// result in compilation errors.
type UnsafeArchiveServer interface {
	mustEmbedUnimplementedArchiveServer()
}

func RegisterArchiveServer(s grpc.ServiceRegistrar, srv ArchiveServer) {
	s.RegisterService(&Archive_ServiceDesc, srv)
}

func _Archive_AppLoggerAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppLoggerAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveServer).AppLoggerAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Archive_AppLoggerAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveServer).AppLoggerAdd(ctx, req.(*AppLoggerAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Archive_AppLoggerFindList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppLoggerFindListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveServer).AppLoggerFindList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Archive_AppLoggerFindList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveServer).AppLoggerFindList(ctx, req.(*AppLoggerFindListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Archive_ServiceDesc is the grpc.ServiceDesc for Archive service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Archive_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "archiveclient.Archive",
	HandlerType: (*ArchiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppLoggerAdd",
			Handler:    _Archive_AppLoggerAdd_Handler,
		},
		{
			MethodName: "AppLoggerFindList",
			Handler:    _Archive_AppLoggerFindList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "archive.proto",
}
