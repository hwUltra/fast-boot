package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/gormV2"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsCategoryListLogic {
	return &PmsCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsCategoryListLogic) PmsCategoryList(in *pmsPb.ListReq) (*pmsPb.PmsCategoryListResp, error) {
	sModel := model.PmsCategoryModel{}
	items := make([]*model.PmsCategoryModel, 0)
	l.svcCtx.GormConn.Model(sModel).
		Scopes(
			gormV2.Paginate(int(in.PageNum), int(in.PageSize)),
			sModel.WithKeywords(in.Keywords)).
		Order("sort asc,id asc").Find(&items)

	return &pmsPb.PmsCategoryListResp{List: makeTreeList(items, 0)}, nil
}

func makeTreeList(list []*model.PmsCategoryModel, parentID int64) []*pmsPb.PmsCategory {
	res := make([]*pmsPb.PmsCategory, 0)
	for _, item := range list {
		if item.ParentId == parentID {
			var it pmsPb.PmsCategory
			_ = copier.Copy(&it, item)
			it.CreatedAt = item.CreatedAt.String()
			it.Children = makeTreeList(list, item.Id)
			res = append(res, &it)
		}
	}
	return res
}
