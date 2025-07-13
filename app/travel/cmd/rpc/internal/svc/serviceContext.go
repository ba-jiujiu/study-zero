package svc

import (
	"study-zero/app/travel/cmd/rpc/internal/config"
	model "study-zero/app/travel/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	HomestayModel model.HomestayModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,

		HomestayModel: model.NewHomestayModel(conn, c.Cache),
	}
}
