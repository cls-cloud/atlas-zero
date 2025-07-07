package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"report/internal/config"
	"report/internal/dao/query"
	"report/internal/svc/dbs"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Rds    *redis.Redis
	Query  *query.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := dbs.NewDb(c)
	return &ServiceContext{
		Config: c,
		Rds:    redis.MustNewRedis(c.Data.Redis),
		Db:     db,
		Query:  query.Use(db),
	}
}
