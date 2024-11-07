package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"strings"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptDelLogic {
	return &DeptDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptDelLogic) DeptDel(in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {
	ids := strings.Split(in.Ids, ",")
	l.svcCtx.GormClient.GormDb.Delete(&model.SysDeptModel{}, ids)

	return &sysPb.SuccessResp{}, nil
}
