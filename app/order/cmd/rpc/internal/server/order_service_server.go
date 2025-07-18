// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: order.proto

package server

import (
	"context"

	"study-zero/app/order/cmd/rpc/internal/logic"
	"study-zero/app/order/cmd/rpc/internal/svc"
	"study-zero/app/order/cmd/rpc/pb"
)

type OrderServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderServiceServer
}

func NewOrderServiceServer(svcCtx *svc.ServiceContext) *OrderServiceServer {
	return &OrderServiceServer{
		svcCtx: svcCtx,
	}
}

// create homestay order
func (s *OrderServiceServer) CreateHomestayOrder(ctx context.Context, in *pb.CreateHomestayOrderRequest) (*pb.CreateHomestayOrderResponse, error) {
	l := logic.NewCreateHomestayOrderLogic(ctx, s.svcCtx)
	return l.CreateHomestayOrder(in)
}

// homestay order detail
func (s *OrderServiceServer) HomestayOrderDetail(ctx context.Context, in *pb.HomestayOrderDetailRequest) (*pb.HomestayOrderDetailResponse, error) {
	l := logic.NewHomestayOrderDetailLogic(ctx, s.svcCtx)
	return l.HomestayOrderDetail(in)
}

// update homestay order trade state
func (s *OrderServiceServer) UpdateHomestayOrderTradeState(ctx context.Context, in *pb.UpdateHomestayOrderTradeStateRequest) (*pb.UpdateHomestayOrderTradeStateResponse, error) {
	l := logic.NewUpdateHomestayOrderTradeStateLogic(ctx, s.svcCtx)
	return l.UpdateHomestayOrderTradeState(in)
}

// user homestay order list
func (s *OrderServiceServer) UserHomestayOrderList(ctx context.Context, in *pb.UserHomestayOrderListRequest) (*pb.UserHomestayOrderListResponse, error) {
	l := logic.NewUserHomestayOrderListLogic(ctx, s.svcCtx)
	return l.UserHomestayOrderList(in)
}
