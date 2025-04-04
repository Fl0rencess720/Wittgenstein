// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: gateway/seminar/v1/seminar.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Seminar_CreateTopic_FullMethodName       = "/Wittgenstein.v1.Seminar/CreateTopic"
	Seminar_GetTopicsMetadata_FullMethodName = "/Wittgenstein.v1.Seminar/GetTopicsMetadata"
	Seminar_GetTopic_FullMethodName          = "/Wittgenstein.v1.Seminar/GetTopic"
	Seminar_DeleteTopic_FullMethodName       = "/Wittgenstein.v1.Seminar/DeleteTopic"
	Seminar_StartTopic_FullMethodName        = "/Wittgenstein.v1.Seminar/StartTopic"
	Seminar_StopTopic_FullMethodName         = "/Wittgenstein.v1.Seminar/StopTopic"
	Seminar_ResumeTopic_FullMethodName       = "/Wittgenstein.v1.Seminar/ResumeTopic"
)

// SeminarClient is the client API for Seminar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeminarClient interface {
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicReply, error)
	// 获取用户所有讨论主题的元信息，用于前端展示
	GetTopicsMetadata(ctx context.Context, in *GetTopicsMetadataRequest, opts ...grpc.CallOption) (*GetTopicsMetadataReply, error)
	// 获取讨论主题的详细信息，进入讨论时加载
	GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*GetTopicReply, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicReply, error)
	StartTopic(ctx context.Context, in *StartTopicRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamOutputReply], error)
	StopTopic(ctx context.Context, in *StopTopicRequest, opts ...grpc.CallOption) (*StopTopicReply, error)
	ResumeTopic(ctx context.Context, in *StartTopicRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamOutputReply], error)
}

type seminarClient struct {
	cc grpc.ClientConnInterface
}

func NewSeminarClient(cc grpc.ClientConnInterface) SeminarClient {
	return &seminarClient{cc}
}

func (c *seminarClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTopicReply)
	err := c.cc.Invoke(ctx, Seminar_CreateTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seminarClient) GetTopicsMetadata(ctx context.Context, in *GetTopicsMetadataRequest, opts ...grpc.CallOption) (*GetTopicsMetadataReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTopicsMetadataReply)
	err := c.cc.Invoke(ctx, Seminar_GetTopicsMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seminarClient) GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*GetTopicReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTopicReply)
	err := c.cc.Invoke(ctx, Seminar_GetTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seminarClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*DeleteTopicReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTopicReply)
	err := c.cc.Invoke(ctx, Seminar_DeleteTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seminarClient) StartTopic(ctx context.Context, in *StartTopicRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamOutputReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Seminar_ServiceDesc.Streams[0], Seminar_StartTopic_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StartTopicRequest, StreamOutputReply]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Seminar_StartTopicClient = grpc.ServerStreamingClient[StreamOutputReply]

func (c *seminarClient) StopTopic(ctx context.Context, in *StopTopicRequest, opts ...grpc.CallOption) (*StopTopicReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopTopicReply)
	err := c.cc.Invoke(ctx, Seminar_StopTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seminarClient) ResumeTopic(ctx context.Context, in *StartTopicRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamOutputReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Seminar_ServiceDesc.Streams[1], Seminar_ResumeTopic_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StartTopicRequest, StreamOutputReply]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Seminar_ResumeTopicClient = grpc.ServerStreamingClient[StreamOutputReply]

// SeminarServer is the server API for Seminar service.
// All implementations must embed UnimplementedSeminarServer
// for forward compatibility.
type SeminarServer interface {
	CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicReply, error)
	// 获取用户所有讨论主题的元信息，用于前端展示
	GetTopicsMetadata(context.Context, *GetTopicsMetadataRequest) (*GetTopicsMetadataReply, error)
	// 获取讨论主题的详细信息，进入讨论时加载
	GetTopic(context.Context, *GetTopicRequest) (*GetTopicReply, error)
	DeleteTopic(context.Context, *DeleteTopicRequest) (*DeleteTopicReply, error)
	StartTopic(*StartTopicRequest, grpc.ServerStreamingServer[StreamOutputReply]) error
	StopTopic(context.Context, *StopTopicRequest) (*StopTopicReply, error)
	ResumeTopic(*StartTopicRequest, grpc.ServerStreamingServer[StreamOutputReply]) error
	mustEmbedUnimplementedSeminarServer()
}

// UnimplementedSeminarServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSeminarServer struct{}

func (UnimplementedSeminarServer) CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedSeminarServer) GetTopicsMetadata(context.Context, *GetTopicsMetadataRequest) (*GetTopicsMetadataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicsMetadata not implemented")
}
func (UnimplementedSeminarServer) GetTopic(context.Context, *GetTopicRequest) (*GetTopicReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}
func (UnimplementedSeminarServer) DeleteTopic(context.Context, *DeleteTopicRequest) (*DeleteTopicReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedSeminarServer) StartTopic(*StartTopicRequest, grpc.ServerStreamingServer[StreamOutputReply]) error {
	return status.Errorf(codes.Unimplemented, "method StartTopic not implemented")
}
func (UnimplementedSeminarServer) StopTopic(context.Context, *StopTopicRequest) (*StopTopicReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopTopic not implemented")
}
func (UnimplementedSeminarServer) ResumeTopic(*StartTopicRequest, grpc.ServerStreamingServer[StreamOutputReply]) error {
	return status.Errorf(codes.Unimplemented, "method ResumeTopic not implemented")
}
func (UnimplementedSeminarServer) mustEmbedUnimplementedSeminarServer() {}
func (UnimplementedSeminarServer) testEmbeddedByValue()                 {}

// UnsafeSeminarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeminarServer will
// result in compilation errors.
type UnsafeSeminarServer interface {
	mustEmbedUnimplementedSeminarServer()
}

func RegisterSeminarServer(s grpc.ServiceRegistrar, srv SeminarServer) {
	// If the following call pancis, it indicates UnimplementedSeminarServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Seminar_ServiceDesc, srv)
}

func _Seminar_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeminarServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seminar_CreateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeminarServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seminar_GetTopicsMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicsMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeminarServer).GetTopicsMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seminar_GetTopicsMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeminarServer).GetTopicsMetadata(ctx, req.(*GetTopicsMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seminar_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeminarServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seminar_GetTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeminarServer).GetTopic(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seminar_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeminarServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seminar_DeleteTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeminarServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seminar_StartTopic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StartTopicRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SeminarServer).StartTopic(m, &grpc.GenericServerStream[StartTopicRequest, StreamOutputReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Seminar_StartTopicServer = grpc.ServerStreamingServer[StreamOutputReply]

func _Seminar_StopTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeminarServer).StopTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seminar_StopTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeminarServer).StopTopic(ctx, req.(*StopTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seminar_ResumeTopic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StartTopicRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SeminarServer).ResumeTopic(m, &grpc.GenericServerStream[StartTopicRequest, StreamOutputReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Seminar_ResumeTopicServer = grpc.ServerStreamingServer[StreamOutputReply]

// Seminar_ServiceDesc is the grpc.ServiceDesc for Seminar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Seminar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Wittgenstein.v1.Seminar",
	HandlerType: (*SeminarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _Seminar_CreateTopic_Handler,
		},
		{
			MethodName: "GetTopicsMetadata",
			Handler:    _Seminar_GetTopicsMetadata_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _Seminar_GetTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _Seminar_DeleteTopic_Handler,
		},
		{
			MethodName: "StopTopic",
			Handler:    _Seminar_StopTopic_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartTopic",
			Handler:       _Seminar_StartTopic_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ResumeTopic",
			Handler:       _Seminar_ResumeTopic_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gateway/seminar/v1/seminar.proto",
}
