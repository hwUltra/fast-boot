package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"strings"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictTypeDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictTypeDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeDelLogic {
	return &DictTypeDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictTypeDelLogic) DictTypeDel(in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {
	ids := strings.Split(in.Ids, ",")
	l.svcCtx.GormConn.Delete(&model.SysDictTypeModel{}, ids)
	return &sysPb.SuccessResp{}, nil
}
