package user

import (
	"context"
	"github.com/jinzhu/copier"
	client "go-mindoc/app/rpc/client/user"
	"go-mindoc/models"

	"go-mindoc/app/api/internal/svc"
	"go-mindoc/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	resp = new(types.UserLoginResp)
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &client.LoginReq{
		AuthType: models.UserAuthTypeSystem,
		AuthKey:  req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var res types.UserLoginResp
	_ = copier.Copy(&res, loginResp)
	return &res, nil
}
