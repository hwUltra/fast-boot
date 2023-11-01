package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"strings"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsGoodsDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsGoodsDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsGoodsDelLogic {
	return &PmsGoodsDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsGoodsDelLogic) PmsGoodsDel(in *pmsPb.IdsReq) (*pmsPb.SuccessResp, error) {
	ids := strings.Split(in.Ids, ",")
	l.svcCtx.GormConn.Delete(&model.PmsGoodsModel{}, ids)
	return &pmsPb.SuccessResp{}, nil
}
