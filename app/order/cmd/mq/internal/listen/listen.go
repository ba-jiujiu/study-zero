package listen

import (
	"context"
	"study-zero/app/order/cmd/mq/internal/config"
	"study-zero/app/order/cmd/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

func Mqs(c config.Config) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var service []service.Service

	service = append(service, KqMqs(c, ctx, svcContext)...)

	return service
}
