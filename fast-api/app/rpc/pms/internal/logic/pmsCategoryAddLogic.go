package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsCategoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsCategoryAddLogic {
	return &PmsCategoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsCategoryAddLogic) PmsCategoryAdd(in *pmsPb.PmsCategoryForm) (*pmsPb.SuccessIdResp, error) {
	item := model.PmsCategoryModel{}
	if err := copier.Copy(&item, in); err != nil {
		return nil, err
	}
	l.svcCtx.GormClient.GormDb.Create(&item)
	return &pmsPb.SuccessIdResp{Id: item.Id}, nil
}
