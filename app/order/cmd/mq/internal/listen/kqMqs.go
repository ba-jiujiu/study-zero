package listen

import (
	"context"
	"study-zero/app/order/cmd/mq/internal/config"
	"study-zero/app/order/cmd/mq/internal/queues"
	"study-zero/app/order/cmd/mq/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		kq.MustNewQueue(
			c.PaymentUpdateStatusConf,
			queues.NewPaymentUpdateStatusMq(ctx, svcContext),
		),
	}
}
