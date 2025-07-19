package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"resource/internal/config"
	"resource/internal/dao/query"
	"resource/internal/middleware"
	"resource/internal/svc/dbs"
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
	rds := redis.MustNewRedis(c.Data.Redis)
	return &ServiceContext{
		Config: c,
		Rds:    rds,
		Db:     db,
		Query:  query.Use(db),
		Auth:   middleware.NewAuthMiddleware(c, rds).Handle,
	}
}
