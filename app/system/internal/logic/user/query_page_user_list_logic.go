package user

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPageUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageUserListLogic {
	return &QueryPageUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPageUserListLogic) QueryPageUserList(req *types.QueryPageUserListReq) (resp *types.QueryPageUserListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
