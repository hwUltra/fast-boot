package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"strings"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDelLogic {
	return &DictDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDelLogic) DictDel(in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {
	ids := strings.Split(in.Ids, ",")
	l.svcCtx.GormClient.GormDb.Delete(&model.SysDictModel{}, ids)

	return &sysPb.SuccessResp{}, nil
}
