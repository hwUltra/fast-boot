// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package pms

import (
	"context"

	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AnyReq                       = pmsPb.AnyReq
	IdReq                        = pmsPb.IdReq
	IdResp                       = pmsPb.IdResp
	IdsReq                       = pmsPb.IdsReq
	ListReq                      = pmsPb.ListReq
	OptionItem                   = pmsPb.OptionItem
	OptionsResp                  = pmsPb.OptionsResp
	PmsBrand                     = pmsPb.PmsBrand
	PmsBrandForm                 = pmsPb.PmsBrandForm
	PmsBrandListReq              = pmsPb.PmsBrandListReq
	PmsBrandListResp             = pmsPb.PmsBrandListResp
	PmsCategory                  = pmsPb.PmsCategory
	PmsCategoryAttribute         = pmsPb.PmsCategoryAttribute
	PmsCategoryAttributeAddReq   = pmsPb.PmsCategoryAttributeAddReq
	PmsCategoryAttributeListReq  = pmsPb.PmsCategoryAttributeListReq
	PmsCategoryAttributeListResp = pmsPb.PmsCategoryAttributeListResp
	PmsCategoryForm              = pmsPb.PmsCategoryForm
	PmsCategoryListReq           = pmsPb.PmsCategoryListReq
	PmsCategoryListResp          = pmsPb.PmsCategoryListResp
	PmsGoods                     = pmsPb.PmsGoods
	PmsGoodsAttribute            = pmsPb.PmsGoodsAttribute
	PmsGoodsForm                 = pmsPb.PmsGoodsForm
	PmsGoodsFormAttribute        = pmsPb.PmsGoodsFormAttribute
	PmsGoodsFormSku              = pmsPb.PmsGoodsFormSku
	PmsGoodsInfo                 = pmsPb.PmsGoodsInfo
	PmsGoodsListReq              = pmsPb.PmsGoodsListReq
	PmsGoodsListResp             = pmsPb.PmsGoodsListResp
	PmsSku                       = pmsPb.PmsSku
	RepOptionItem                = pmsPb.RepOptionItem
	RepOptionsResp               = pmsPb.RepOptionsResp
	Shop                         = pmsPb.Shop
	ShopForm                     = pmsPb.ShopForm
	ShopListResp                 = pmsPb.ShopListResp
	SuccessIdResp                = pmsPb.SuccessIdResp
	SuccessResp                  = pmsPb.SuccessResp

	Pms interface {
		ShopGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Shop, error)
		ShopList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ShopListResp, error)
		ShopAdd(ctx context.Context, in *ShopForm, opts ...grpc.CallOption) (*SuccessIdResp, error)
		ShopUpdate(ctx context.Context, in *ShopForm, opts ...grpc.CallOption) (*SuccessResp, error)
		ShopDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		ShopOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*OptionsResp, error)
		PmsCategoryGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsCategory, error)
		PmsCategoryList(ctx context.Context, in *PmsCategoryListReq, opts ...grpc.CallOption) (*PmsCategoryListResp, error)
		PmsCategoryAdd(ctx context.Context, in *PmsCategoryForm, opts ...grpc.CallOption) (*SuccessIdResp, error)
		PmsCategoryUpdate(ctx context.Context, in *PmsCategoryForm, opts ...grpc.CallOption) (*SuccessResp, error)
		PmsCategoryDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		PmsCategoryOptions(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RepOptionsResp, error)
		PmsCategoryAttributeList(ctx context.Context, in *PmsCategoryAttributeListReq, opts ...grpc.CallOption) (*PmsCategoryAttributeListResp, error)
		PmsCategoryAttributeAdd(ctx context.Context, in *PmsCategoryAttributeAddReq, opts ...grpc.CallOption) (*SuccessResp, error)
		PmsBrandGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsBrand, error)
		PmsBrandList(ctx context.Context, in *PmsBrandListReq, opts ...grpc.CallOption) (*PmsBrandListResp, error)
		PmsBrandAdd(ctx context.Context, in *PmsBrandForm, opts ...grpc.CallOption) (*SuccessIdResp, error)
		PmsBrandUpdate(ctx context.Context, in *PmsBrandForm, opts ...grpc.CallOption) (*SuccessResp, error)
		PmsBrandDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		PmsBrandOptions(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*OptionsResp, error)
		PmsGoodsList(ctx context.Context, in *PmsGoodsListReq, opts ...grpc.CallOption) (*PmsGoodsListResp, error)
		PmsGoodsGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsGoodsInfo, error)
		PmsGoodsDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		PmsGoodsEdit(ctx context.Context, in *PmsGoodsForm, opts ...grpc.CallOption) (*SuccessIdResp, error)
	}

	defaultPms struct {
		cli zrpc.Client
	}
)

func NewPms(cli zrpc.Client) Pms {
	return &defaultPms{
		cli: cli,
	}
}

func (m *defaultPms) ShopGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Shop, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.ShopGet(ctx, in, opts...)
}

func (m *defaultPms) ShopList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ShopListResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.ShopList(ctx, in, opts...)
}

func (m *defaultPms) ShopAdd(ctx context.Context, in *ShopForm, opts ...grpc.CallOption) (*SuccessIdResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.ShopAdd(ctx, in, opts...)
}

func (m *defaultPms) ShopUpdate(ctx context.Context, in *ShopForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.ShopUpdate(ctx, in, opts...)
}

func (m *defaultPms) ShopDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.ShopDel(ctx, in, opts...)
}

func (m *defaultPms) ShopOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*OptionsResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.ShopOptions(ctx, in, opts...)
}

func (m *defaultPms) PmsCategoryGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsCategory, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsCategoryGet(ctx, in, opts...)
}

func (m *defaultPms) PmsCategoryList(ctx context.Context, in *PmsCategoryListReq, opts ...grpc.CallOption) (*PmsCategoryListResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsCategoryList(ctx, in, opts...)
}

func (m *defaultPms) PmsCategoryAdd(ctx context.Context, in *PmsCategoryForm, opts ...grpc.CallOption) (*SuccessIdResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsCategoryAdd(ctx, in, opts...)
}

func (m *defaultPms) PmsCategoryUpdate(ctx context.Context, in *PmsCategoryForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsCategoryUpdate(ctx, in, opts...)
}

func (m *defaultPms) PmsCategoryDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsCategoryDel(ctx, in, opts...)
}

func (m *defaultPms) PmsCategoryOptions(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RepOptionsResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsCategoryOptions(ctx, in, opts...)
}

func (m *defaultPms) PmsCategoryAttributeList(ctx context.Context, in *PmsCategoryAttributeListReq, opts ...grpc.CallOption) (*PmsCategoryAttributeListResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsCategoryAttributeList(ctx, in, opts...)
}

func (m *defaultPms) PmsCategoryAttributeAdd(ctx context.Context, in *PmsCategoryAttributeAddReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsCategoryAttributeAdd(ctx, in, opts...)
}

func (m *defaultPms) PmsBrandGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsBrand, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsBrandGet(ctx, in, opts...)
}

func (m *defaultPms) PmsBrandList(ctx context.Context, in *PmsBrandListReq, opts ...grpc.CallOption) (*PmsBrandListResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsBrandList(ctx, in, opts...)
}

func (m *defaultPms) PmsBrandAdd(ctx context.Context, in *PmsBrandForm, opts ...grpc.CallOption) (*SuccessIdResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsBrandAdd(ctx, in, opts...)
}

func (m *defaultPms) PmsBrandUpdate(ctx context.Context, in *PmsBrandForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsBrandUpdate(ctx, in, opts...)
}

func (m *defaultPms) PmsBrandDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsBrandDel(ctx, in, opts...)
}

func (m *defaultPms) PmsBrandOptions(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*OptionsResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsBrandOptions(ctx, in, opts...)
}

func (m *defaultPms) PmsGoodsList(ctx context.Context, in *PmsGoodsListReq, opts ...grpc.CallOption) (*PmsGoodsListResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsGoodsList(ctx, in, opts...)
}

func (m *defaultPms) PmsGoodsGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsGoodsInfo, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsGoodsGet(ctx, in, opts...)
}

func (m *defaultPms) PmsGoodsDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsGoodsDel(ctx, in, opts...)
}

func (m *defaultPms) PmsGoodsEdit(ctx context.Context, in *PmsGoodsForm, opts ...grpc.CallOption) (*SuccessIdResp, error) {
	client := pmsPb.NewPmsClient(m.cli.Conn())
	return client.PmsGoodsEdit(ctx, in, opts...)
}
