package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/xerr"
	"github.com/hwUltra/fb-tools/gormV2"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictPageLogic {
	return &DictPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictPageLogic) DictPage(in *sysPb.DictPageReq) (*sysPb.DictPageResp, error) {
	dictModel := model.SysDictModel{}

	total := int64(0)
	if err := l.svcCtx.GormConn.Model(dictModel).
		Scopes(dictModel.WithKeywords(in.Keywords)).
		Count(&total).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get  err : %v , in :%+v", err, in)
	}
	list := make([]*sysPb.SysDict, 0)
	if total > 0 {
		items := make([]*model.SysDictModel, 0)
		l.svcCtx.GormConn.Model(dictModel).
			Scopes(
				gormV2.Paginate(int(in.PageNum), int(in.PageSize))).
			Order("id asc").Find(&items)

		for _, item := range items {
			var it sysPb.SysDict
			_ = copier.Copy(&it, item)

			list = append(list, &it)
		}
	}

	return &sysPb.DictPageResp{List: list, Total: total}, nil
}