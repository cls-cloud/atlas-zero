package _config

import (
	"context"
	"strings"
	"toolkit/errx"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.CodeReq) error {
	ids := strings.Split(req.Code, ",")
	q := l.svcCtx.Query
	if _, err := q.SysConfig.WithContext(l.ctx).Where(q.SysConfig.ConfigID.In(ids...)).Unscoped().Delete(); err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
