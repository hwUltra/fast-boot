package notice

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/hwUltra/fb-tools/jwtx"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyPageLogic {
	return &MyPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyPageLogic) MyPage(req *types.SysNoticePageReq) (resp *types.SysNoticePageResp, err error) {
	uid := jwtx.GetUid(l.ctx)
	result, err := l.svcCtx.SysRpc.SysNoticePage(l.ctx, &sysPb.SysNoticePageReq{
		Title:         req.Title,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		PublishStatus: req.Status,
		PageNum:       req.PageNum,
		PageSize:      req.PageSize,
		Uid:           uid,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*types.SysNotice, 0)
	_ = copier.Copy(&list, result.List)

	return &types.SysNoticePageResp{
		Total: result.Total,
		List:  list,
	}, nil

}
