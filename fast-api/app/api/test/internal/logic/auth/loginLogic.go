package auth

import (
	"context"
	"fast-boot/app/api/test/internal/svc"
	"fast-boot/app/api/test/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.TokenResponse, err error) {

	//menu, err := l.svcCtx.SysRpc.AddSysMenu(l.ctx, &sysclient.AddSysMenuReq{
	//	ParentId:  1,
	//	Order:     99,
	//	Title:     "test",
	//	Icon:      "fa",
	//	Uri:       "/test",
	//	Extension: "xxxbbb",
	//	Show:      1,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("menu", menu.String())

	//result, err := l.svcCtx.SysRpc.UpdateSysMenu(l.ctx, &sysclient.UpdateSysMenuReq{
	//	Id:        1,
	//	ParentId:  0,
	//	Order:     1,
	//	Title:     "test0001",
	//	Icon:      "fa",
	//	Uri:       "/test0001",
	//	Extension: "xxxbbb",
	//	Show:      1,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("result", result.String())

	//uid, err := l.svcCtx.SysRpc.AddSysUsers(l.ctx, &sysclient.AddSysUsersReq{
	//	Username: "kyle",
	//	Name:     "kyle",
	//	Password: "123123",
	//})
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("uid", uid.String())

	//userInfo, err := l.svcCtx.SysRpc.GetSysUsersByUserName(l.ctx,
	//	&sysclient.GetSysUsersByUserNameReq{Username: "kyle"})
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("userInfo=", userInfo.String())

	return resp, err

}
