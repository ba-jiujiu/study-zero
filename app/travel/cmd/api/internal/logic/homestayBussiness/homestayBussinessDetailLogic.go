package homestayBussiness

import (
	"context"

	"study-zero/app/travel/cmd/api/internal/svc"
	"study-zero/app/travel/cmd/api/internal/types"
	"study-zero/app/travel/model"
	"study-zero/app/user/cmd/rpc/usercenter"
	"study-zero/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBussinessDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// boss detail
func NewHomestayBussinessDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBussinessDetailLogic {
	return &HomestayBussinessDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBussinessDetailLogic) HomestayBussinessDetail(req *types.HomestayBussinessDetailReq) (resp *types.HomestayBussinessDetailResp, err error) {
	detail, err := l.svcCtx.HomestayBusinessModel.FindOne(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), " HomestayBussinessDetail  FindOne db fail ,id  : %d , err : %v", req.Id, err)
	}

	var typeHomestayBusinessBoss types.HomestayBusinessBoss
	if detail != nil {

		userResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
			Id: detail.UserId,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("get boss info fail"), "get boss info fail ,  userId : %d ,err:%v", detail.UserId, err)
		}
		if userResp != nil {
			_ = copier.Copy(&typeHomestayBusinessBoss, userResp)
		}
	}

	return &types.HomestayBussinessDetailResp{
		Boss: typeHomestayBusinessBoss,
	}, nil
}
