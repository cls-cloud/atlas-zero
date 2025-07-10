package data

import (
	"context"
	"github.com/jinzhu/copier"
	"toolkit/errx"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DataByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DataByTypeLogic {
	return &DataByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DataByTypeLogic) DataByType(req *types.CodeReq) (resp []*types.DictDataBase, err error) {
	resp = make([]*types.DictDataBase, 0)
	q := l.svcCtx.Query
	sysDictData, err := q.SysDictDatum.WithContext(l.ctx).Where(q.SysDictDatum.DictType.Eq(req.Code)).Find()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	err = copier.Copy(&resp, sysDictData)
	if err != nil {
		return nil, err
	}
	return
}
