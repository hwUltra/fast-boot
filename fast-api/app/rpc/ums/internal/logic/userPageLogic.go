package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"
	"fast-boot/common/globalkey"
	"fast-boot/common/xerr"
	"github.com/hwUltra/fb-tools/gormx"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
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
	if err := l.svcCtx.GormClient.GormDb.Model(userModel).
		Scopes(userModel.WithKeywords(in.Keywords)).
		Count(&total).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get  err : %v , in :%+v", err, in)
	}
	list := make([]*umsPb.User, 0)
	if total > 0 {
		items := make([]*model.UserModel, 0)
		l.svcCtx.GormClient.GormDb.Model(userModel).
			Scopes(
				gormx.Paginate(int(in.PageNum), int(in.PageSize)),
				userModel.WithKeywords(in.Keywords)).
			Order("id asc").Find(&items)

		if err := copier.CopyWithOption(&list, items, globalkey.CopierProtoOptions); err != nil {
			return nil, err
		}
	}

	return &umsPb.UserPageResp{List: list, Total: total}, nil
}
