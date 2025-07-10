package role

import (
	"context"
	"system/internal/dao/model"
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

func (l *UpdateLogic) Update(req *types.AddOrUpdateRoleReq) error {
	if req.RoleID == 0 {
		return errx.BizErr("角色ID不能为空")
	}
	q := l.svcCtx.Query
	if oRoles, err := q.SysRole.WithContext(l.ctx).Where(q.SysRole.RoleKey.Eq(req.RoleKey)).
		Where(q.SysRole.RoleID.Neq(req.RoleID)).
		First(); err == nil && oRoles != nil {
		return errx.BizErr("角色编码已存在")
	}
	omit := utils.StructToMapOmit(req.RoleBase, nil, []string{"SuperAdmin"}, true)
	if _, err := q.SysRole.WithContext(l.ctx).Where(q.SysRole.RoleID.Eq(req.RoleID)).Updates(omit); err != nil {
		return errx.GORMErr(err)
	}
	if len(req.MenuIds) != 0 {
		if _, err := q.SysRoleMenu.WithContext(l.ctx).Where(q.SysRoleMenu.RoleID.Eq(req.RoleID)).Unscoped().Delete(); err != nil {
			return errx.GORMErr(err)
		}
		roleMenus := make([]*model.SysRoleMenu, 0, len(req.MenuIds))
		for _, menuId := range req.MenuIds {
			roleMenus = append(roleMenus, &model.SysRoleMenu{
				RoleID: req.RoleID,
				MenuID: menuId,
			})
		}
		if err := q.SysRoleMenu.WithContext(l.ctx).CreateInBatches(roleMenus, len(roleMenus)); err != nil {
			return errx.GORMErr(err)
		}
	}
	return nil
}
