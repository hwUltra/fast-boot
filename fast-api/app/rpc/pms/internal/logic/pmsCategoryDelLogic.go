package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"strings"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsCategoryDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsCategoryDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsCategoryDelLogic {
	return &PmsCategoryDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsCategoryDelLogic) PmsCategoryDel(in *pmsPb.IdsReq) (*pmsPb.SuccessResp, error) {
	ids := strings.Split(in.Ids, ",")
	l.svcCtx.GormConn.Delete(&model.PmsCategoryModel{}, ids)
	return &pmsPb.SuccessResp{}, nil
}
