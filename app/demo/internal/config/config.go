// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"ovra/toolkit/configshared"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	RestConf    rest.RestConf
	RpcConf     zrpc.RpcServerConf
	Tenant      configshared.TenantConfig
	Data        configshared.DataConfig
	JwtAuth     configshared.JwtAuthConfig
	ApiDecrypt  configshared.ApiDecryptConfig
	Captcha     configshared.CaptchaConfig
	Idempotency configshared.IdempotencyConfig
}
