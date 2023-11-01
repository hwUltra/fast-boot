package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsBrandAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsBrandAddLogic {
	return &PmsBrandAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsBrandAddLogic) PmsBrandAdd(in *pmsPb.PmsBrandForm) (*pmsPb.SuccessIdResp, error) {
	item := &model.PmsBrandModel{}
	if err := copier.Copy(&item, in); err != nil {
		return nil, err
	}
	l.svcCtx.GormConn.Create(&item)
	return &pmsPb.SuccessIdResp{Id: item.Id}, nil
}
