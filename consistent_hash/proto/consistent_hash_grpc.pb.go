// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/consistent_hash.proto

package consistent_hash

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
	ConsistentHashServer_GetValue_FullMethodName     = "/consistent_hash.ConsistentHashServer/GetValue"
	ConsistentHashServer_SetKV_FullMethodName        = "/consistent_hash.ConsistentHashServer/SetKV"
	ConsistentHashServer_DeleteKey_FullMethodName    = "/consistent_hash.ConsistentHashServer/DeleteKey"
	ConsistentHashServer_ListenGossip_FullMethodName = "/consistent_hash.ConsistentHashServer/ListenGossip"
)

// ConsistentHashServerClient is the client API for ConsistentHashServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsistentHashServerClient interface {
	GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error)
	SetKV(ctx context.Context, in *SetKVRequest, opts ...grpc.CallOption) (*SetKVResponse, error)
	DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*DeleteKeyResponse, error)
	ListenGossip(ctx context.Context, in *Gossip, opts ...grpc.CallOption) (*Gossip, error)
}

type consistentHashServerClient struct {
	cc grpc.ClientConnInterface
}

func NewConsistentHashServerClient(cc grpc.ClientConnInterface) ConsistentHashServerClient {
	return &consistentHashServerClient{cc}
}

func (c *consistentHashServerClient) GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetValueResponse)
	err := c.cc.Invoke(ctx, ConsistentHashServer_GetValue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consistentHashServerClient) SetKV(ctx context.Context, in *SetKVRequest, opts ...grpc.CallOption) (*SetKVResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetKVResponse)
	err := c.cc.Invoke(ctx, ConsistentHashServer_SetKV_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consistentHashServerClient) DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*DeleteKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteKeyResponse)
	err := c.cc.Invoke(ctx, ConsistentHashServer_DeleteKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consistentHashServerClient) ListenGossip(ctx context.Context, in *Gossip, opts ...grpc.CallOption) (*Gossip, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Gossip)
	err := c.cc.Invoke(ctx, ConsistentHashServer_ListenGossip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsistentHashServerServer is the server API for ConsistentHashServer service.
// All implementations must embed UnimplementedConsistentHashServerServer
// for forward compatibility.
type ConsistentHashServerServer interface {
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)
	SetKV(context.Context, *SetKVRequest) (*SetKVResponse, error)
	DeleteKey(context.Context, *DeleteKeyRequest) (*DeleteKeyResponse, error)
	ListenGossip(context.Context, *Gossip) (*Gossip, error)
	mustEmbedUnimplementedConsistentHashServerServer()
}

// UnimplementedConsistentHashServerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConsistentHashServerServer struct{}

func (UnimplementedConsistentHashServerServer) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedConsistentHashServerServer) SetKV(context.Context, *SetKVRequest) (*SetKVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetKV not implemented")
}
func (UnimplementedConsistentHashServerServer) DeleteKey(context.Context, *DeleteKeyRequest) (*DeleteKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKey not implemented")
}
func (UnimplementedConsistentHashServerServer) ListenGossip(context.Context, *Gossip) (*Gossip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListenGossip not implemented")
}
func (UnimplementedConsistentHashServerServer) mustEmbedUnimplementedConsistentHashServerServer() {}
func (UnimplementedConsistentHashServerServer) testEmbeddedByValue()                              {}

// UnsafeConsistentHashServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsistentHashServerServer will
// result in compilation errors.
type UnsafeConsistentHashServerServer interface {
	mustEmbedUnimplementedConsistentHashServerServer()
}

func RegisterConsistentHashServerServer(s grpc.ServiceRegistrar, srv ConsistentHashServerServer) {
	// If the following call pancis, it indicates UnimplementedConsistentHashServerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ConsistentHashServer_ServiceDesc, srv)
}

func _ConsistentHashServer_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsistentHashServerServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsistentHashServer_GetValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsistentHashServerServer).GetValue(ctx, req.(*GetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsistentHashServer_SetKV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetKVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsistentHashServerServer).SetKV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsistentHashServer_SetKV_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsistentHashServerServer).SetKV(ctx, req.(*SetKVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsistentHashServer_DeleteKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsistentHashServerServer).DeleteKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsistentHashServer_DeleteKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsistentHashServerServer).DeleteKey(ctx, req.(*DeleteKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsistentHashServer_ListenGossip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Gossip)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsistentHashServerServer).ListenGossip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsistentHashServer_ListenGossip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsistentHashServerServer).ListenGossip(ctx, req.(*Gossip))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsistentHashServer_ServiceDesc is the grpc.ServiceDesc for ConsistentHashServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsistentHashServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "consistent_hash.ConsistentHashServer",
	HandlerType: (*ConsistentHashServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValue",
			Handler:    _ConsistentHashServer_GetValue_Handler,
		},
		{
			MethodName: "SetKV",
			Handler:    _ConsistentHashServer_SetKV_Handler,
		},
		{
			MethodName: "DeleteKey",
			Handler:    _ConsistentHashServer_DeleteKey_Handler,
		},
		{
			MethodName: "ListenGossip",
			Handler:    _ConsistentHashServer_ListenGossip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consistent_hash.proto",
}

const (
	ProxyService_GetValue_FullMethodName              = "/consistent_hash.ProxyService/GetValue"
	ProxyService_SetKV_FullMethodName                 = "/consistent_hash.ProxyService/SetKV"
	ProxyService_DeleteKey_FullMethodName             = "/consistent_hash.ProxyService/DeleteKey"
	ProxyService_RedirectGossipMessage_FullMethodName = "/consistent_hash.ProxyService/RedirectGossipMessage"
	ProxyService_GetAllServerInfo_FullMethodName      = "/consistent_hash.ProxyService/GetAllServerInfo"
)

// ProxyServiceClient is the client API for ProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyServiceClient interface {
	// those three apis are for client usage
	GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error)
	SetKV(ctx context.Context, in *SetKVRequest, opts ...grpc.CallOption) (*SetKVResponse, error)
	DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*DeleteKeyResponse, error)
	// redirect all kinds of message
	// which could create the network partition manually
	RedirectGossipMessage(ctx context.Context, in *RedirectGossipMessageRequest, opts ...grpc.CallOption) (*RedirectGossipMessageResponse, error)
	GetAllServerInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Gossip, error)
}

type proxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyServiceClient(cc grpc.ClientConnInterface) ProxyServiceClient {
	return &proxyServiceClient{cc}
}

func (c *proxyServiceClient) GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetValueResponse)
	err := c.cc.Invoke(ctx, ProxyService_GetValue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyServiceClient) SetKV(ctx context.Context, in *SetKVRequest, opts ...grpc.CallOption) (*SetKVResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetKVResponse)
	err := c.cc.Invoke(ctx, ProxyService_SetKV_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyServiceClient) DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*DeleteKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteKeyResponse)
	err := c.cc.Invoke(ctx, ProxyService_DeleteKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyServiceClient) RedirectGossipMessage(ctx context.Context, in *RedirectGossipMessageRequest, opts ...grpc.CallOption) (*RedirectGossipMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RedirectGossipMessageResponse)
	err := c.cc.Invoke(ctx, ProxyService_RedirectGossipMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyServiceClient) GetAllServerInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Gossip, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Gossip)
	err := c.cc.Invoke(ctx, ProxyService_GetAllServerInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServiceServer is the server API for ProxyService service.
// All implementations must embed UnimplementedProxyServiceServer
// for forward compatibility.
type ProxyServiceServer interface {
	// those three apis are for client usage
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)
	SetKV(context.Context, *SetKVRequest) (*SetKVResponse, error)
	DeleteKey(context.Context, *DeleteKeyRequest) (*DeleteKeyResponse, error)
	// redirect all kinds of message
	// which could create the network partition manually
	RedirectGossipMessage(context.Context, *RedirectGossipMessageRequest) (*RedirectGossipMessageResponse, error)
	GetAllServerInfo(context.Context, *Empty) (*Gossip, error)
	mustEmbedUnimplementedProxyServiceServer()
}

// UnimplementedProxyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProxyServiceServer struct{}

func (UnimplementedProxyServiceServer) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedProxyServiceServer) SetKV(context.Context, *SetKVRequest) (*SetKVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetKV not implemented")
}
func (UnimplementedProxyServiceServer) DeleteKey(context.Context, *DeleteKeyRequest) (*DeleteKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKey not implemented")
}
func (UnimplementedProxyServiceServer) RedirectGossipMessage(context.Context, *RedirectGossipMessageRequest) (*RedirectGossipMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedirectGossipMessage not implemented")
}
func (UnimplementedProxyServiceServer) GetAllServerInfo(context.Context, *Empty) (*Gossip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllServerInfo not implemented")
}
func (UnimplementedProxyServiceServer) mustEmbedUnimplementedProxyServiceServer() {}
func (UnimplementedProxyServiceServer) testEmbeddedByValue()                      {}

// UnsafeProxyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServiceServer will
// result in compilation errors.
type UnsafeProxyServiceServer interface {
	mustEmbedUnimplementedProxyServiceServer()
}

func RegisterProxyServiceServer(s grpc.ServiceRegistrar, srv ProxyServiceServer) {
	// If the following call pancis, it indicates UnimplementedProxyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProxyService_ServiceDesc, srv)
}

func _ProxyService_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProxyService_GetValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).GetValue(ctx, req.(*GetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyService_SetKV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetKVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).SetKV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProxyService_SetKV_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).SetKV(ctx, req.(*SetKVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyService_DeleteKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).DeleteKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProxyService_DeleteKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).DeleteKey(ctx, req.(*DeleteKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyService_RedirectGossipMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedirectGossipMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).RedirectGossipMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProxyService_RedirectGossipMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).RedirectGossipMessage(ctx, req.(*RedirectGossipMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyService_GetAllServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).GetAllServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProxyService_GetAllServerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).GetAllServerInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ProxyService_ServiceDesc is the grpc.ServiceDesc for ProxyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProxyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "consistent_hash.ProxyService",
	HandlerType: (*ProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValue",
			Handler:    _ProxyService_GetValue_Handler,
		},
		{
			MethodName: "SetKV",
			Handler:    _ProxyService_SetKV_Handler,
		},
		{
			MethodName: "DeleteKey",
			Handler:    _ProxyService_DeleteKey_Handler,
		},
		{
			MethodName: "RedirectGossipMessage",
			Handler:    _ProxyService_RedirectGossipMessage_Handler,
		},
		{
			MethodName: "GetAllServerInfo",
			Handler:    _ProxyService_GetAllServerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consistent_hash.proto",
}
