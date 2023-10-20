package svc

import (
	"beyond/application/user/rpc/internal/config"
	"beyond/application/user/rpc/internal/model"
	"database/sql"
	"log"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	DB        *gorm.DB
	Config    config.Config
	Redis     *redis.Redis
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb := redis.MustNewRedis(c.RedisConf)
	conn := sqlx.NewMysql(c.Mysql.DSN)

	return &ServiceContext{
		Config:    c,
		Redis:     rdb,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}

// Init Gorm的Init
func Init(c config.Config) (db *gorm.DB) {
	var (
		sqlDB *sql.DB
		err   error
	)
	mysqlConf := mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/beyond_user?charset=utf8&parseTime=True&loc=Local",
	}
	gormConfig := configLog(c.Mysql.LogMode)
	if db, err = gorm.Open(mysql.New(mysqlConf), gormConfig); err != nil {
		log.Fatal("opens database failed: ", err)
	}
	if sqlDB, err = db.DB(); err != nil {
		log.Fatal("db.db() failed: ", err)
	}

	sqlDB.SetMaxIdleConns(c.Mysql.MaxIdleCons)
	sqlDB.SetMaxOpenConns(c.Mysql.MaxOpenCons)
	return
}

// configLog 根据配置决定是否开启日志
func configLog(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	}
	return
}
