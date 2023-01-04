package user

import (
	"context"
	"github.com/haimait/go-mindoc/models"
	"github.com/haimait/go-mindoc/user/api/internal/svc"
	"github.com/haimait/go-mindoc/user/api/internal/types"
	"github.com/haimait/go-mindoc/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx context.Context

	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	resp = new(types.UserLoginResp)
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
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
