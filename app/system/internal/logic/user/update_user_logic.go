package user

import (
	"context"
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
	_, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.UserID.Eq(req.UserID)).Updates(userToMap)
	if err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
