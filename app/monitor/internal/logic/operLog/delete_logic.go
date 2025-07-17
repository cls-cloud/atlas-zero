package operLog

import (
	"context"
	"strings"
	"toolkit/errx"

	"monitor/internal/svc"
	"monitor/internal/types"

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

func (l *DeleteLogic) Delete(req *types.IdsReq) error {
	ids := strings.Split(req.Ids, ",")
	q := l.svcCtx.Query
	if _, err := q.SysOperLog.WithContext(l.ctx).Where(q.SysOperLog.OperID.In(ids...)).Unscoped().Delete(); err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
