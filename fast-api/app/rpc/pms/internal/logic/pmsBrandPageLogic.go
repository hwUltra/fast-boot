package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/xerr"
	"github.com/hwUltra/fb-tools/gormV2"
	"github.com/hwUltra/fb-tools/gormV2/types"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsBrandPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsBrandPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsBrandPageLogic {
	return &PmsBrandPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsBrandPageLogic) PmsBrandPage(in *pmsPb.PmsBrandPageReq) (*pmsPb.PmsBrandPageResp, error) {
	m := model.PmsBrandModel{}

	total := int64(0)
	if err := l.svcCtx.GormConn.Model(m).
		Scopes(
			m.WithShopId(in.ShopId),
			m.WithKeywords(in.Keywords)).
		Count(&total).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get  err : %v , in :%+v", err, in)
	}

	list := make([]*pmsPb.PmsBrand, 0)

	if total > 0 {
		items := make([]*model.PmsBrandModel, 0)
		l.svcCtx.GormConn.Model(m).
			Scopes(
				gormV2.Paginate(int(in.PageNum), int(in.PageSize)),
				m.WithShopId(in.ShopId),
				m.WithKeywords(in.Keywords)).
			Order("id asc").Find(&items)

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

	}

	return &pmsPb.PmsBrandPageResp{List: list, Total: total}, nil
}
