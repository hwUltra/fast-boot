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

type DictTypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeListLogic {
	return &DictTypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictTypeListLogic) DictTypeList(in *sysPb.DictTypeListReq) (*sysPb.DictTypeListResp, error) {
	dictTypeModel := model.SysDictTypeModel{}
	total := int64(0)
	if err := l.svcCtx.GormConn.Model(dictTypeModel).
		Scopes(dictTypeModel.WithKeywords(in.Keywords)).
		Count(&total).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get  err : %v , in :%+v", err, in)
	}
	list := make([]*sysPb.SysDictType, 0)
	if total > 0 {
		items := make([]*model.SysDictTypeModel, 0)
		l.svcCtx.GormConn.Model(dictTypeModel).
			Scopes(
				gormV2.Paginate(int(in.PageNum), int(in.PageSize)),
				dictTypeModel.WithKeywords(in.Keywords)).
			Order("id asc").Find(&items)

		for _, item := range items {
			var it sysPb.SysDictType
			_ = copier.Copy(&it, item)

			list = append(list, &it)
		}
	}

	return &sysPb.DictTypeListResp{List: list, Total: total}, nil
}
