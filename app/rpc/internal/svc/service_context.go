package svc

import (
	"github.com/haimait/go-mindoc/app/rpc/internal/config"
	"github.com/haimait/go-mindoc/models"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	//RedisClient *redis.Redis
	DB *gorm.DB

	//model
	//UserModel models.SysUser
	//UserModel models.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		//RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
		//	r.Type = c.Redis.Type
		//	r.Pass = c.Redis.Pass
		//}),
		DB: models.DB,
		//UserModel: &models.SysUser{},
		//UserModel: models.NewSysUserModel(),
	}
}
