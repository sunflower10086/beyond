package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConf redis.RedisConf
	Mysql     struct {
		DSN         string
		LogMode     bool
		MaxOpenCons int
		MaxIdleCons int
	}
	CacheRedis cache.CacheConf
}
