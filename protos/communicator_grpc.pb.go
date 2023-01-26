// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/communicator.proto

package communicator

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

// CommunicatorClient is the client API for Communicator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunicatorClient interface {
	Send(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type communicatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicatorClient(cc grpc.ClientConnInterface) CommunicatorClient {
	return &communicatorClient{cc}
}

func (c *communicatorClient) Send(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/communicator.Communicator/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicatorServer is the server API for Communicator service.
// All implementations must embed UnimplementedCommunicatorServer
// for forward compatibility
type CommunicatorServer interface {
	Send(context.Context, *MessageRequest) (*MessageResponse, error)
	mustEmbedUnimplementedCommunicatorServer()
}

// UnimplementedCommunicatorServer must be embedded to have forward compatible implementations.
type UnimplementedCommunicatorServer struct {
}

func (UnimplementedCommunicatorServer) Send(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedCommunicatorServer) mustEmbedUnimplementedCommunicatorServer() {}

// UnsafeCommunicatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunicatorServer will
// result in compilation errors.
type UnsafeCommunicatorServer interface {
	mustEmbedUnimplementedCommunicatorServer()
}

func RegisterCommunicatorServer(s grpc.ServiceRegistrar, srv CommunicatorServer) {
	s.RegisterService(&Communicator_ServiceDesc, srv)
}

func _Communicator_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicatorServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communicator.Communicator/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicatorServer).Send(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Communicator_ServiceDesc is the grpc.ServiceDesc for Communicator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Communicator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "communicator.Communicator",
	HandlerType: (*CommunicatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Communicator_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/communicator.proto",
}
