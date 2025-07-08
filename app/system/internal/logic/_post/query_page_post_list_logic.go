package _post

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPagePostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPagePostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPagePostListLogic {
	return &QueryPagePostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPagePostListLogic) QueryPagePostList(req *types.QueryPagePostListReq) (resp *types.QueryPagePostListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
