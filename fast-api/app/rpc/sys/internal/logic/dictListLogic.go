package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/gormV2"
	"fast-boot/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictListLogic {
	return &DictListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictListLogic) DictList(in *sysPb.DictListReq) (*sysPb.DictListResp, error) {
	dictModel := model.SysDictModel{}
	total := int64(0)
	if err := l.svcCtx.GormConn.Model(dictModel).
		Where("type_code = ?", in.TypeCode).
		Count(&total).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get  err : %v , in :%+v", err, in)
	}
	list := make([]*sysPb.SysDict, 0)
	if total > 0 {
		items := make([]*model.SysDictModel, 0)
		l.svcCtx.GormConn.Model(dictModel).
			Where("type_code = ?", in.TypeCode).
			Scopes(gormV2.Paginate(int(in.PageNum), int(in.PageSize))).
			Order("sort asc,id asc").Find(&items)

		for _, item := range items {
			var it sysPb.SysDict
			_ = copier.Copy(&it, item)

			list = append(list, &it)
		}
	}

	return &sysPb.DictListResp{List: list, Total: total}, nil
}
