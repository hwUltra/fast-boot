package logic

import (
	"context"
	"fast-boot/app/rpc/model"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictOptionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictOptionsLogic {
	return &DictOptionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictOptionsLogic) DictOptions(in *sysPb.TypeReq) (*sysPb.DictOptionsResp, error) {
	dictModel := model.SysDictModel{}

	var items []*sysPb.DictOption

	l.svcCtx.GormConn.Model(dictModel).
		Select("value", "name as label ").
		Where("type_code = ?", in.TypeCode).
		Order("id asc ").
		Scan(&items)

	return &sysPb.DictOptionsResp{Items: items}, nil

}
