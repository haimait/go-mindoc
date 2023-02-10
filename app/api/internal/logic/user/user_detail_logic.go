package user

import (
	"context"
	"github.com/jinzhu/copier"
	client "go-mindoc/app/rpc/client/user"

	"go-mindoc/app/api/internal/svc"
	"go-mindoc/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailReq) (resp *types.UserDetailResp, err error) {
	loginResp, err := l.svcCtx.UserRpc.UserDetail(l.ctx, &client.UserDetailReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	var res types.UserDetailResp
	_ = copier.Copy(&res, loginResp)
	return &res, nil
}
