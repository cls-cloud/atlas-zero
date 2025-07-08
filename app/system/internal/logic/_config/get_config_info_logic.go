package _config

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigInfoLogic {
	return &GetConfigInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigInfoLogic) GetConfigInfo() (resp *types.ConfigInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
