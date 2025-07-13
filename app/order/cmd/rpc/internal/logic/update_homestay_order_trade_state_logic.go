package logic

import (
	"context"

	"study-zero/app/order/cmd/rpc/internal/svc"
	"study-zero/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHomestayOrderTradeStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomestayOrderTradeStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomestayOrderTradeStateLogic {
	return &UpdateHomestayOrderTradeStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// update homestay order trade state
func (l *UpdateHomestayOrderTradeStateLogic) UpdateHomestayOrderTradeState(in *pb.UpdateHomestayOrderTradeStateRequest) (*pb.UpdateHomestayOrderTradeStateResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateHomestayOrderTradeStateResponse{}, nil
}
