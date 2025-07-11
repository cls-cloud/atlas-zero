package client

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

func (l *UpdateLogic) Update(req *types.ModifyClientReq) error {
	toMapOmit := utils.StructToMapOmit(req.ClientBase, nil, nil, true)
	if _, err := l.svcCtx.Query.SysClient.WithContext(l.ctx).Where(l.svcCtx.Query.SysClient.ID.Eq(req.ID)).Updates(toMapOmit); err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
