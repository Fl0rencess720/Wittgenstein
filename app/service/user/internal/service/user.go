package service

import (
	"context"

	v1 "github.com/Fl0rencess720/Ayana/api/gateway/user/v1"
	"github.com/Fl0rencess720/Ayana/app/service/user/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	accessToken, refreshToken, err := s.uc.Login(ctx, req.Phone, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
func (s *UserService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	if err := s.uc.Register(ctx, req.Phone, req.Password); err != nil {
		return nil, err
	}
	return &v1.RegisterReply{Message: "success"}, nil
}

func (s *UserService) SetProfile(ctx context.Context, req *v1.SetProfileRequest) (*v1.SetProfileReply, error) {
	if err := s.uc.SetProfile(ctx, req.Phone, biz.Profile{Name: req.Profile.Name, Avatar: req.Profile.Avatar}); err != nil {
		return nil, err
	}
	return &v1.SetProfileReply{}, nil
}

func (s *UserService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*v1.GetProfileReply, error) {
	profile, err := s.uc.GetProfile(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	return &v1.GetProfileReply{Profile: &v1.Profile{Name: profile.Name, Avatar: profile.Avatar}}, nil
}
