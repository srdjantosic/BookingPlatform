// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: user_reservation_service.proto

package user_reservation

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
	UserReservationService_SayHi_FullMethodName                   = "/user_reservation.UserReservationService/SayHi"
	UserReservationService_GetReservationByGuestId_FullMethodName = "/user_reservation.UserReservationService/GetReservationByGuestId"
)

// UserReservationServiceClient is the client API for UserReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserReservationServiceClient interface {
	SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiResponse, error)
	GetReservationByGuestId(ctx context.Context, in *GetReservationRequest, opts ...grpc.CallOption) (*GetReservationResponse, error)
}

type userReservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserReservationServiceClient(cc grpc.ClientConnInterface) UserReservationServiceClient {
	return &userReservationServiceClient{cc}
}

func (c *userReservationServiceClient) SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiResponse, error) {
	out := new(HiResponse)
	err := c.cc.Invoke(ctx, UserReservationService_SayHi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userReservationServiceClient) GetReservationByGuestId(ctx context.Context, in *GetReservationRequest, opts ...grpc.CallOption) (*GetReservationResponse, error) {
	out := new(GetReservationResponse)
	err := c.cc.Invoke(ctx, UserReservationService_GetReservationByGuestId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserReservationServiceServer is the server API for UserReservationService service.
// All implementations must embed UnimplementedUserReservationServiceServer
// for forward compatibility
type UserReservationServiceServer interface {
	SayHi(context.Context, *HiRequest) (*HiResponse, error)
	GetReservationByGuestId(context.Context, *GetReservationRequest) (*GetReservationResponse, error)
	mustEmbedUnimplementedUserReservationServiceServer()
}

// UnimplementedUserReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserReservationServiceServer struct {
}

func (UnimplementedUserReservationServiceServer) SayHi(context.Context, *HiRequest) (*HiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (UnimplementedUserReservationServiceServer) GetReservationByGuestId(context.Context, *GetReservationRequest) (*GetReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationByGuestId not implemented")
}
func (UnimplementedUserReservationServiceServer) mustEmbedUnimplementedUserReservationServiceServer() {
}

// UnsafeUserReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserReservationServiceServer will
// result in compilation errors.
type UnsafeUserReservationServiceServer interface {
	mustEmbedUnimplementedUserReservationServiceServer()
}

func RegisterUserReservationServiceServer(s grpc.ServiceRegistrar, srv UserReservationServiceServer) {
	s.RegisterService(&UserReservationService_ServiceDesc, srv)
}

func _UserReservationService_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserReservationServiceServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserReservationService_SayHi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserReservationServiceServer).SayHi(ctx, req.(*HiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserReservationService_GetReservationByGuestId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserReservationServiceServer).GetReservationByGuestId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserReservationService_GetReservationByGuestId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserReservationServiceServer).GetReservationByGuestId(ctx, req.(*GetReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserReservationService_ServiceDesc is the grpc.ServiceDesc for UserReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_reservation.UserReservationService",
	HandlerType: (*UserReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _UserReservationService_SayHi_Handler,
		},
		{
			MethodName: "GetReservationByGuestId",
			Handler:    _UserReservationService_GetReservationByGuestId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_reservation_service.proto",
}
