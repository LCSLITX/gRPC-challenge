// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: crypto.proto

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptoServiceClient interface {
	CreateCrypto(ctx context.Context, in *Crypto, opts ...grpc.CallOption) (*CryptoId, error)
	ReadCrypto(ctx context.Context, in *CryptoId, opts ...grpc.CallOption) (*Crypto, error)
	UpdateCrypto(ctx context.Context, in *CryptoId, opts ...grpc.CallOption) (*Crypto, error)
	DeleteCrypto(ctx context.Context, in *CryptoId, opts ...grpc.CallOption) (*empty.Empty, error)
	ListCryptos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (CryptoService_ListCryptosClient, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) CreateCrypto(ctx context.Context, in *Crypto, opts ...grpc.CallOption) (*CryptoId, error) {
	out := new(CryptoId)
	err := c.cc.Invoke(ctx, "/crypto.CryptoService/CreateCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) ReadCrypto(ctx context.Context, in *CryptoId, opts ...grpc.CallOption) (*Crypto, error) {
	out := new(Crypto)
	err := c.cc.Invoke(ctx, "/crypto.CryptoService/ReadCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) UpdateCrypto(ctx context.Context, in *CryptoId, opts ...grpc.CallOption) (*Crypto, error) {
	out := new(Crypto)
	err := c.cc.Invoke(ctx, "/crypto.CryptoService/UpdateCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) DeleteCrypto(ctx context.Context, in *CryptoId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/crypto.CryptoService/DeleteCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) ListCryptos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (CryptoService_ListCryptosClient, error) {
	stream, err := c.cc.NewStream(ctx, &CryptoService_ServiceDesc.Streams[0], "/crypto.CryptoService/ListCryptos", opts...)
	if err != nil {
		return nil, err
	}
	x := &cryptoServiceListCryptosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CryptoService_ListCryptosClient interface {
	Recv() (*Crypto, error)
	grpc.ClientStream
}

type cryptoServiceListCryptosClient struct {
	grpc.ClientStream
}

func (x *cryptoServiceListCryptosClient) Recv() (*Crypto, error) {
	m := new(Crypto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CryptoServiceServer is the server API for CryptoService service.
// All implementations must embed UnimplementedCryptoServiceServer
// for forward compatibility
type CryptoServiceServer interface {
	CreateCrypto(context.Context, *Crypto) (*CryptoId, error)
	ReadCrypto(context.Context, *CryptoId) (*Crypto, error)
	UpdateCrypto(context.Context, *CryptoId) (*Crypto, error)
	DeleteCrypto(context.Context, *CryptoId) (*empty.Empty, error)
	ListCryptos(*empty.Empty, CryptoService_ListCryptosServer) error
	mustEmbedUnimplementedCryptoServiceServer()
}

// UnimplementedCryptoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCryptoServiceServer struct {
}

func (UnimplementedCryptoServiceServer) CreateCrypto(context.Context, *Crypto) (*CryptoId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCrypto not implemented")
}
func (UnimplementedCryptoServiceServer) ReadCrypto(context.Context, *CryptoId) (*Crypto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCrypto not implemented")
}
func (UnimplementedCryptoServiceServer) UpdateCrypto(context.Context, *CryptoId) (*Crypto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCrypto not implemented")
}
func (UnimplementedCryptoServiceServer) DeleteCrypto(context.Context, *CryptoId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCrypto not implemented")
}
func (UnimplementedCryptoServiceServer) ListCryptos(*empty.Empty, CryptoService_ListCryptosServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCryptos not implemented")
}
func (UnimplementedCryptoServiceServer) mustEmbedUnimplementedCryptoServiceServer() {}

// UnsafeCryptoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoServiceServer will
// result in compilation errors.
type UnsafeCryptoServiceServer interface {
	mustEmbedUnimplementedCryptoServiceServer()
}

func RegisterCryptoServiceServer(s grpc.ServiceRegistrar, srv CryptoServiceServer) {
	s.RegisterService(&CryptoService_ServiceDesc, srv)
}

func _CryptoService_CreateCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Crypto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CreateCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoService/CreateCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CreateCrypto(ctx, req.(*Crypto))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_ReadCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).ReadCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoService/ReadCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).ReadCrypto(ctx, req.(*CryptoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_UpdateCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).UpdateCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoService/UpdateCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).UpdateCrypto(ctx, req.(*CryptoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_DeleteCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).DeleteCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoService/DeleteCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).DeleteCrypto(ctx, req.(*CryptoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_ListCryptos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CryptoServiceServer).ListCryptos(m, &cryptoServiceListCryptosServer{stream})
}

type CryptoService_ListCryptosServer interface {
	Send(*Crypto) error
	grpc.ServerStream
}

type cryptoServiceListCryptosServer struct {
	grpc.ServerStream
}

func (x *cryptoServiceListCryptosServer) Send(m *Crypto) error {
	return x.ServerStream.SendMsg(m)
}

// CryptoService_ServiceDesc is the grpc.ServiceDesc for CryptoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crypto.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCrypto",
			Handler:    _CryptoService_CreateCrypto_Handler,
		},
		{
			MethodName: "ReadCrypto",
			Handler:    _CryptoService_ReadCrypto_Handler,
		},
		{
			MethodName: "UpdateCrypto",
			Handler:    _CryptoService_UpdateCrypto_Handler,
		},
		{
			MethodName: "DeleteCrypto",
			Handler:    _CryptoService_DeleteCrypto_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCryptos",
			Handler:       _CryptoService_ListCryptos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "crypto.proto",
}
