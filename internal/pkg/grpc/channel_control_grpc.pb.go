// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: channel_control.proto

package grpc

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
	ChannelControl_SetValue_FullMethodName = "/channelControl.v1.ChannelControl/SetValue"
)

// ChannelControlClient is the client API for ChannelControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelControlClient interface {
	SetValue(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ControlValue, ControlValue], error)
}

type channelControlClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelControlClient(cc grpc.ClientConnInterface) ChannelControlClient {
	return &channelControlClient{cc}
}

func (c *channelControlClient) SetValue(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ControlValue, ControlValue], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChannelControl_ServiceDesc.Streams[0], ChannelControl_SetValue_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ControlValue, ControlValue]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChannelControl_SetValueClient = grpc.BidiStreamingClient[ControlValue, ControlValue]

// ChannelControlServer is the server API for ChannelControl service.
// All implementations must embed UnimplementedChannelControlServer
// for forward compatibility.
type ChannelControlServer interface {
	SetValue(grpc.BidiStreamingServer[ControlValue, ControlValue]) error
	mustEmbedUnimplementedChannelControlServer()
}

// UnimplementedChannelControlServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChannelControlServer struct{}

func (UnimplementedChannelControlServer) SetValue(grpc.BidiStreamingServer[ControlValue, ControlValue]) error {
	return status.Errorf(codes.Unimplemented, "method SetValue not implemented")
}
func (UnimplementedChannelControlServer) mustEmbedUnimplementedChannelControlServer() {}
func (UnimplementedChannelControlServer) testEmbeddedByValue()                        {}

// UnsafeChannelControlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelControlServer will
// result in compilation errors.
type UnsafeChannelControlServer interface {
	mustEmbedUnimplementedChannelControlServer()
}

func RegisterChannelControlServer(s grpc.ServiceRegistrar, srv ChannelControlServer) {
	// If the following call pancis, it indicates UnimplementedChannelControlServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChannelControl_ServiceDesc, srv)
}

func _ChannelControl_SetValue_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChannelControlServer).SetValue(&grpc.GenericServerStream[ControlValue, ControlValue]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChannelControl_SetValueServer = grpc.BidiStreamingServer[ControlValue, ControlValue]

// ChannelControl_ServiceDesc is the grpc.ServiceDesc for ChannelControl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelControl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channelControl.v1.ChannelControl",
	HandlerType: (*ChannelControlServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SetValue",
			Handler:       _ChannelControl_SetValue_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "channel_control.proto",
}
