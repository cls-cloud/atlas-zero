package client

import (
	"context"
	"toolkit/errx"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeStatusLogic {
	return &ChangeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeStatusLogic) ChangeStatus(req *types.ClientQuery) (resp []*types.ClientBase, err error) {
	q := l.svcCtx.Query
	if _, err := q.SysClient.WithContext(l.ctx).Where(q.SysClient.ClientID.Eq(req.ClientId)).Update(q.SysClient.Status, req.Status); err != nil {
		return nil, errx.GORMErr(err)
	}
	return
}
