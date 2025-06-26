package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"study-zero/app/user/cmd/rpc/usercenter"
	"study-zero/pkg/xerr"

	"study-zero/app/user/cmd/rpc/internal/svc"
	"study-zero/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByAuthKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByAuthKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByAuthKeyLogic {
	return &GetUserAuthByAuthKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByAuthKeyLogic) GetUserAuthByAuthKey(in *pb.GetUserAuthByAuthKeyReq) (*pb.GetUserAuthResp, error) {
	authType, err := l.svcCtx.UserAuthModel.FindOneByAuthTypeAuthKey(l.ctx, in.AuthType, in.AuthKey)
	if err != nil && errors.Is(err, sqlx.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user auth failed"), "err: %v, authKey: %s, authType: %s", err, in.AuthKey, in.AuthType)
	}

	var _authType usercenter.UserAuth
	_ = copier.Copy(&_authType, authType)

	return &pb.GetUserAuthResp{
		UserAuth: &_authType,
	}, nil
}
