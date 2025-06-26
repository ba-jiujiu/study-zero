package user

import (
	"context"
	"github.com/jinzhu/copier"
	"study-zero/app/user/cmd/rpc/usercenter"
	"study-zero/pkg/ctxdata"

	"study-zero/app/user/cmd/api/internal/svc"
	"study-zero/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get user info
func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {

	userId := ctxdata.GetUidFromCtx(l.ctx)

	user, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: userId,
	})

	if err != nil {
		return nil, err
	}

	var _user types.User
	_ = copier.Copy(&_user, user)

	resp.UserInfo = _user

	return
}
