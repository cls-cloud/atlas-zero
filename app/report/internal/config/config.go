package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Data DataConfig
}

type DataConfig struct {
	Database DatabaseConfig
	Redis    redis.RedisConf
	Cache    CacheConfig
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
