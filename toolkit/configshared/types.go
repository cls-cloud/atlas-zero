package configshared

import "github.com/zeromicro/go-zero/core/stores/redis"

type JwtAuthConfig struct {
	AccessSecret         string
	AccessExpire         int64
	MultipleLoginDevices bool
}

type ApiDecryptConfig struct {
	Enabled    bool
	HeaderFlag string
	PublicKey  string
	PrivateKey string
}

type CaptchaConfig struct {
	Enabled bool
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

type DatabaseConfig struct {
	Logger string
	DatabaseCfg
}

type DatabaseCfg struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Ssl      bool
}

type CacheConfig struct {
	Expire int
}

type RsaConfig struct {
	PubB64 string
	PriB64 string
}

type IdempotencyConfig struct {
	Enabled       bool
	Header        string
	ExpireSeconds int
	IncludePaths  []string
	ExcludePaths  []string
}
