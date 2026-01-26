// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package tenant

import (
	"context"
	"fmt"
	"ovra/toolkit/auth"
	"ovra/toolkit/tenant"

	"ovra/app/system/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearLogic {
	return &ClearLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClearLogic) Clear() error {
	userId := auth.GetUserId(l.ctx)
	key := fmt.Sprintf(tenant.TENANT_KEY, userId)
	ex, err := l.svcCtx.Rds.ExistsCtx(l.ctx, key)
	if err != nil {
		return err
	}
	if !ex {
		return nil
	}
	_, err = l.svcCtx.Rds.DelCtx(l.ctx, key)
	if err != nil {
		return err
	}
	return nil
}
