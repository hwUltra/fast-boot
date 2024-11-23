package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/hwUltra/fb-tools/gormx"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysNoticePageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysNoticePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysNoticePageLogic {
	return &SysNoticePageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysNoticePageLogic) SysNoticePage(in *sysPb.SysNoticePageReq) (*sysPb.SysNoticePageResp, error) {

	snModel := model.SysNoticeModel{}
	total := int64(0)
	if err := l.svcCtx.GormClient.GormDb.Model(snModel).
		Scopes(
			snModel.WithCreatedAt(in.StartTime, in.EndTime),
			snModel.WithTitle(in.Title),
			snModel.WithUid(in.Uid),
			snModel.WithStatus(in.PublishStatus)).
		Count(&total).Error; err != nil {
		return nil, err
	}
	list := make([]*sysPb.SysNotice, 0)
	if total > 0 {
		items := make([]*model.SysNoticeModel, 0)
		l.svcCtx.GormClient.GormDb.Model(snModel).
			Scopes(
				snModel.WithCreatedAt(in.StartTime, in.EndTime),
				snModel.WithTitle(in.Title),
				snModel.WithUid(in.Uid),
				snModel.WithStatus(in.PublishStatus),
				gormx.Paginate(int(in.PageNum), int(in.PageSize))).
			Order("id desc").Find(&items)

		for _, item := range items {
			var it sysPb.SysNotice
			_ = copier.Copy(&it, item)
			list = append(list, &it)
		}
	}

	return &sysPb.SysNoticePageResp{List: list, Total: total}, nil
}
