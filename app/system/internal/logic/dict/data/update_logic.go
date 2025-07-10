package data

import (
	"context"
	"toolkit/errx"
	"toolkit/utils"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.ModifyDictDataReq) error {
	toMapOmit := utils.StructToMapOmit(req.DictDataBase, nil, []string{"CreateTime"}, true)
	if _, err := l.svcCtx.Query.SysDictDatum.WithContext(l.ctx).Where(l.svcCtx.Query.SysDictDatum.DictCode.Eq(req.DictCode)).Updates(toMapOmit); err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
