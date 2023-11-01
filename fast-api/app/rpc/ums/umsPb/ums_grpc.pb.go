// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: pb/ums.proto

package umsPb

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

// UmsClient is the client API for Ums service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UmsClient interface {
	// UserUpdate
	ByOpenId(ctx context.Context, in *OpenIdReq, opts ...grpc.CallOption) (*UserInfoResp, error)
	Bind(ctx context.Context, in *BindReq, opts ...grpc.CallOption) (*UserInfoResp, error)
	// -------
	UserGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*User, error)
	UserList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*UserListResp, error)
	UserAdd(ctx context.Context, in *UserForm, opts ...grpc.CallOption) (*SuccessIdResp, error)
	UserUpdate(ctx context.Context, in *UserForm, opts ...grpc.CallOption) (*SuccessResp, error)
	UserDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
	// third
	UserThirdList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*UserThirdListResp, error)
}

type umsClient struct {
	cc grpc.ClientConnInterface
}

func NewUmsClient(cc grpc.ClientConnInterface) UmsClient {
	return &umsClient{cc}
}

func (c *umsClient) ByOpenId(ctx context.Context, in *OpenIdReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, "/umsPb.ums/ByOpenId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) Bind(ctx context.Context, in *BindReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, "/umsPb.ums/Bind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) UserGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/umsPb.ums/UserGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) UserList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	out := new(UserListResp)
	err := c.cc.Invoke(ctx, "/umsPb.ums/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) UserAdd(ctx context.Context, in *UserForm, opts ...grpc.CallOption) (*SuccessIdResp, error) {
	out := new(SuccessIdResp)
	err := c.cc.Invoke(ctx, "/umsPb.ums/UserAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) UserUpdate(ctx context.Context, in *UserForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/umsPb.ums/UserUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) UserDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/umsPb.ums/UserDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) UserThirdList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*UserThirdListResp, error) {
	out := new(UserThirdListResp)
	err := c.cc.Invoke(ctx, "/umsPb.ums/UserThirdList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UmsServer is the server API for Ums service.
// All implementations must embed UnimplementedUmsServer
// for forward compatibility
type UmsServer interface {
	// UserUpdate
	ByOpenId(context.Context, *OpenIdReq) (*UserInfoResp, error)
	Bind(context.Context, *BindReq) (*UserInfoResp, error)
	// -------
	UserGet(context.Context, *IdReq) (*User, error)
	UserList(context.Context, *ListReq) (*UserListResp, error)
	UserAdd(context.Context, *UserForm) (*SuccessIdResp, error)
	UserUpdate(context.Context, *UserForm) (*SuccessResp, error)
	UserDel(context.Context, *IdsReq) (*SuccessResp, error)
	// third
	UserThirdList(context.Context, *ListReq) (*UserThirdListResp, error)
	mustEmbedUnimplementedUmsServer()
}

// UnimplementedUmsServer must be embedded to have forward compatible implementations.
type UnimplementedUmsServer struct {
}

func (UnimplementedUmsServer) ByOpenId(context.Context, *OpenIdReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByOpenId not implemented")
}
func (UnimplementedUmsServer) Bind(context.Context, *BindReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bind not implemented")
}
func (UnimplementedUmsServer) UserGet(context.Context, *IdReq) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGet not implemented")
}
func (UnimplementedUmsServer) UserList(context.Context, *ListReq) (*UserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedUmsServer) UserAdd(context.Context, *UserForm) (*SuccessIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}
func (UnimplementedUmsServer) UserUpdate(context.Context, *UserForm) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
func (UnimplementedUmsServer) UserDel(context.Context, *IdsReq) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDel not implemented")
}
func (UnimplementedUmsServer) UserThirdList(context.Context, *ListReq) (*UserThirdListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserThirdList not implemented")
}
func (UnimplementedUmsServer) mustEmbedUnimplementedUmsServer() {}

// UnsafeUmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UmsServer will
// result in compilation errors.
type UnsafeUmsServer interface {
	mustEmbedUnimplementedUmsServer()
}

func RegisterUmsServer(s grpc.ServiceRegistrar, srv UmsServer) {
	s.RegisterService(&Ums_ServiceDesc, srv)
}

func _Ums_ByOpenId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).ByOpenId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umsPb.ums/ByOpenId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).ByOpenId(ctx, req.(*OpenIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_Bind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).Bind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umsPb.ums/Bind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).Bind(ctx, req.(*BindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_UserGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).UserGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umsPb.ums/UserGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).UserGet(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umsPb.ums/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).UserList(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umsPb.ums/UserAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).UserAdd(ctx, req.(*UserForm))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umsPb.ums/UserUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).UserUpdate(ctx, req.(*UserForm))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_UserDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).UserDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umsPb.ums/UserDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).UserDel(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_UserThirdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).UserThirdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umsPb.ums/UserThirdList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).UserThirdList(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Ums_ServiceDesc is the grpc.ServiceDesc for Ums service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ums_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "umsPb.ums",
	HandlerType: (*UmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByOpenId",
			Handler:    _Ums_ByOpenId_Handler,
		},
		{
			MethodName: "Bind",
			Handler:    _Ums_Bind_Handler,
		},
		{
			MethodName: "UserGet",
			Handler:    _Ums_UserGet_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _Ums_UserList_Handler,
		},
		{
			MethodName: "UserAdd",
			Handler:    _Ums_UserAdd_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _Ums_UserUpdate_Handler,
		},
		{
			MethodName: "UserDel",
			Handler:    _Ums_UserDel_Handler,
		},
		{
			MethodName: "UserThirdList",
			Handler:    _Ums_UserThirdList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/ums.proto",
}
