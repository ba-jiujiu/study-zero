package homestayBussiness

import (
	"context"

	"study-zero/app/travel/cmd/api/internal/svc"
	"study-zero/app/travel/cmd/api/internal/types"
	"study-zero/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBussinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// business list
func NewHomestayBussinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBussinessListLogic {
	return &HomestayBussinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBussinessListLogic) HomestayBussinessList(req *types.HomestayBussinessListReq) (resp *types.HomestayBussinessListResp, err error) {
	builder := l.svcCtx.HomestayBusinessModel.SelectBuilder()
	list, err := l.svcCtx.HomestayBusinessModel.FindPageListByIdDESC(l.ctx, builder, req.LastId, req.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HomestayBussinessList FindPageListByIdDESC db fail ,  req : %+v , err:%v", req, err)
	}

	var homestayBusinessList []types.HomestayBusinessListInfo
	if len(list) > 0 {
		for _, homestayBusiness := range list {
			var typeHomestayBusinessInfo types.HomestayBusinessListInfo
			_ = copier.Copy(&typeHomestayBusinessInfo, homestayBusiness)

			homestayBusinessList = append(homestayBusinessList, typeHomestayBusinessInfo)
		}
	}

	return &types.HomestayBussinessListResp{
		List: homestayBusinessList,
	}, nil
}
