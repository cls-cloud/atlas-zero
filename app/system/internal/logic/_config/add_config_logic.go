package _config

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigLogic {
	return &AddConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddConfigLogic) AddConfig(req *types.ModifyConfigReq) error {
	// todo: add your logic here and delete this line

	return nil
}
