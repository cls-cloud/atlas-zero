package menu

import (
	"context"
	"toolkit/errx"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TreeSelectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTreeSelectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TreeSelectLogic {
	return &TreeSelectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TreeSelectLogic) TreeSelect() (resp []*types.RoleMenuTree, err error) {
	resp = make([]*types.RoleMenuTree, 0)
	q := l.svcCtx.Query
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
	tree := l.Tree(menuTree, "0")
	resp = tree

	return
}

func (l *TreeSelectLogic) Tree(node []*types.RoleMenuTree, pid string) []*types.RoleMenuTree {
	res := make([]*types.RoleMenuTree, 0)
	for _, v := range node {
		if v.ParentId == pid {
			v.Children = l.Tree(node, v.Id)
			res = append(res, v)
		}
	}
	return res
}
