package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"system/internal/svc"
	"system/internal/types"
	"toolkit/errx"
)

type QueryUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserDetailLogic {
	return &QueryUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserDetailLogic) QueryUserDetail(req *types.IdReq) (resp *types.UserDetailResp, err error) {
	q := l.svcCtx.Query
	resp = new(types.UserDetailResp)

	// 查询系统所有角色
	roles, err := q.SysRole.WithContext(l.ctx).Find()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	if err = copier.Copy(&resp.Roles, roles); err != nil {
		return nil, err
	}
	if req == nil {
		return
	}
	// 查询用户基本信息
	id := req.Id
	sysUser, err := q.SysUser.WithContext(l.ctx).
		Where(q.SysUser.UserID.Eq(id)).
		First()
	user := new(types.UserRoles)
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	if err = copier.Copy(&user, sysUser); err != nil {
		return nil, err
	}
	resp.User = user
	// 查询用户的角色ID列表
	userRoles, err := q.SysUserRole.WithContext(l.ctx).
		Select(q.SysUserRole.RoleID).
		Where(q.SysUserRole.UserID.Eq(id)).
		Find()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	roleIds := make([]string, 0)
	for _, r := range userRoles {
		roleIds = append(roleIds, r.RoleID)
	}

	// 查询用户的岗位ID列表
	postIds := make([]string, 0)
	userPosts, err := q.SysUserPost.WithContext(l.ctx).
		Select(q.SysUserPost.PostID).
		Where(q.SysUserPost.UserID.Eq(id)).
		Find()
	if err != nil {
		return nil, errx.GORMErr(err)
	}
	for _, p := range userPosts {
		postIds = append(postIds, p.PostID)
	}
	posts := make([]*types.PostBase, 0)
	resp.Posts = posts
	resp.RoleIds = roleIds
	resp.PostIds = postIds
	return
}
