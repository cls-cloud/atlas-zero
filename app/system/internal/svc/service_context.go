package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"system/internal/config"
	"system/internal/dao/query"
	"system/internal/middleware"
	"system/internal/svc/dbs"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Rds    *redis.Redis
	Query  *query.Query
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := dbs.NewDb(c)
	return &ServiceContext{
		Config: c,
		Rds:    redis.MustNewRedis(c.Data.Redis),
		Db:     db,
		Query:  query.Use(db),
		Auth:   middleware.NewAuthMiddleware(c).Handle,
	}
}
