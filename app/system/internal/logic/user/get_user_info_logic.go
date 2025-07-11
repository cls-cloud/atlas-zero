package user

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"toolkit/errx"
	"toolkit/helper"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserInfoResp, err error) {
	// 获取用户权限
	// 获取用户角色
	q := l.svcCtx.Query
	userId := helper.GetUserId(l.ctx)
	resp = new(types.UserInfoResp)
	sysUser := l.svcCtx.Query.SysUser
	user, err := sysUser.WithContext(l.ctx).Where(sysUser.UserID.Eq(userId)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.GORMErrMsg(err, "用户不存在")
		}
		return nil, err
	}
	sysRole := l.svcCtx.Query.SysRole
	sysUserRole := l.svcCtx.Query.SysUserRole
	var roleIds []string
	err = sysUserRole.WithContext(l.ctx).Select(sysUserRole.RoleID).Where(sysUserRole.UserID.Eq(userId)).Scan(&roleIds)
	if err != nil {
		return nil, err
	}
	isAdmin := false
	for _, item := range roleIds {
		if item == "1" {
			isAdmin = true
		}
	}
	sysRole.WithContext(l.ctx).Where(sysRole.RoleID.In())
	roles, err := sysRole.WithContext(l.ctx).
		Where(sysRole.RoleID.In(roleIds...)).
		Find()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.GORMErr(err)
		}
		return nil, err
	}
	roleKeys := make([]string, len(roles))
	for i, item := range roles {
		roleKeys[i] = item.RoleKey
	}
	if isAdmin {
		//TODO Permissions 菜单权限
		resp.Permissions = append(resp.Permissions, "*:*:*")
		resp.Roles = append(resp.Roles, "superadmin")
	} else {
		resp.Roles = roleKeys
		permissions := make([]string, 0)
		err = q.SysMenu.WithContext(l.ctx).Select(q.SysMenu.Perms).
			LeftJoin(q.SysRoleMenu, q.SysRoleMenu.MenuID.EqCol(q.SysMenu.MenuID)).
			LeftJoin(q.SysUserRole, q.SysUserRole.RoleID.EqCol(q.SysRoleMenu.RoleID)).
			Where(q.SysUserRole.UserID.Eq(user.UserID)).
			Where(q.SysMenu.Perms.IsNotNull(), q.SysMenu.Perms.Neq("")).
			Distinct(q.SysMenu.Perms).
			Scan(&permissions)
		if err != nil {
			return nil, err
		}
		resp.Permissions = permissions
	}

	err = copier.Copy(&resp.User, user)
	err = copier.Copy(&resp.User.Roles, roles)
	if err != nil {
		return nil, err
	}
	return
}
