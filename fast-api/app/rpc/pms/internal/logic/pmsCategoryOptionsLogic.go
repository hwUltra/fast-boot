package logic

import (
	"context"
	"fast-boot/app/rpc/model"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsCategoryOptionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsCategoryOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsCategoryOptionsLogic {
	return &PmsCategoryOptionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsCategoryOptionsLogic) PmsCategoryOptions(in *pmsPb.IdReq) (*pmsPb.RepOptionsResp, error) {
	var items []*model.PmsCategoryModel
	l.svcCtx.GormClient.GormDb.Model(model.PmsCategoryModel{}).
		Where("shop_id = ?", in.Id).
		Order("sort desc ,id asc ").
		Find(&items)

	return &pmsPb.RepOptionsResp{Items: makeTree(items, 0)}, nil
}

func makeTree(list []*model.PmsCategoryModel, parentID int64) []*pmsPb.RepOptionItem {
	res := make([]*pmsPb.RepOptionItem, 0)
	for _, item := range list {
		if item.ParentId == parentID {
			var ot pmsPb.RepOptionItem
			ot.Label = item.Name
			ot.Value = item.Id
			ot.Children = makeTree(list, item.Id)
			res = append(res, &ot)
		}
	}
	return res
}
