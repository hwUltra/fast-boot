package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsCategoryGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsCategoryGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsCategoryGetLogic {
	return &PmsCategoryGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsCategoryGetLogic) PmsCategoryGet(in *pmsPb.IdReq) (*pmsPb.PmsCategory, error) {
	item := model.PmsCategoryModel{}
	l.svcCtx.GormConn.First(&item, in.Id)
	var info pmsPb.PmsCategory
	if err := copier.Copy(&info, item); err != nil {
		return nil, err
	}
	return &info, nil
}
