package user

import (
	"context"
	"toolkit/errx"
	"toolkit/helper"

	"system/internal/svc"
	"system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPwdLogic {
	return &ResetPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPwdLogic) ResetPwd(req *types.ResetPwdReq) error {
	q := l.svcCtx.Query
	if len(req.Password) < 5 || len(req.Password) > 20 {
		return errx.BizErr("密码长度为5 - 20")
	}
	_, err := q.SysUser.WithContext(l.ctx).Where(q.SysUser.UserID.Eq(req.UserId)).
		Update(q.SysUser.Password, helper.Encrypt(req.Password))
	if err != nil {
		return errx.GORMErr(err)
	}
	return nil
}
