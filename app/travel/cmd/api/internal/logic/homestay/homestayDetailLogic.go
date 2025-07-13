package homestay

import (
	"context"

	"study-zero/app/travel/cmd/api/internal/svc"
	"study-zero/app/travel/cmd/api/internal/types"
	"study-zero/app/travel/cmd/rpc/travel"
	"study-zero/pkg/tool"
	"study-zero/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// homestay room detail
func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDetailLogic {
	return &HomestayDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayDetailLogic) HomestayDetail(req *types.HomestayDetailReq) (resp *types.HomestayDetailResp, err error) {
	homestayDetail, err := l.svcCtx.TravelRpc.HomestayDetail(l.ctx, &travel.HomestayDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get homestay detail fail"), " get homestay detail db err , id : %d , err : %v ", req.Id, err)
	}

	var homestay types.Homestay
	if homestayDetail != nil {
		_ = copier.Copy(&homestay, homestayDetail)

		homestay.FoodPrice = tool.Fen2Yuan(homestayDetail.Homestay.FoodPrice)
		homestay.HomestayPrice = tool.Fen2Yuan(homestayDetail.Homestay.HomestayPrice)
		homestay.MarketHomestayPrice = tool.Fen2Yuan(homestayDetail.Homestay.MarketHomestayPrice)
	}

	return &types.HomestayDetailResp{
		Homestay: homestay,
	}, nil
}
