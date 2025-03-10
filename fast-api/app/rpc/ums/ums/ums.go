// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: ums.proto

package ums

import (
	"context"

	"fast-boot/app/rpc/ums/umsPb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BindReq             = umsPb.BindReq
	IdReq               = umsPb.IdReq
	IdResp              = umsPb.IdResp
	IdsReq              = umsPb.IdsReq
	OpenIdReq           = umsPb.OpenIdReq
	PageReq             = umsPb.PageReq
	SuccessIdResp       = umsPb.SuccessIdResp
	SuccessResp         = umsPb.SuccessResp
	User                = umsPb.User
	UserAddress         = umsPb.UserAddress
	UserAddressListResp = umsPb.UserAddressListResp
	UserCarts           = umsPb.UserCarts
	UserFeedback        = umsPb.UserFeedback
	UserForm            = umsPb.UserForm
	UserInfoResp        = umsPb.UserInfoResp
	UserPageResp        = umsPb.UserPageResp
	UserThird           = umsPb.UserThird
	UserThirdPageResp   = umsPb.UserThirdPageResp

	Ums interface {
		// UserUpdate
		ByOpenId(ctx context.Context, in *OpenIdReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		Bind(ctx context.Context, in *BindReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		// -------
		UserGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*User, error)
		UserPage(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*UserPageResp, error)
		UserAdd(ctx context.Context, in *UserForm, opts ...grpc.CallOption) (*SuccessIdResp, error)
		UserUpdate(ctx context.Context, in *UserForm, opts ...grpc.CallOption) (*SuccessResp, error)
		UserDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		// third
		UserThirdPage(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*UserThirdPageResp, error)
	}

	defaultUms struct {
		cli zrpc.Client
	}
)

func NewUms(cli zrpc.Client) Ums {
	return &defaultUms{
		cli: cli,
	}
}

// UserUpdate
func (m *defaultUms) ByOpenId(ctx context.Context, in *OpenIdReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := umsPb.NewUmsClient(m.cli.Conn())
	return client.ByOpenId(ctx, in, opts...)
}

func (m *defaultUms) Bind(ctx context.Context, in *BindReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := umsPb.NewUmsClient(m.cli.Conn())
	return client.Bind(ctx, in, opts...)
}

// -------
func (m *defaultUms) UserGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*User, error) {
	client := umsPb.NewUmsClient(m.cli.Conn())
	return client.UserGet(ctx, in, opts...)
}

func (m *defaultUms) UserPage(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*UserPageResp, error) {
	client := umsPb.NewUmsClient(m.cli.Conn())
	return client.UserPage(ctx, in, opts...)
}

func (m *defaultUms) UserAdd(ctx context.Context, in *UserForm, opts ...grpc.CallOption) (*SuccessIdResp, error) {
	client := umsPb.NewUmsClient(m.cli.Conn())
	return client.UserAdd(ctx, in, opts...)
}

func (m *defaultUms) UserUpdate(ctx context.Context, in *UserForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := umsPb.NewUmsClient(m.cli.Conn())
	return client.UserUpdate(ctx, in, opts...)
}

func (m *defaultUms) UserDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := umsPb.NewUmsClient(m.cli.Conn())
	return client.UserDel(ctx, in, opts...)
}

// third
func (m *defaultUms) UserThirdPage(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*UserThirdPageResp, error) {
	client := umsPb.NewUmsClient(m.cli.Conn())
	return client.UserThirdPage(ctx, in, opts...)
}
