package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/hwUltra/fb-tools/gormx"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserThirdPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserThirdPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserThirdPageLogic {
	return &UserThirdPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserThirdPageLogic) UserThirdPage(in *umsPb.PageReq) (*umsPb.UserThirdPageResp, error) {
	gModel := model.UserThirdModel{}

	total := int64(0)
	if err := l.svcCtx.GormClient.GormDb.Model(gModel).
		Scopes(gModel.WithKeywords(in.Keywords)).
		Count(&total).Error; err != nil {
		return nil, err
	}
	list := make([]*umsPb.UserThird, 0)
	if total > 0 {
		items := make([]*model.UserThirdModel, 0)
		l.svcCtx.GormClient.GormDb.Model(gModel).
			Scopes(
				gormx.Paginate(int(in.PageNum), int(in.PageSize)),
				gModel.WithKeywords(in.Keywords)).
			Order("id asc").Find(&items)

		for _, item := range items {
			var it umsPb.UserThird
			_ = copier.Copy(&it, item)

			list = append(list, &it)
		}
	}

	return &umsPb.UserThirdPageResp{List: list, Total: total}, nil
}
