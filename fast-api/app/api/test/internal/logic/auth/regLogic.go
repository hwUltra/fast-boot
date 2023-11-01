package auth

import (
	"context"
	"fast-boot/app/api/test/internal/svc"
	"fast-boot/app/api/test/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegLogic {
	return &RegLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegLogic) Reg(req *types.RegRequest) (resp *types.TokenResponse, err error) {

	//memberAddResp, err := l.svcCtx.Ums.MemberAdd(l.ctx, &umsclient.MemberAddReq{
	//	Username: req.UserName,
	//	Password: req.Password,
	//	Phone:    req.Mobile,
	//})
	//
	//if err != nil {
	//	reqStr, _ := json.Marshal(req)
	//	logx.WithContext(l.ctx).Errorf("会员注册失败,参数: %s,响应：%s", reqStr, err.Error())
	//	return nil, err
	//}
	//
	//return &types.TokenResponse{
	//	Token: memberAddResp.Token,
	//}, nil

	return resp, err

}
