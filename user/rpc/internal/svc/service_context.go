package svc

import (
	"github.com/haimait/go-mindoc/models"
	"github.com/haimait/go-mindoc/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	DB          *gorm.DB
	//UserModel   models.SysUser
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		DB: models.DB,
	}
}