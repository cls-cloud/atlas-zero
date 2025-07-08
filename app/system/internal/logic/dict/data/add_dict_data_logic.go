package data

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDictDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDictDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictDataLogic {
	return &AddDictDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDictDataLogic) AddDictData(req *types.ModifyDictDataReq) error {
	// todo: add your logic here and delete this line

	return nil
}
