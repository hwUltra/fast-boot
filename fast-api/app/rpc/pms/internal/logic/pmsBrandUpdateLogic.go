package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"google.golang.org/grpc/status"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsBrandUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsBrandUpdateLogic {
	return &PmsBrandUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsBrandUpdateLogic) PmsBrandUpdate(in *pmsPb.PmsBrandForm) (*pmsPb.SuccessResp, error) {
	info := model.PmsBrandModel{}
	l.svcCtx.GormConn.First(&info, in.Id)

	if info.Id == 0 {
		logx.WithContext(l.ctx).Errorf("不存在: %s", in.Id)
		return nil, status.Error(100, "不存在")
	}
	info.ShopId = in.ShopId
	info.Name = in.Name
	info.Logo = in.Logo
	info.Sort = in.Sort
	res := l.svcCtx.GormConn.Save(&info)
	if res.Error != nil {
		return nil, res.Error
	}
	return &pmsPb.SuccessResp{}, nil
}
