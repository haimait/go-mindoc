package svc

import (
	"github.com/haimait/go-mindoc/models"
	"github.com/haimait/go-mindoc/user/api/internal/config"
	"github.com/haimait/go-mindoc/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
	DB      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB(c.DB.DataSource)
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		DB:      models.DB,
	}
}
