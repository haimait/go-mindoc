package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		JwtKey             string
		TakenExpire        int64
		RefreshTokenExpire int64
	}
	DB struct {
		DataSource string
	}

	Redis struct {
		Addr     string
		Password string
		DB       int
	}

	UserRpcConf zrpc.RpcClientConf
}
