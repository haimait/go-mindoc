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
	uc models.UserModel
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

func (l *LoginLogic) loginByMobile(AuthKey, password string) (u *models.SysUser, err error) {
	//var userModel1 = models.SysUser{}
	//userModel1.Username = AuthKey
	//userModel1.Phone = AuthKey

	//var uc models.UserModel
	//uc = &userModel1
	//uc.GetUserInfo()

	var resp = models.SysUser{}
	var userModel = models.SysUser{}
	userModel.Username = AuthKey
	userModel.Phone = AuthKey
	//方法一
	user, err := userModel.GetUserInfo()
	//方法二
	//l.uc = &userModel
	//user, err := l.uc.GetUserInfo()

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
