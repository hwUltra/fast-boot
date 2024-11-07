package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsCategoryAttributeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsCategoryAttributeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsCategoryAttributeAddLogic {
	return &PmsCategoryAttributeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsCategoryAttributeAddLogic) PmsCategoryAttributeAdd(in *pmsPb.PmsCategoryAttributeAddReq) (*pmsPb.SuccessResp, error) {

	if err := l.svcCtx.GormClient.GormDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("type = ?", in.Type).
			Where("category_id = ?", in.CategoryId).
			Delete(&model.PmsCategoryAttributeModel{}).Error; err != nil {
			return err
		}

		for _, name := range in.Attributes {
			if err := tx.FirstOrCreate(&model.PmsCategoryAttributeModel{}, &model.PmsCategoryAttributeModel{
				Name:       name,
				Type:       int8(in.Type),
				CategoryId: in.CategoryId,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pmsPb.SuccessResp{}, nil
}
