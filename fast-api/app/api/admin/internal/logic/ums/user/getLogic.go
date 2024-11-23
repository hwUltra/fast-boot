package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.UserGetReq) (resp *types.UserInfo, err error) {
	res, err := l.svcCtx.UmsRpc.UserGet(l.ctx, &umsPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var info types.UserInfo
	if err := copier.Copy(&info, res); err != nil {
		return nil, err
	}

	return &info, nil
}
