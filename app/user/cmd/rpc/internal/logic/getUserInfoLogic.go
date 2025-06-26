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

var ErrUserNoExist = xerr.NewErrMsg("user not found")

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && errors.Is(err, sqlx.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "query Id %d user db err: %v", in.Id, err)
	}

	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExist, "Id %d user not exist", in.Id)
	}

	var _user *usercenter.User
	_ = copier.Copy(_user, user)

	return &pb.GetUserInfoResp{
		User: _user,
	}, nil
}
