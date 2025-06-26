package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2/cache"
	"study-zero/app/user/cmd/rpc/usercenter"
	"study-zero/app/user/model"
	"study-zero/pkg/xerr"

	"github.com/silenceper/wechat/v2"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/zeromicro/go-zero/core/logx"
	"study-zero/app/user/cmd/api/internal/svc"
	"study-zero/app/user/cmd/api/internal/types"
)

var ErrWxMiniAuthFailed = xerr.NewErrMsg("wx mini program auth failed")

type WxMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewWxMiniAuthLogic wechat mini auth
func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxMiniAuthLogic) WxMiniAuth(req *types.WXMiniAuthReq) (resp *types.TokenResp, err error) {

	// get wechat mini program instance
	wxApp := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     l.svcCtx.Config.WxMiniConf.AppId,
		AppSecret: l.svcCtx.Config.WxMiniConf.Secret,
		Cache:     cache.NewMemory(),
	})

	authRes, err := wxApp.GetAuth().Code2Session(req.Code)
	if err != nil || authRes.ErrCode != 0 || authRes.OpenID == "" {
		return nil, errors.Wrapf(ErrWxMiniAuthFailed, "wx authrization failed, err: %v, code: %s, authRes: %#v", err, req.Code, authRes)
	}

	// parse authRes
	userData, err := wxApp.GetEncryptor().Decrypt(authRes.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailed, "parse wx user data failed, err: %v, encryptedData: %s, iv: %s, sessionKey: %s", err, req.EncryptedData, req.IV, authRes.SessionKey)
	}

	var userId int64
	userAuthRes, err := l.svcCtx.UsercenterRpc.GetUserAuthByAuthKey(l.ctx, &usercenter.GetUserAuthByAuthKeyReq{
		AuthType: model.UserAuthTypeSmallWX,
		AuthKey:  authRes.OpenID,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailed, "call user center rpc failed, err: %v", err)
	}

	// user never bind with wx mini program
	if userAuthRes == nil || userAuthRes.UserAuth.Id == 0 {
		registerRes, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.RegisterReq{
			AuthType: model.UserAuthTypeSmallWX,
			Mobile:   userData.PhoneNumber,
			Nickname: userData.NickName,
			AuthKey:  authRes.OpenID,
		})
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailed, "UsercenterRpc.Register err :%v, authResult : %+v", err, authRes)
		}

		return &types.TokenResp{
			AccessToken:  registerRes.AccessToken,
			AccessExpire: registerRes.AccessExpire,
			RefreshAfter: registerRes.AccessExpire,
		}, nil
	}

	// user already bind with wx mini program, login
	userId = userAuthRes.UserAuth.UserId
	token, err := l.svcCtx.UsercenterRpc.GenerateToken(l.ctx, &usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailed, "call user center rpc generate token failed, err: %v", err)
	}

	return &types.TokenResp{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
		RefreshAfter: token.AccessExpire,
	}, nil
}
