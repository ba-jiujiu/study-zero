package logic

import (
	"context"

	"study-zero/app/order/cmd/rpc/internal/svc"
	"study-zero/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHomestayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateHomestayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHomestayOrderLogic {
	return &CreateHomestayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// create homestay order
func (l *CreateHomestayOrderLogic) CreateHomestayOrder(in *pb.CreateHomestayOrderRequest) (*pb.CreateHomestayOrderResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateHomestayOrderResponse{}, nil
}
