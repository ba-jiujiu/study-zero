package homestay

import (
	"context"
	"study-zero/app/travel/cmd/api/internal/svc"
	"study-zero/app/travel/cmd/api/internal/types"
	"study-zero/app/travel/model"
	"study-zero/pkg/tool"
	"study-zero/pkg/xerr"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type HomestayListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// homestay room list
func NewHomestayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayListLogic {
	return &HomestayListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayListLogic) HomestayList(req *types.HomestayListReq) (resp *types.HomestayListResp, err error) {
	whereBuilder := l.svcCtx.HomestayActivityModel.SelectBuilder().Where(squirrel.Eq{
		"row_type": model.HomestayActivityPreferredType,
		"status":   model.HomestayActivityUpStatus,
	})

	list, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, whereBuilder, req.Page, req.PageSize, "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get activity homestay id set fail rowType: %s ,err : %v", model.HomestayActivityPreferredType, err)
	}

	var res []types.Homestay
	if len(list) > 0 {
		mr.MapReduceVoid(func(source chan<- any) {
			for _, homestayActivity := range list {
				source <- homestayActivity.DataId
			}
		}, func(item any, writer mr.Writer[*model.Homestay], cancel func(error)) {
			id := item.(int64)

			homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, id)
			if err != nil && err != model.ErrNotFound {
				logx.WithContext(l.ctx).Errorf("ActivityHomestayListLogic ActivityHomestayList 获取活动数据失败 id : %d ,err : %v", id, err)
				return
			}
			writer.Write(homestay)
		}, func(pipe <-chan *model.Homestay, cancel func(error)) {
			for homestay := range pipe {
				var typeHomestay types.Homestay
				_ = copier.Copy(&typeHomestay, homestay)

				typeHomestay.FoodPrice = tool.Fen2Yuan(homestay.FoodPrice)
				typeHomestay.HomestayPrice = tool.Fen2Yuan(homestay.HomestayPrice)
				typeHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestay.MarketHomestayPrice)

				res = append(res, typeHomestay)
			}
		})
	}

	return &types.HomestayListResp{
		List: res,
	}, nil
}
