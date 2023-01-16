package svc

import (
	"github.com/haimait/go-mindoc/app/api/internal/config"
	"github.com/haimait/go-mindoc/app/api/internal/middleware"
	client "github.com/haimait/go-mindoc/app/rpc/client/user"
	"github.com/haimait/go-mindoc/models"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	Auth    rest.Middleware
	DB      *gorm.DB
	RDB     *redis.Client
	UserRpc client.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB(c.DB.DataSource)
	models.NewRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	return &ServiceContext{
		Config:  c,
		UserRpc: client.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		DB:      models.DB,
		RDB:     models.RDB,
		Auth:    middleware.NewAuthMiddleware(c.JwtAuth.JwtKey).Handle,
	}
}
