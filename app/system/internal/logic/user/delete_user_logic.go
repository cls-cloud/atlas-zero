package user

import (
	"context"
	"strings"
	"toolkit/errx"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.CodeReq) error {
	userIds := strings.Split(req.Code, ",")
	for _, userId := range userIds {
		q := l.svcCtx.Query
		_, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.UserID.Eq(userId)).Unscoped().Delete()
		if err != nil {
			return errx.GORMErr(err)
		}
		//删除角色 岗位
		_, err = q.SysUserRole.WithContext(l.ctx).Where(q.SysUserRole.UserID.Eq(userId)).Unscoped().Delete()
		if err != nil {
			return errx.GORMErr(err)
		}
		_, err = q.SysUserPost.WithContext(l.ctx).Where(q.SysUserPost.UserID.Eq(userId)).Unscoped().Delete()
		if err != nil {
			return errx.GORMErr(err)
		}
	}

	return nil
}
