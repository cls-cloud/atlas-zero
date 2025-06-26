package menu

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPageMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageMenuListLogic {
	return &QueryPageMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPageMenuListLogic) QueryPageMenuList(req *types.QueryPageMenuListReq) (resp *types.QueryPageMenuListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
