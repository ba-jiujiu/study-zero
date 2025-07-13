package HomestayOrder

import (
	"context"

	"study-zero/app/order/cmd/api/internal/svc"
	"study-zero/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHomestayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// create homestay order
func NewCreateHomestayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHomestayOrderLogic {
	return &CreateHomestayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateHomestayOrderLogic) CreateHomestayOrder(req *types.CreateHomestayOrderRequest) (resp *types.CreateHomestayOrderResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
