package homestayBussiness

import (
	"context"

	"study-zero/app/travel/cmd/api/internal/svc"
	"study-zero/app/travel/cmd/api/internal/types"
	"study-zero/app/travel/model"
	"study-zero/app/user/cmd/rpc/pb"
	"study-zero/app/user/cmd/rpc/usercenter"
	"study-zero/pkg/xerr"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GoodBossLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// good boss
func NewGoodBossLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodBossLogic {
	return &GoodBossLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodBossLogic) GoodBoss(req *types.GoodBossReq) (resp *types.GoodBossResp, err error) {
	builder := l.svcCtx.HomestayActivityModel.SelectBuilder().Where(squirrel.Eq{
		"row_type":   model.HomestayActivityGoodBusiType,
		"row_status": model.HomestayActivityUpStatus,
	})
	list, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, builder, 0, 10, "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get GoodBoss db err. rowType: %s ,err : %v", model.HomestayActivityGoodBusiType, err)
	}

	var goodBossList []types.HomestayBusinessBoss
	if len(list) > 0 {

		mr.MapReduceVoid(
			func(source chan<- any) {
				for _, homestayActivity := range list {
					source <- homestayActivity.DataId
				}
			}, func(item any, write mr.Writer[*usercenter.User], cancel func(error)) {
				id := item.(int64)

				userResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
					Id: id,
				})

				if err != nil {
					logx.WithContext(l.ctx).Errorf("GoodListLogic GoodList fail userId : %d ,err:%v", id, err)
					return
				}

				if userResp != nil && userResp.User != nil && userResp.User.Id > 0 {
					write.Write(userResp.User)
				}
			}, func(pipe <-chan *usercenter.User, cancel func(error)) {
				for item := range pipe {
					var typeHomestayBusiness types.HomestayBusinessBoss
					_ = copier.Copy(&typeHomestayBusiness, item)

					goodBossList = append(goodBossList, typeHomestayBusiness)
				}
			})
	}

	return &types.GoodBossResp{
		List: goodBossList,
	}, nil
}
