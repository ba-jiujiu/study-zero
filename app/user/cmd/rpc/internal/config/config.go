package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	// JwtAuth configuration for JWT authentication
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	// DB configuration for connecting to the database
	DB struct {
		DataSource string
	}
	// Cache configuration for caching user data
	Cache cache.CacheConf
}
