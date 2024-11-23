package logs

import (
	"context"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VisitStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVisitStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VisitStatsLogic {
	return &VisitStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VisitStatsLogic) VisitStats() (resp []types.VisitStatsItem, err error) {
	return []types.VisitStatsItem{
		{
			Type:             "pv",
			Title:            "浏览量",
			TodayCount:       1175,
			TotalCount:       458801,
			GrowthRate:       0.68,
			GranularityLabel: "日",
		},
		{
			Type:             "uv",
			Title:            "访客数",
			TodayCount:       100,
			TotalCount:       2000,
			GrowthRate:       0,
			GranularityLabel: "日",
		},
		{
			Type:             "ip",
			Title:            "IP数",
			TodayCount:       176,
			TotalCount:       28233,
			GrowthRate:       0.22,
			GranularityLabel: "日",
		},
	}, err
}
