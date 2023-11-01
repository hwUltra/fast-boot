package logic

import (
	"context"
	"fast-boot/app/rpc/model"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictTypeOptionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictTypeOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeOptionsLogic {
	return &DictTypeOptionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictTypeOptionsLogic) DictTypeOptions(in *sysPb.AnyReq) (*sysPb.DictTypeOptionsResp, error) {
	dictModel := model.SysDictTypeModel{}

	var items []*sysPb.DictTypeOption

	l.svcCtx.GormConn.Model(dictModel).
		Select("`code` as `value`", "`name` as `label` ").
		Where("`status` = 1").
		Order("`id` asc ").
		Scan(&items)

	return &sysPb.DictTypeOptionsResp{Items: items}, nil

}
