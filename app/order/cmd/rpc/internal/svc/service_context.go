package svc

import (
	"study-zero/app/order/cmd/rpc/internal/config"
	"study-zero/app/order/model"
	"study-zero/app/travel/cmd/rpc/travel"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	TravelRpc travel.Travel

	HomestayOrderModel model.HomestayOrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:             c,
		TravelRpc:          travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
		HomestayOrderModel: model.NewHomestayOrderModel(conn, c.Cache),
	}
}
