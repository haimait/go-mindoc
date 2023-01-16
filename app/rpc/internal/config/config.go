package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	JwtAuth struct {
		JwtKey             string
		TakenExpire        int64
		RefreshTokenExpire int64
	}
	Cache cache.CacheConf
	DB    struct {
		DataSource string
	}
	//Redis struct {
	//	Addr     string
	//	Password string
	//	DB       int
	//}
}
