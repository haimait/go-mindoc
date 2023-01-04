package logic

import (
	"context"
	"github.com/haimait/go-mindoc/helper"
	"github.com/haimait/go-mindoc/models"
	"github.com/pkg/errors"

	"github.com/haimait/go-mindoc/user/rpc/internal/svc"
	"github.com/haimait/go-mindoc/user/rpc/pb-desc/types/pb"

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

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var user = new(models.SysUser)
	var err error
	switch in.AuthType {
	case models.UserAuthTypeSystem:
		user, err = l.loginByMobile(in.AuthKey, in.Password)
	default:
		return nil, errors.New("不支持类型")
	}
	if err != nil {
		return nil, err
	}

	//2、Generate the token, so that the service doesn't call rpc internally
	token, err := helper.GenerateToken(user.UserId, user.Username, 3600*24*30)
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		err = errors.New("用户名或密码不正确")
		return nil, err
	}

	return &pb.LoginResp{
		AccessToken: token,
		//AccessExpire: tokenResp.AccessExpire,
		//RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
func (l *LoginLogic) loginByMobile(AuthKey, password string) (*models.SysUser, error) {

	//user, err := l.svcCtx.DB.FindOneByMobile(l.ctx, mobile)
	ub := new(models.SysUser)
	err := l.svcCtx.DB.Where("phone = ? or username =? AND  status = '2' ",
		AuthKey, AuthKey).First(ub).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return ub, errors.New("用户名或密码不正确")
	}
	_, err = helper.CompareHashAndPassword(ub.Password, password)
	if err != nil {
		logx.Errorf("user login error, %s", err.Error())
		return ub, errors.New("用户名或密码不正确")
	}
	return ub, nil

}
