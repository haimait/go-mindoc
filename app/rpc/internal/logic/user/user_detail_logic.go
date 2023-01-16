package userlogic

import (
	"context"
	client "github.com/haimait/go-mindoc/app/rpc/client/user"
	"github.com/haimait/go-mindoc/models"
	"github.com/jinzhu/copier"

	"github.com/haimait/go-mindoc/app/rpc/internal/svc"
	"github.com/haimait/go-mindoc/app/rpc/pb-desc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDetailLogic) UserDetail(in *pb.UserDetailReq) (*pb.UserDetailResp, error) {
	var err error
	//var resp models.SysUser
	var userModel = models.SysUser{}
	userModel.UserId = int(in.UserId)
	user, err := userModel.GetUserDetail()
	if err != nil || user.UserId == 0 {
		logx.Errorf("GetUserDetail error, %s", err.Error())
		return nil, err
	}

	var ub client.UserBasic
	err = copier.Copy(&ub, user)
	if err != nil {
		logx.Errorf("Copy &pb.UserBasic error, %s", err.Error())
		return nil, err
	}

	return &pb.UserDetailResp{
		UserBasic: &ub,
	}, nil
}
