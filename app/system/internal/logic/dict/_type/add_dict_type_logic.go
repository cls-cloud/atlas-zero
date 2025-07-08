package _type

import (
	"context"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDictTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictTypeLogic {
	return &AddDictTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDictTypeLogic) AddDictType(req *types.ModifyDictTypeReq) error {
	// todo: add your logic here and delete this line

	return nil
}
