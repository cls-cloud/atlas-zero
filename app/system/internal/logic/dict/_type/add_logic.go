package _type

import (
	"context"
	"system/internal/dao/model"
	"toolkit/errx"
	"toolkit/utils"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.ModifyDictTypeReq) error {
	dictType := &model.SysDictType{
		DictID:   utils.GetID(),
		DictName: req.DictName,
		DictType: req.DictType,
		Remark:   req.Remark,
	}
	if err := l.svcCtx.Query.SysDictType.WithContext(l.ctx).Create(dictType); err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
