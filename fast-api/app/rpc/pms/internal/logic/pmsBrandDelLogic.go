package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"strings"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsBrandDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsBrandDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsBrandDelLogic {
	return &PmsBrandDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsBrandDelLogic) PmsBrandDel(in *pmsPb.IdsReq) (*pmsPb.SuccessResp, error) {
	ids := strings.Split(in.Ids, ",")
	l.svcCtx.GormClient.GormDb.Delete(&model.PmsBrandModel{}, ids)
	return &pmsPb.SuccessResp{}, nil
}
