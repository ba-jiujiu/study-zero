package svc

import (
	"study-zero/app/travel/cmd/api/internal/config"
	"study-zero/app/travel/cmd/rpc/travel"
	"study-zero/app/travel/model"
	"study-zero/app/user/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	// rpc
	UsercenterRpc usercenter.Usercenter
	TravelRpc     travel.Travel

	// model
	HomestayModel         model.HomestayModel
	HomestayActivityModel model.HomestayActivityModel
	HomestayBusinessModel model.HomestayBusinessModel
	HomestayCommentModel  model.HomestayCommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,

		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		TravelRpc:     travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),

		HomestayModel:         model.NewHomestayModel(conn, c.Cache),
		HomestayActivityModel: model.NewHomestayActivityModel(conn, c.Cache),
		HomestayBusinessModel: model.NewHomestayBusinessModel(conn, c.Cache),
		HomestayCommentModel:  model.NewHomestayCommentModel(conn, c.Cache),
	}
}
