package _type

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictTypeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictTypeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictTypeInfoLogic {
	return &GetDictTypeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictTypeInfoLogic) GetDictTypeInfo() (resp *types.DictTypeInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
