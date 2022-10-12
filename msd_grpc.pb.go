// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: msd.proto

package msd

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

// ApiV1Client is the client API for ApiV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiV1Client interface {
	// 获取数据请求
	Get(ctx context.Context, in *GetDataFrameRequest, opts ...grpc.CallOption) (*GetDataFrameResponse, error)
	// 获取某个表中符合条件的 Object 信息
	ListObject(ctx context.Context, in *ListObjectRequest, opts ...grpc.CallOption) (*ListObjectResponse, error)
	// 更新数据请求
	Update(ctx context.Context, opts ...grpc.CallOption) (ApiV1_UpdateClient, error)
	// 单次更新数据请求，在某些场景，stream 更新可能无法使用，如 grpcweb
	UpdateOnce(ctx context.Context, in *UpdateDataFrameRequest, opts ...grpc.CallOption) (*UpdateDataFrameResponse, error)
	// 获取表信息请求
	GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*GetTableResponse, error)
	// 更新元数据
	UpdateMeta(ctx context.Context, in *UpdateMetaRequest, opts ...grpc.CallOption) (*UpdateMetaResponse, error)
	// 转发相关的操作
	Forward(ctx context.Context, in *ForwardRequest, opts ...grpc.CallOption) (*ForwardResponse, error)
	// 执行SQL
	SqlQuery(ctx context.Context, in *SqlRequest, opts ...grpc.CallOption) (*SqlResponse, error)
}

type apiV1Client struct {
	cc grpc.ClientConnInterface
}

func NewApiV1Client(cc grpc.ClientConnInterface) ApiV1Client {
	return &apiV1Client{cc}
}

func (c *apiV1Client) Get(ctx context.Context, in *GetDataFrameRequest, opts ...grpc.CallOption) (*GetDataFrameResponse, error) {
	out := new(GetDataFrameResponse)
	err := c.cc.Invoke(ctx, "/msd.ApiV1/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiV1Client) ListObject(ctx context.Context, in *ListObjectRequest, opts ...grpc.CallOption) (*ListObjectResponse, error) {
	out := new(ListObjectResponse)
	err := c.cc.Invoke(ctx, "/msd.ApiV1/ListObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiV1Client) Update(ctx context.Context, opts ...grpc.CallOption) (ApiV1_UpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &ApiV1_ServiceDesc.Streams[0], "/msd.ApiV1/Update", opts...)
	if err != nil {
		return nil, err
	}
	x := &apiV1UpdateClient{stream}
	return x, nil
}

type ApiV1_UpdateClient interface {
	Send(*UpdateDataFrameRequest) error
	CloseAndRecv() (*UpdateDataFrameResponse, error)
	grpc.ClientStream
}

type apiV1UpdateClient struct {
	grpc.ClientStream
}

func (x *apiV1UpdateClient) Send(m *UpdateDataFrameRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *apiV1UpdateClient) CloseAndRecv() (*UpdateDataFrameResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpdateDataFrameResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *apiV1Client) UpdateOnce(ctx context.Context, in *UpdateDataFrameRequest, opts ...grpc.CallOption) (*UpdateDataFrameResponse, error) {
	out := new(UpdateDataFrameResponse)
	err := c.cc.Invoke(ctx, "/msd.ApiV1/UpdateOnce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiV1Client) GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*GetTableResponse, error) {
	out := new(GetTableResponse)
	err := c.cc.Invoke(ctx, "/msd.ApiV1/GetTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiV1Client) UpdateMeta(ctx context.Context, in *UpdateMetaRequest, opts ...grpc.CallOption) (*UpdateMetaResponse, error) {
	out := new(UpdateMetaResponse)
	err := c.cc.Invoke(ctx, "/msd.ApiV1/UpdateMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiV1Client) Forward(ctx context.Context, in *ForwardRequest, opts ...grpc.CallOption) (*ForwardResponse, error) {
	out := new(ForwardResponse)
	err := c.cc.Invoke(ctx, "/msd.ApiV1/Forward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiV1Client) SqlQuery(ctx context.Context, in *SqlRequest, opts ...grpc.CallOption) (*SqlResponse, error) {
	out := new(SqlResponse)
	err := c.cc.Invoke(ctx, "/msd.ApiV1/SqlQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiV1Server is the server API for ApiV1 service.
// All implementations must embed UnimplementedApiV1Server
// for forward compatibility
type ApiV1Server interface {
	// 获取数据请求
	Get(context.Context, *GetDataFrameRequest) (*GetDataFrameResponse, error)
	// 获取某个表中符合条件的 Object 信息
	ListObject(context.Context, *ListObjectRequest) (*ListObjectResponse, error)
	// 更新数据请求
	Update(ApiV1_UpdateServer) error
	// 单次更新数据请求，在某些场景，stream 更新可能无法使用，如 grpcweb
	UpdateOnce(context.Context, *UpdateDataFrameRequest) (*UpdateDataFrameResponse, error)
	// 获取表信息请求
	GetTable(context.Context, *GetTableRequest) (*GetTableResponse, error)
	// 更新元数据
	UpdateMeta(context.Context, *UpdateMetaRequest) (*UpdateMetaResponse, error)
	// 转发相关的操作
	Forward(context.Context, *ForwardRequest) (*ForwardResponse, error)
	// 执行SQL
	SqlQuery(context.Context, *SqlRequest) (*SqlResponse, error)
	mustEmbedUnimplementedApiV1Server()
}

// UnimplementedApiV1Server must be embedded to have forward compatible implementations.
type UnimplementedApiV1Server struct {
}

func (UnimplementedApiV1Server) Get(context.Context, *GetDataFrameRequest) (*GetDataFrameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedApiV1Server) ListObject(context.Context, *ListObjectRequest) (*ListObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObject not implemented")
}
func (UnimplementedApiV1Server) Update(ApiV1_UpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedApiV1Server) UpdateOnce(context.Context, *UpdateDataFrameRequest) (*UpdateDataFrameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOnce not implemented")
}
func (UnimplementedApiV1Server) GetTable(context.Context, *GetTableRequest) (*GetTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTable not implemented")
}
func (UnimplementedApiV1Server) UpdateMeta(context.Context, *UpdateMetaRequest) (*UpdateMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMeta not implemented")
}
func (UnimplementedApiV1Server) Forward(context.Context, *ForwardRequest) (*ForwardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Forward not implemented")
}
func (UnimplementedApiV1Server) SqlQuery(context.Context, *SqlRequest) (*SqlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SqlQuery not implemented")
}
func (UnimplementedApiV1Server) mustEmbedUnimplementedApiV1Server() {}

// UnsafeApiV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiV1Server will
// result in compilation errors.
type UnsafeApiV1Server interface {
	mustEmbedUnimplementedApiV1Server()
}

func RegisterApiV1Server(s grpc.ServiceRegistrar, srv ApiV1Server) {
	s.RegisterService(&ApiV1_ServiceDesc, srv)
}

func _ApiV1_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiV1Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msd.ApiV1/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiV1Server).Get(ctx, req.(*GetDataFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiV1_ListObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiV1Server).ListObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msd.ApiV1/ListObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiV1Server).ListObject(ctx, req.(*ListObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiV1_Update_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ApiV1Server).Update(&apiV1UpdateServer{stream})
}

type ApiV1_UpdateServer interface {
	SendAndClose(*UpdateDataFrameResponse) error
	Recv() (*UpdateDataFrameRequest, error)
	grpc.ServerStream
}

type apiV1UpdateServer struct {
	grpc.ServerStream
}

func (x *apiV1UpdateServer) SendAndClose(m *UpdateDataFrameResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *apiV1UpdateServer) Recv() (*UpdateDataFrameRequest, error) {
	m := new(UpdateDataFrameRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ApiV1_UpdateOnce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDataFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiV1Server).UpdateOnce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msd.ApiV1/UpdateOnce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiV1Server).UpdateOnce(ctx, req.(*UpdateDataFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiV1_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiV1Server).GetTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msd.ApiV1/GetTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiV1Server).GetTable(ctx, req.(*GetTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiV1_UpdateMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiV1Server).UpdateMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msd.ApiV1/UpdateMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiV1Server).UpdateMeta(ctx, req.(*UpdateMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiV1_Forward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiV1Server).Forward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msd.ApiV1/Forward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiV1Server).Forward(ctx, req.(*ForwardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiV1_SqlQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiV1Server).SqlQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msd.ApiV1/SqlQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiV1Server).SqlQuery(ctx, req.(*SqlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiV1_ServiceDesc is the grpc.ServiceDesc for ApiV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msd.ApiV1",
	HandlerType: (*ApiV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ApiV1_Get_Handler,
		},
		{
			MethodName: "ListObject",
			Handler:    _ApiV1_ListObject_Handler,
		},
		{
			MethodName: "UpdateOnce",
			Handler:    _ApiV1_UpdateOnce_Handler,
		},
		{
			MethodName: "GetTable",
			Handler:    _ApiV1_GetTable_Handler,
		},
		{
			MethodName: "UpdateMeta",
			Handler:    _ApiV1_UpdateMeta_Handler,
		},
		{
			MethodName: "Forward",
			Handler:    _ApiV1_Forward_Handler,
		},
		{
			MethodName: "SqlQuery",
			Handler:    _ApiV1_SqlQuery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Update",
			Handler:       _ApiV1_Update_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "msd.proto",
}
