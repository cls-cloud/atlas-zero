package _post

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostInfoLogic {
	return &GetPostInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostInfoLogic) GetPostInfo() (resp *types.PostInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
