package queues

import (
	"context"
	"encoding/json"
	"study-zero/app/order/cmd/mq/internal/svc"
	"study-zero/app/order/cmd/rpc/orderservice"
	orderModel "study-zero/app/order/model"
	paymentModel "study-zero/app/payment/model"
	"study-zero/pkg/kqueue"
	"study-zero/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentUpdateStatusMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentUpdateStatusMq(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentUpdateStatusMq {
	return &PaymentUpdateStatusMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (mq *PaymentUpdateStatusMq) Consume(ctx context.Context, key, val string) error {

	var message kqueue.ThirdPaymentUpdatePayStatusNotifyMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(mq.ctx).Error("PaymentUpdateStatusMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := mq.execService(message); err != nil {
		logx.WithContext(mq.ctx).Error("PaymentUpdateStatusMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (mq *PaymentUpdateStatusMq) execService(message kqueue.ThirdPaymentUpdatePayStatusNotifyMessage) error {

	orderTradeState := mq.getOrderTradeStateByPaymentTradeState(message.PayStatus)
	if orderTradeState != -99 {
		// update order status
		_, err := mq.svcCtx.OrderRpc.UpdateHomestayOrderTradeState(mq.ctx, &orderservice.UpdateHomestayOrderTradeStateRequest{
			Sn:         message.OrderSn,
			TradeState: orderTradeState,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("update homestay order state fail"), "update homestay order state fail err : %v ,message:%+v", err, message)
		}
	}

	return nil
}

func (mq *PaymentUpdateStatusMq) getOrderTradeStateByPaymentTradeState(thirdPaymentPayStatus int64) int64 {
	switch thirdPaymentPayStatus {
	case paymentModel.ThirdPaymentPayTradeStateSuccess:
		return orderModel.HomestayOrderTradeStateWaitUse
	case paymentModel.ThirdPaymentPayTradeStateRefund:
		return orderModel.HomestayOrderTradeStateRefund
	default:
		return -99
	}
}
