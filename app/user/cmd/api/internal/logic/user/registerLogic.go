package user

import (
	"context"
	"github.com/pkg/errors"
	"study-zero/app/user/cmd/rpc/usercenter"
	"study-zero/app/user/model"

	"study-zero/app/user/cmd/api/internal/svc"
	"study-zero/app/user/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// register
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {

	res, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		AuthKey:  req.Mobile,
		AuthType: model.UserAuthTypeSystem,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	// copier toValue must be pointer, formValue can be pointer or value
	_ = copier.Copy(resp, res)
	return resp, nil
}
