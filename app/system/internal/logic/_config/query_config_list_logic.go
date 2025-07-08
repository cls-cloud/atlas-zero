package _config

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryConfigListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryConfigListLogic {
	return &QueryConfigListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryConfigListLogic) QueryConfigList(req *types.QueryConfigListReq) (resp []types.ConfigDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
