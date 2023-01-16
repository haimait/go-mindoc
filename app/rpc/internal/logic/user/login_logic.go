package userlogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/haimait/go-mindoc/app/rpc/internal/svc"
	"github.com/haimait/go-mindoc/app/rpc/pb-desc/types/pb"
	"github.com/haimait/go-mindoc/helper"
	"github.com/haimait/go-mindoc/models"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	//uc models.UserModel
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// 1、search userInfo
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

	//2、Generate the token
	token, err := helper.GenerateToken(user.UserId, user.Username, l.svcCtx.Config.JwtAuth.JwtKey, l.svcCtx.Config.JwtAuth.TakenExpire)
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		err = errors.New("用户名或密码不正确")
		return nil, err
	}

	// 3、生成用于刷新token的token
	refreshToken, err := helper.GenerateToken(user.UserId, user.Username, l.svcCtx.Config.JwtAuth.JwtKey, l.svcCtx.Config.JwtAuth.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil
}

func (l *LoginLogic) loginByMobile(AuthKey, password string) (u *models.SysUser, err error) {
	var resp = models.SysUser{}
	var userModel = models.SysUser{}

	//方法一
	user, err := userModel.GetUserDetailWithLogin(AuthKey)

	//方法二 直接查mysql
	//err := l.svcCtx.DB.Where("phone = ? or username =? AND  status = '2' ",
	//	AuthKey, AuthKey).First(userModel).Error

	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return &resp, errors.New("用户名或密码不正确")
	}
	_, err = helper.CompareHashAndPassword(user.Password, password)
	if err != nil {
		logx.Errorf("user login error, %s", err.Error())
		return &resp, errors.New("用户名或密码不正确")
	}
	err = copier.Copy(&resp, user)
	return &resp, nil

}
