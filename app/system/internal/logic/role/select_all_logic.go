package role

import (
	"context"
	"strconv"
	"strings"
	"system/internal/dao/model"
	"toolkit/errx"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectAllLogic {
	return &SelectAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectAllLogic) SelectAll(req *types.SelectAllReq) error {
	userIds := strings.Split(req.UserIds, ",")
	q := l.svcCtx.Query
	userRoles := make([]*model.SysUserRole, len(userIds))
	for i, id := range userIds {
		userId, _ := strconv.ParseInt(id, 10, 64)
		userRole := &model.SysUserRole{
			UserID: userId,
			RoleID: req.RoleId,
		}
		userRoles[i] = userRole
	}
	if err := q.SysUserRole.WithContext(l.ctx).CreateInBatches(userRoles, len(userRoles)); err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
