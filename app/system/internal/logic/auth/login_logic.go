package auth

import (
	"context"
	"fmt"
	"github.com/mssola/useragent"
	"monitor/pb/monitor"
	"net/http"
	"strings"
	"system/internal/svc"
	"system/internal/types"
	"time"
	"toolkit/auth"
	"toolkit/constants"
	"toolkit/errx"
	"toolkit/helper"
	"toolkit/ip"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginStatus := false
	resp, err = l.LoginHandler(req, err)
	q := l.svcCtx.Query
	if err == nil {
		loginStatus = true
		l.LoginInfo("登陆成功", loginStatus, req)
		//	更新登陆时间
		userMap := map[string]interface{}{
			"login_date": time.Now(),
			"login_ip":   ip.GetClientIP(l.r),
		}
		if _, err = q.SysUser.WithContext(l.ctx).Where(q.SysUser.UserName.Eq(req.Username), q.SysUser.TenantID.Eq(req.TenantId)).Updates(userMap); err != nil {
			logx.Error(err)
		}
	} else {
		l.LoginInfo(err.Error(), loginStatus, req)
	}
	return
}

func (l *LoginLogic) LoginHandler(req *types.LoginReq, err error) (*types.LoginResp, error) {
	//判断clientId是否正常
	q := l.svcCtx.Query
	if req.ClientId == "" {
		return nil, errx.AuthErr("clientId不能为空")
	}
	client, err := q.SysClient.WithContext(l.ctx).Where(q.SysClient.ClientID.Eq(req.ClientId)).
		Where(q.SysClient.GrantType.Like(fmt.Sprintf("%%%s%%", req.GrantType))).
		First()
	if err != nil {
		return nil, errx.GORMErrMsg(err, fmt.Sprintf("客户端id: %s 认证类型：%s 认证异常!", req.ClientId, req.GrantType))
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
	var accessExpire = int64(client.Timeout)
	if int64(client.Timeout) == 0 {
		accessExpire = l.svcCtx.Config.JwtAuth.AccessExpire
	}
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	userid := user.UserID
	userInfo := auth.UserInfo{
		UserId:   userid,
		TenantId: user.TenantID,
		ClientId: client.ClientID,
		Timeout:  client.Timeout,
	}
	token, err := auth.GenerateToken(userInfo, accessSecret, accessExpire)
	if err != nil {
		return nil, err
	}

	err = auth.NewAuth(l.svcCtx.Rds, &userInfo).SetToken(l.ctx, fmt.Sprintf(constants.TOKEN_CACHE, client.ClientID, userid), token, int64(client.ActiveTimeout), accessExpire)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{AccessToken: token}, nil
}

func (l *LoginLogic) LoginInfo(msg string, status bool, req *types.LoginReq) {
	q := l.svcCtx.Query
	client, _ := q.SysClient.WithContext(l.ctx).Where(q.SysClient.ClientID.Eq(req.ClientId)).
		Where(q.SysClient.GrantType.Like(fmt.Sprintf("%%%s%%", req.GrantType))).
		First()
	location := "Unknown"
	cuRip := ip.GetClientIP(l.r)
	ipAddr := "Unknown"
	logx.Info("请求 IP: ", cuRip)
	if !ip.IsPrivateIP(cuRip) {
		ipData, _ := ip.LookupIP(cuRip)
		location = fmt.Sprintf("%s|%s|%s", ipData.Country, ipData.Region, ipData.City)
		ipAddr = cuRip
	} else {
		ipAddr = "内网IP"
	}
	userAgentStr := l.r.Header.Get("User-Agent")
	ua := useragent.New(userAgentStr)
	browserName, _ := ua.Browser()
	deviceType := "PC"
	if ua.Mobile() {
		deviceType = "Mobile"
	}
	statusStr := "1"
	if status {
		statusStr = "0"
	}
	if _, err := l.svcCtx.LoginInfoRpc.Save(l.ctx, &monitor.LoginInfoReq{
		Username:      req.Username,
		TenantId:      req.TenantId,
		ClientKey:     client.ClientKey,
		DeviceType:    deviceType,
		Ipaddr:        ipAddr,
		LoginLocation: location,
		Browser:       browserName,
		Os:            ParseOS(ua.OS()),
		Status:        statusStr,
		Msg:           msg,
	}); err != nil {
		logx.Error(err)
	}
}

func ParseOS(ua string) string {
	ua = strings.ToLower(ua)
	switch {
	case strings.Contains(ua, "windows nt 10.0"):
		return "Windows 10"
	case strings.Contains(ua, "windows nt 11.0"):
		return "Windows 11"
	case strings.Contains(ua, "mac os x"):
		return "OSX"
	case strings.Contains(ua, "android"):
		return "Android"
	case strings.Contains(ua, "harmonyos"):
		return "Harmony"
	case strings.Contains(ua, "linux"):
		return "Linux"
	case strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad"):
		return "iPhone"
	default:
		return "Unknown"
	}
}
