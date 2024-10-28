package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/xerr"
	"github.com/hwUltra/fb-tools/gormV2"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPageLogic {
	return &UserPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPageLogic) UserPage(in *umsPb.PageReq) (*umsPb.UserPageResp, error) {
	userModel := model.UserModel{}

	total := int64(0)
	if err := l.svcCtx.GormConn.Model(userModel).
		Scopes(userModel.WithKeywords(in.Keywords)).
		Count(&total).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get  err : %v , in :%+v", err, in)
	}
	list := make([]*umsPb.User, 0)
	if total > 0 {
		items := make([]*model.UserModel, 0)
		l.svcCtx.GormConn.Model(userModel).
			Preload("Addresses").
			Scopes(
				gormV2.Paginate(int(in.PageNum), int(in.PageSize)),
				userModel.WithKeywords(in.Keywords)).
			Order("id asc").Find(&items)

		for _, item := range items {
			var it umsPb.User
			_ = copier.Copy(&it, item)

			list = append(list, &it)
		}
	}

	return &umsPb.UserPageResp{List: list, Total: total}, nil
}
