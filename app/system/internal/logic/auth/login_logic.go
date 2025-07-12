package auth

import (
	"context"
	"fmt"
	"system/internal/svc"
	"system/internal/types"
	"toolkit/constants"
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
	//判断clientId是否正常
	q := l.svcCtx.Query
	if req.ClientId == "" {
		return nil, errx.AuthErr("clientId不能为空")
	}
	client, err := q.SysClient.WithContext(l.ctx).Where(q.SysClient.ClientID.Eq(req.ClientId)).First()
	if err != nil {
		return nil, errx.GORMErrMsg(err, "clientId不存在")
	}
	_, err = q.SysTenant.WithContext(l.ctx).Where(q.SysTenant.TenantID.Eq(req.TenantId)).First()
	if err != nil {
		return nil, errx.GORMErrMsg(err, "租户不存在")
	}
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
	user, err := sysUser.WithContext(l.ctx).Where(sysUser.UserName.Eq(req.Username)).
		Where(sysUser.TenantID.Eq(req.TenantId)).
		First()

	if err != nil {
		return nil, errx.GORMErrMsg(err, "用户不存在")
	}
	//判断密码是否匹配
	if !helper.Verify(user.Password, req.Password) {
		return nil, errx.AuthErr("密码验证失败")
	}
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	userid := user.UserID
	token, err := helper.GenerateToken(userid, user.TenantID, accessSecret, accessExpire)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.Rds.SetexCtx(l.ctx, fmt.Sprintf(constants.TOKEN_CACHE, client.ClientID, userid), token, int(client.Timeout))
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{AccessToken: token}, nil
}
