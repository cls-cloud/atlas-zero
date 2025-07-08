package data

import (
	"context"
	"github.com/jinzhu/copier"
	"toolkit/errx"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDataByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDataByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDataByTypeLogic {
	return &GetDataByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDataByTypeLogic) GetDataByType(req *types.CodeReq) (resp []types.DictDataDetailResp, err error) {
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
