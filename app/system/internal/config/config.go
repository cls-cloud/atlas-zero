package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	RestConf   rest.RestConf
	RpcConf    zrpc.RpcServerConf
	Tenant     TenantConfig
	Data       DataConfig
	MonitorRpc zrpc.RpcClientConf
	JwtAuth    struct {
		AccessSecret         string
		AccessExpire         int64
		MultipleLoginDevices bool
	}
	Idempotency IdempotencyConfig
	ApiDecrypt  struct {
		Enabled    bool
		HeaderFlag string
		PublicKey  string
		PrivateKey string
	}
	Captcha struct {
		Enabled bool
	}
}

type TenantConfig struct {
	Enabled      bool
	IgnoreTables []string
}

type DataConfig struct {
	Database DatabaseConfig
	Redis    redis.RedisConf
	Cache    CacheConfig
}

type IdempotencyConfig struct {
	Enabled       bool
	Header        string
	ExpireSeconds int
	IncludePaths  []string
	ExcludePaths  []string
}

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}
type CacheConfig struct {
	Expire int
}
