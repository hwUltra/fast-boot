package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/zeromicro/go-zero/core/logx"
)

type DictSimListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictSimListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictSimListLogic {
	return &DictSimListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictSimListLogic) DictSimList(in *sysPb.AnyReq) (*sysPb.DictSimListResp, error) {
	dictModel := model.SysDictModel{}
	items := make([]*model.SysDictModel, 0)
	l.svcCtx.GormConn.Model(dictModel).
		Preload("DataList").
		Order("sort asc,id asc").Find(&items)
	list := make([]*sysPb.DictSimItem, 0)
	for _, item := range items {
		DataSimItems := make([]*sysPb.DictDataSimItem, 0)
		for _, data := range item.DataList {
			DataSimItems = append(DataSimItems, &sysPb.DictDataSimItem{
				Value:   data.Value,
				Label:   data.Label,
				TagType: data.Tag,
			})
		}
		list = append(list, &sysPb.DictSimItem{
			Name:         item.Name,
			DictCode:     item.Code,
			DictDataList: DataSimItems,
		})
	}

	return &sysPb.DictSimListResp{List: list}, nil
}
