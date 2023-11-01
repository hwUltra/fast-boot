package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsBrandGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsBrandGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsBrandGetLogic {
	return &PmsBrandGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsBrandGetLogic) PmsBrandGet(in *pmsPb.IdReq) (*pmsPb.PmsBrand, error) {
	item := model.PmsBrandModel{}
	l.svcCtx.GormConn.First(&item, in.Id)
	var info pmsPb.PmsBrand
	if err := copier.Copy(&info, item); err != nil {
		return nil, err
	}
	info.CreatedAt = item.CreatedAt.String()
	return &info, nil
}
