package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"study-zero/app/user/cmd/rpc/internal/svc"
	"study-zero/app/user/cmd/rpc/pb"
	"study-zero/pkg/ctxdata"
	"study-zero/pkg/xerr"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrGenerateTokenError = xerr.NewErrMsg("generate token failed")

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accssToken, err := l.getJwtToken(now, accessExpire, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "getJwtToken err userId:%d , err:%v", in.UserId, err)
	}
	return &pb.GenerateTokenResp{
		AccessToken:  accssToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(iat, expired, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + expired
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodES256)
	token.Claims = claims

	return token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
}
