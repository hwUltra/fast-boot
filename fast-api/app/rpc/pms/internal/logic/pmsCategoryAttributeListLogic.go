package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/hwUltra/fb-tools/gormx/types"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type PmsCategoryAttributeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsCategoryAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsCategoryAttributeListLogic {
	return &PmsCategoryAttributeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsCategoryAttributeListLogic) PmsCategoryAttributeList(in *pmsPb.PmsCategoryAttributeListReq) (*pmsPb.PmsCategoryAttributeListResp, error) {
	m := model.PmsCategoryAttributeModel{}
	items := make([]*model.PmsCategoryAttributeModel, 0)
	l.svcCtx.GormClient.GormDb.Model(m).
		Where("type = ?", in.Type).
		Where("category_id = ?", in.CategoryId).
		Order("id asc").Find(&items)

	list := make([]*pmsPb.PmsCategoryAttribute, 0)

	if err := copier.CopyWithOption(&list, items, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters: []copier.TypeConverter{
			{
				SrcType: types.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(types.Time)
					if !ok {
						return nil, status.Error(100, "src type not matching")
					}
					return s.String(), nil
				},
			},
		},
	}); err != nil {
		return nil, err
	}

	return &pmsPb.PmsCategoryAttributeListResp{List: list}, nil
}
