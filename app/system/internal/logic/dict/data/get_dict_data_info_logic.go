package data

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictDataInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictDataInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictDataInfoLogic {
	return &GetDictDataInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictDataInfoLogic) GetDictDataInfo() (resp *types.DictDataInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
