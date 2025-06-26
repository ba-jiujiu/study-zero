package logic

import (
	"context"
	"github.com/pkg/errors"
	"study-zero/app/user/model"
	"study-zero/pkg/tool"
	"study-zero/pkg/xerr"

	"study-zero/app/user/cmd/rpc/internal/svc"
	"study-zero/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUsernamePwdError = xerr.NewErrMsg("username or password error")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.GenerateTokenResp, error) {

	var userId int64
	var err error

	switch in.AuthType {
	case model.UserAuthTypeSystem:
		userId, err = l.loginByMobile(in.AuthKey, in.Password)
	default:
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	if err != nil {
		return nil, err
	}

	genTokenService := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := genTokenService.GenerateToken(&pb.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return tokenResp, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find user by mobile %s error", mobile)
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserNoExist, "mobile %s not exist", mobile)
	}

	if tool.Md5ByString(password) != user.Password {
		return 0, errors.Wrapf(ErrUsernamePwdError, "username or password error")
	}

	return user.Id, nil
}

func (l *LoginLogic) loginByWx() error {
	return nil
}
