package logs

import (
	"context"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VisitTrendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVisitTrendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VisitTrendLogic {
	return &VisitTrendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VisitTrendLogic) VisitTrend() (resp *types.VisitTrendResp, err error) {

	return &types.VisitTrendResp{
		Dates:  []string{"2024-10-15", "2024-10-16", "2024-10-17", "2024-10-18", "2024-10-19", "2024-10-20", "2024-10-21", "2024-10-22"},
		PvList: []int64{3678, 3761, 3778, 4371, 1395, 1247, 4171, 1175},
		UvList: []int64{678, 761, 778, 471, 139, 127, 411, 117},
		IpList: []int64{485, 470, 478, 434, 185, 175, 498, 176},
	}, nil
}
