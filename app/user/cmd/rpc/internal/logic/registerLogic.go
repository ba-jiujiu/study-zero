package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"study-zero/app/user/cmd/rpc/usercenter"
	"study-zero/app/user/model"
	"study-zero/pkg/tool"
	"study-zero/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"study-zero/app/user/cmd/rpc/internal/svc"
)

var ErrUserAlreadyRegisterError = xerr.NewErrMsg("user has been registered")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *usercenter.RegisterReq) (*usercenter.GenerateTokenResp, error) {
	// check if user already exists
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "mobile:%s,err:%v", in.Mobile, err)
	}

	if user != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists mobile:%s,err:%v", in.Mobile, err)
	}

	var userId int64
	if err = l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := new(model.User)
		user.Mobile = in.Mobile
		if len(in.Nickname) == 0 {
			user.Nickname = tool.Krand(8, tool.KC_RAND_KIND_ALL)
		}
		// wxMiniProgram register does not require password, phone register is require passwd valid in api
		if len(in.Password) > 0 {
			user.Password = tool.Md5ByString(in.Password)
		}
		// insert user model
		result, err := l.svcCtx.UserModel.Insert(ctx, user)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, user)
		}
		userId, err = result.LastInsertId()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user insertResult.LastInsertId err:%v,user:%+v", err, user)
		}

		userAuth := new(model.UserAuth)
		userAuth.UserId = userId
		userAuth.AuthKey = in.AuthKey
		userAuth.AuthType = in.AuthType
		// insert userAuth model
		if _, err = l.svcCtx.UserAuthModel.Insert(ctx, userAuth); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user_auth Insert err:%v,userAuth:%v", err, userAuth)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	// now successful register user gen token
	tokenService := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := tokenService.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}

	return token, nil
}
