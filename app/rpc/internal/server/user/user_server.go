// Code generated by goctl. DO NOT EDIT.
// Source: core.proto

package server

import (
	"context"

	"github.com/haimait/go-mindoc/app/rpc/internal/logic/user"
	"github.com/haimait/go-mindoc/app/rpc/internal/svc"
	"github.com/haimait/go-mindoc/app/rpc/pb-desc/types/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := userlogic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := userlogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := userlogic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserServer) UpdateUserInfo(ctx context.Context, in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	l := userlogic.NewUpdateUserInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserInfo(in)
}
