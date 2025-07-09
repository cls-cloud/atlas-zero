package menu

import (
	"context"
	"toolkit/errx"

	"system/internal/svc"
	"system/internal/types"

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
	roleMenus, err := q.SysRoleMenu.WithContext(l.ctx).Where(q.SysRoleMenu.RoleID.Eq(req.Id)).Find()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	menuIds := make([]int64, 0, len(roleMenus))
	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, roleMenu.MenuID)
	}
	resp.CheckedKeys = menuIds
	sysMenus, err := q.SysMenu.WithContext(l.ctx).Find()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	var menuTree []*types.RoleMenuTree
	for _, menu := range sysMenus {
		menuTree = append(menuTree, &types.RoleMenuTree{
			Id:       menu.MenuID,
			ParentId: menu.ParentID,
			Icon:     menu.Icon,
			MenuType: menu.MenuType,
			Label:    menu.MenuName,
		})
	}
	tree := l.Tree(menuTree, 0)
	resp.Menus = tree
	return
}

func (l *RoleMenuTreeLogic) Tree(node []*types.RoleMenuTree, pid int64) []*types.RoleMenuTree {
	res := make([]*types.RoleMenuTree, 0)
	for _, v := range node {
		if v.ParentId == pid {
			v.Children = l.Tree(node, v.Id)
			res = append(res, v)
		}
	}
	return res
}
