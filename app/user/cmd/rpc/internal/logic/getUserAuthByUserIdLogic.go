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

type GetUserAuthByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByUserIdLogic {
	return &GetUserAuthByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByUserIdLogic) GetUserAuthByUserId(in *pb.GetUserAuthByUserIdReq) (*pb.GetUserAuthResp, error) {
	authType, err := l.svcCtx.UserAuthModel.FindOneByUserIdAuthType(l.ctx, in.UserId, in.AuthType)
	if err != nil && errors.Is(err, sqlx.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user auth failed"), "err: %v, userId: %d, authType: %s", err, in.UserId, in.AuthType)
	}

	var _authType usercenter.UserAuth
	_ = copier.Copy(&_authType, authType)

	return &pb.GetUserAuthResp{
		UserAuth: &_authType,
	}, nil
}
