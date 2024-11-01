package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"google.golang.org/grpc/status"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsCategoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsCategoryUpdateLogic {
	return &PmsCategoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsCategoryUpdateLogic) PmsCategoryUpdate(in *pmsPb.PmsCategoryForm) (*pmsPb.SuccessResp, error) {
	info := model.PmsCategoryModel{}
	l.svcCtx.GormConn.First(&info, in.Id)

	if info.Id == 0 {
		logx.WithContext(l.ctx).Errorf("不存在: %s", in.Id)
		return nil, status.Error(100, "不存在")
	}
	info.ShopId = in.ShopId
	info.ParentId = in.ParentId
	info.Sort = in.Sort
	info.Name = in.Name
	info.Icon = in.Icon
	info.Sort = in.Sort
	info.Visible = in.Visible
	res := l.svcCtx.GormConn.Save(&info)
	if res.Error != nil {
		return nil, res.Error
	}
	return &pmsPb.SuccessResp{}, nil
}
