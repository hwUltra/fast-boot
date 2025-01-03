// Code generated by goctl. DO NOT EDIT.
// Source: oms.proto

package oms

import (
	"context"

	"fast-boot/app/rpc/oms/omsPb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AnyReq         = omsPb.AnyReq
	IdReq          = omsPb.IdReq
	IdResp         = omsPb.IdResp
	IdsReq         = omsPb.IdsReq
	ListReq        = omsPb.ListReq
	Order          = omsPb.Order
	OrderAddReq    = omsPb.OrderAddReq
	OrderListResp  = omsPb.OrderListResp
	OrderUpdateReq = omsPb.OrderUpdateReq
	SuccessIdResp  = omsPb.SuccessIdResp
	SuccessResp    = omsPb.SuccessResp

	Oms interface {
		// -------
		OrderGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Order, error)
		OrderList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*OrderListResp, error)
		OrderAdd(ctx context.Context, in *OrderAddReq, opts ...grpc.CallOption) (*SuccessIdResp, error)
		OrderUpdate(ctx context.Context, in *OrderUpdateReq, opts ...grpc.CallOption) (*SuccessResp, error)
	}

	defaultOms struct {
		cli zrpc.Client
	}
)

func NewOms(cli zrpc.Client) Oms {
	return &defaultOms{
		cli: cli,
	}
}

// -------
func (m *defaultOms) OrderGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Order, error) {
	client := omsPb.NewOmsClient(m.cli.Conn())
	return client.OrderGet(ctx, in, opts...)
}

func (m *defaultOms) OrderList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*OrderListResp, error) {
	client := omsPb.NewOmsClient(m.cli.Conn())
	return client.OrderList(ctx, in, opts...)
}

func (m *defaultOms) OrderAdd(ctx context.Context, in *OrderAddReq, opts ...grpc.CallOption) (*SuccessIdResp, error) {
	client := omsPb.NewOmsClient(m.cli.Conn())
	return client.OrderAdd(ctx, in, opts...)
}

func (m *defaultOms) OrderUpdate(ctx context.Context, in *OrderUpdateReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := omsPb.NewOmsClient(m.cli.Conn())
	return client.OrderUpdate(ctx, in, opts...)
}
