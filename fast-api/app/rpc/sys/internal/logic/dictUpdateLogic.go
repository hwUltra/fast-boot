package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"google.golang.org/grpc/status"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictUpdateLogic {
	return &DictUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictUpdateLogic) DictUpdate(in *sysPb.DictForm) (*sysPb.SuccessResp, error) {
	info := model.SysDictModel{}
	l.svcCtx.GormConn.First(&info, in.Id)

	if info.Id == 0 {
		logx.WithContext(l.ctx).Errorf("不存在: %s", in.Id)
		return nil, status.Error(100, "不存在")
	}
	info.Code = in.Code
	info.Name = in.Name
	info.Remark = in.Remark
	info.Status = int8(in.Status)

	res := l.svcCtx.GormConn.Save(&info)
	if res.Error != nil {
		return nil, res.Error
	}
	return &sysPb.SuccessResp{}, nil
}
