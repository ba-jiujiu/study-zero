package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"study-zero/app/user/cmd/api/internal/config"
	"study-zero/app/user/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config config.Config
	// UsercenterRpc
	UsercenterRpc usercenter.Usercenter

	SetUidToCtxMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
