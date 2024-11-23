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

type DictDataPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDataPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataPageLogic {
	return &DictDataPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDataPageLogic) DictDataPage(in *sysPb.DictDataPageReq) (*sysPb.DictDataPageResp, error) {
	dictDataModel := model.SysDictDataModel{}

	total := int64(0)
	if err := l.svcCtx.GormClient.GormDb.Model(dictDataModel).
		Where("dict_id = ?", in.DictId).
		Scopes(dictDataModel.WithKeywords(in.Keywords)).
		Count(&total).Error; err != nil {
		return nil, err
	}
	list := make([]*sysPb.SysDictData, 0)
	if total > 0 {
		items := make([]*model.SysDictDataModel, 0)
		l.svcCtx.GormClient.GormDb.Model(dictDataModel).
			Where("dict_id = ?", in.DictId).
			Scopes(
				dictDataModel.WithKeywords(in.Keywords),
				gormx.Paginate(int(in.PageNum), int(in.PageSize))).
			Order("id asc").Find(&items)

		for _, item := range items {
			var it sysPb.SysDictData
			_ = copier.Copy(&it, item)

			list = append(list, &it)
		}
	}

	return &sysPb.DictDataPageResp{List: list, Total: total}, nil
}
