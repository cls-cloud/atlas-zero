package auth

import (
	"context"
	"strconv"
	"system/internal/svc"
	"system/internal/types"
	"toolkit/errx"
	"toolkit/helper"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if l.svcCtx.Config.Captcha.Enabled {
		data, err := l.svcCtx.Rds.GetCtx(l.ctx, "captcha:"+req.Uuid)
		if err != nil {
			return nil, errx.AuthErr("验证码失效")
		}
		if data != req.Code {
			//验证码错误则重新获取验证码
			_, _ = l.svcCtx.Rds.DelCtx(l.ctx, "captcha:"+req.Uuid)
			return nil, errx.AuthErr("验证码验证失败")
		}
	}
	sysUser := l.svcCtx.Query.SysUser
	user, err := sysUser.WithContext(l.ctx).Where(sysUser.UserName.Eq(req.Username)).First()
	if err != nil {
		return nil, errx.GORMErrMsg(err, "用户不存在")
	}
	//判断密码是否匹配
	if !helper.Verify(user.Password, req.Password) {
		return nil, errx.AuthErr("密码验证失败")
	}
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	userid := strconv.FormatInt(user.UserID, 10)
	token, err := helper.GenerateToken(userid, user.TenantID, accessSecret, accessExpire)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{AccessToken: token}, nil
}
