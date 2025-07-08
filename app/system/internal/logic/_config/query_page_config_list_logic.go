package _config

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageConfigListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPageConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageConfigListLogic {
	return &QueryPageConfigListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPageConfigListLogic) QueryPageConfigList(req *types.QueryPageConfigListReq) (resp *types.QueryPageConfigListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
