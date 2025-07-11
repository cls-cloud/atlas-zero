package tenant

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

func (l *UpdateLogic) Update(req *types.ModifyTenantReq) error {
	toMapOmit := utils.StructToMapOmit(req.TenantBase, nil, nil, true)
	if _, err := l.svcCtx.Query.SysTenant.WithContext(l.ctx).Where(l.svcCtx.Query.SysTenant.ID.Eq(req.ID)).Updates(toMapOmit); err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
