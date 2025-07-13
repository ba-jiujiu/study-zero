package logic

import (
	"context"

	"study-zero/app/order/cmd/rpc/internal/svc"
	"study-zero/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderListLogic {
	return &UserHomestayOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// user homestay order list
func (l *UserHomestayOrderListLogic) UserHomestayOrderList(in *pb.UserHomestayOrderListRequest) (*pb.UserHomestayOrderListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UserHomestayOrderListResponse{}, nil
}
