package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"google.golang.org/grpc/status"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShopUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopUpdateLogic {
	return &ShopUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShopUpdateLogic) ShopUpdate(in *pmsPb.ShopForm) (*pmsPb.SuccessResp, error) {
	info := model.PmsShopModel{}
	l.svcCtx.GormConn.First(&info, in.Id)

	if info.Id == 0 {
		logx.WithContext(l.ctx).Errorf("不存在: %s", in.Id)
		return nil, status.Error(100, "不存在")
	}
	if len(info.Name) > 0 && info.Name != in.Name {
		info.Name = in.Name
	}

	if len(info.Tel) > 0 && info.Tel != in.Tel {
		info.Tel = in.Tel
	}

	if len(info.Notice) > 0 && info.Notice != in.Notice {
		info.Notice = in.Notice
	}

	if info.DistributionFee != in.DistributionFee {
		info.DistributionFee = in.DistributionFee
	}

	res := l.svcCtx.GormConn.Save(&info)
	if res.Error != nil {
		return nil, res.Error
	}
	return &pmsPb.SuccessResp{}, nil
}
