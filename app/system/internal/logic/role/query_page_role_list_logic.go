package role

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPageRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageRoleListLogic {
	return &QueryPageRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPageRoleListLogic) QueryPageRoleList(req *types.QueryPageRoleListReq) (resp *types.QueryPageRoleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
