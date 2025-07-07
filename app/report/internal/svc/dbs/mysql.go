package dbs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	slog "log"
	"os"
	"report/internal/config"
	"time"
)

func NewDb(c config.Config) *gorm.DB {
	diff := c.Data.Database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		diff.Username, diff.Password, diff.Host, diff.Port, diff.Database,
	)
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢查询sql阀值
			Colorful:      true,        //禁用彩色打印
			LogLevel:      logger.Info,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{
			//SingularTable: true, //表名是否加s
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to MySQL database.")
	return db
}
