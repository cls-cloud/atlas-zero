package user

import (
	"context"
	"system/internal/dao/model"
	"toolkit/errx"
	"toolkit/utils"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.AddOrUpdateUserReq) error {
	if req.UserID == 0 {
		return errx.BizErr("用户ID不能为空")
	}
	userToMap := utils.StructToMapOmit(req.UserBase, nil, []string{"Password"}, true)
	//更新数据
	q := l.svcCtx.Query
	if _, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.UserID.Eq(req.UserID)).Updates(userToMap); err != nil {
		return errx.GORMErr(err)
	}
	if _, err := q.SysUserPost.WithContext(l.ctx).Where(q.SysUserPost.UserID.Eq(req.UserID)).Delete(); err != nil {
		return errx.GORMErr(err)
	}
	if _, err := q.SysUserRole.WithContext(l.ctx).Where(q.SysUserRole.UserID.Eq(req.UserID)).Delete(); err != nil {
		return errx.GORMErr(err)
	}
	if len(req.RoleIds) != 0 {
		userRole := make([]*model.SysUserRole, len(req.RoleIds))
		for i, roleId := range req.RoleIds {
			userRole[i] = &model.SysUserRole{
				UserID: req.UserID,
				RoleID: roleId,
			}
		}
		if err := q.SysUserRole.WithContext(l.ctx).CreateInBatches(userRole, len(userRole)); err != nil {
			return errx.GORMErr(err)
		}
	}
	if len(req.PostIds) != 0 {
		userPost := make([]*model.SysUserPost, len(req.PostIds))
		for i, postId := range req.PostIds {
			userPost[i] = &model.SysUserPost{
				UserID: req.UserID,
				PostID: postId,
			}
		}
		if err := q.SysUserPost.WithContext(l.ctx).CreateInBatches(userPost, len(userPost)); err != nil {
			return errx.GORMErr(err)
		}
	}

	return nil
}
