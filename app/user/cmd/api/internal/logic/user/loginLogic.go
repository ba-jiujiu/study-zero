package user

import (
	"context"
	"github.com/jinzhu/copier"
	"study-zero/app/user/cmd/rpc/usercenter"
	"study-zero/app/user/model"

	"study-zero/app/user/cmd/api/internal/svc"
	"study-zero/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.TokenResp, err error) {

	loginRes, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &usercenter.LoginReq{
		AuthType: model.UserAuthTypeSystem,
		Password: req.Password,
		AuthKey:  req.Mobile,
	})

	if err != nil {
		return nil, err
	}

	_ = copier.Copy(resp, loginRes)

	return
}
