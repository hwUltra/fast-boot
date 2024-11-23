package menu

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"fmt"
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

// 获取菜单
func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.SysMenuGetReq) (resp *types.SysMenuInfo, err error) {
	res, err := l.svcCtx.SysRpc.MenuGet(l.ctx, &sysPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var menuInfo types.SysMenuInfo
	if err := copier.Copy(&menuInfo, res); err != nil {
		return nil, err
	}
	fmt.Println("res", res.Params)
	return &menuInfo, nil
}
