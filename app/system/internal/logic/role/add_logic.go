package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"system/internal/dao/model"
	"system/internal/svc"
	"system/internal/types"
	"toolkit/errx"
	"toolkit/utils"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddOrUpdateRoleReq) error {
	q := l.svcCtx.Query
	if oRoles, err := q.SysRole.WithContext(l.ctx).Where(q.SysRole.RoleKey.Eq(req.RoleKey)).First(); err == nil && oRoles != nil {
		return errx.BizErr("角色编码已存在")
	}
	roleId := utils.GetID()
	if req.RoleID != "" {
		roleId = req.RoleID
	}
	role := &model.SysRole{
		RoleID:    roleId,
		RoleKey:   req.RoleKey,
		RoleName:  req.RoleName,
		RoleSort:  req.RoleSort,
		Status:    req.Status,
		Remark:    req.Remark,
		DataScope: req.DataScope,
	}
	if req.TenantID != "" {
		role.TenantID = req.TenantID
	}
	if err := q.SysRole.WithContext(l.ctx).Create(role); err != nil {
		return errx.GORMErr(err)
	}
	if len(req.MenuIds) != 0 {
		roleMenus := make([]*model.SysRoleMenu, 0, len(req.MenuIds))
		for _, menuId := range req.MenuIds {
			roleMenus = append(roleMenus, &model.SysRoleMenu{
				RoleID: roleId,
				MenuID: menuId,
			})
		}
		if err := q.SysRoleMenu.WithContext(l.ctx).CreateInBatches(roleMenus, len(roleMenus)); err != nil {
			return errx.GORMErr(err)
		}
	}
	return nil
}
