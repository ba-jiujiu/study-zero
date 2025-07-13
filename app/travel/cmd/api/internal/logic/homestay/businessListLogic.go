package homestay

import (
	"context"

	"study-zero/app/travel/cmd/api/internal/svc"
	"study-zero/app/travel/cmd/api/internal/types"
	"study-zero/pkg/tool"
	"study-zero/pkg/xerr"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// boss all homestay room
func NewBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessListLogic {
	return &BusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BusinessListLogic) BusinessList(req *types.BusinessListReq) (resp *types.BusinessListResp, err error) {
	whereBuilder := l.svcCtx.HomestayModel.SelectBuilder().Where(squirrel.Eq{
		"homestay_business_id": req.HomestayBusinessId,
	})
	list, err := l.svcCtx.HomestayModel.FindPageListByIdDESC(l.ctx, whereBuilder, req.LastId, req.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HomestayBusinessId: %d ,err : %v", req.HomestayBusinessId, err)
	}

	var _list []types.Homestay
	if len(list) > 0 {
		for _, homestay := range list {

			var typeHomestay types.Homestay
			_ = copier.Copy(&typeHomestay, homestay)

			typeHomestay.FoodPrice = tool.Fen2Yuan(homestay.FoodPrice)
			typeHomestay.HomestayPrice = tool.Fen2Yuan(homestay.HomestayPrice)
			typeHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestay.MarketHomestayPrice)

			_list = append(_list, typeHomestay)
		}
	}

	return &types.BusinessListResp{
		List: _list,
	}, nil
}
