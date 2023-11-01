package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"strings"

	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDelLogic {
	return &UserDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDelLogic) UserDel(in *umsPb.IdsReq) (*umsPb.SuccessResp, error) {
	ids := strings.Split(in.Ids, ",")
	l.svcCtx.GormConn.Delete(&model.UserModel{}, ids)
	return &umsPb.SuccessResp{}, nil
}
