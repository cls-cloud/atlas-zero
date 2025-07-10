package menu

import (
	"context"
	"system/internal/svc"
	"system/internal/types"
	"toolkit/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleMenuTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleMenuTreeLogic {
	return &RoleMenuTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleMenuTreeLogic) RoleMenuTree(req *types.IdReq) (resp *types.RoleMenuTreeResp, err error) {
	resp = new(types.RoleMenuTreeResp)
	q := l.svcCtx.Query
	roleMenus, err := q.SysRoleMenu.WithContext(l.ctx).
		LeftJoin(q.SysMenu, q.SysMenu.MenuID.EqCol(q.SysRoleMenu.MenuID)).
		Where(q.SysMenu.MenuType.Eq(TypeButton)).
		Where(q.SysRoleMenu.RoleID.Eq(req.Id)).Find()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	menuIds := make([]string, 0, len(roleMenus))
	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, roleMenu.MenuID)
	}
	resp.CheckedKeys = menuIds
	treeSelect, err := NewTreeSelectLogic(l.ctx, l.svcCtx).TreeSelect()
	if err != nil {
		return nil, err
	}
	resp.Menus = treeSelect
	return
}
