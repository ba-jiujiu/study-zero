package logic

import (
	"context"

	"study-zero/app/order/cmd/rpc/internal/svc"
	"study-zero/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayOrderDetailLogic {
	return &HomestayOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// homestay order detail
func (l *HomestayOrderDetailLogic) HomestayOrderDetail(in *pb.HomestayOrderDetailRequest) (*pb.HomestayOrderDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.HomestayOrderDetailResponse{}, nil
}
