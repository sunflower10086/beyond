package svc

import (
	"beyond/application/applet/internal/config"
	"beyond/application/user/rpc/user"

	"beyond/pkg/interceptors"

	"beyond/application/applet/internal/middleware"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Auth    rest.Middleware
	Redis   *redis.Redis
	UserRPC user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:  c,
		Auth:    middleware.NewAuthMiddleware(c.Auth.AccessSecret).Handle,
		Redis:   redis.MustNewRedis(c.BizRedis),
		UserRPC: user.NewUser(userRPC),
	}
}
