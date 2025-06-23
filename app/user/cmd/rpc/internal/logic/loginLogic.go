package logic

import (
	"context"

	"study-zero/app/user/cmd/rpc/internal/svc"
	"study-zero/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

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
	// todo: add your logic here and delete this line

	return &pb.GenerateTokenResp{}, nil
}
