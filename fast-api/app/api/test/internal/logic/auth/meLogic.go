package auth

import (
	"context"
	"fast-boot/app/api/test/internal/svc"
	"fast-boot/app/api/test/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type MeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeLogic {
	return &MeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeLogic) Me() (resp *types.UserInfoResponse, err error) {
	//uid := jwtx.GetUid(l.ctx)
	//userInfo, err := l.svcCtx.Ums.FindMemberByUid(l.ctx, &umsclient.UidReq{
	//	Uid: uid,
	//})
	//if err != nil {
	return nil, err
	//}
	//
	//return &types.UserInfoResponse{
	//	Id:        uid,
	//	NickName:  userInfo.Nickname,
	//	AvatarURL: userInfo.Icon,
	//}, nil

}
