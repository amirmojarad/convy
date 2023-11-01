package service

import (
	"context"
	"convy/conf"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
}

type UserService struct {
	cfg            *conf.AppConfig
	logger         *logrus.Entry
	userRepository UserRepository
}

func NewUser(cfg *conf.AppConfig, logger *logrus.Entry, userRepository UserRepository) *UserService {
	return &UserService{
		cfg:            cfg,
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u UserService) CreateUser(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error) {
	return CreateUserResponse{}, nil
}

func (u UserService) GetUser(ctx context.Context, req GetUserRequest) (GetUserResponse, error) {
	return GetUserResponse{}, nil
}

func (u UserService) UpdateUserInformation(ctx context.Context, req UpdateUserInformationRequest) (
	UpdateUserInformationResponse, error) {
	return UpdateUserInformationResponse{}, nil
}

func (u UserService) UpdatePassword(ctx context.Context, req UpdatePasswordRequest) (UpdatePasswordResponse, error) {
	return UpdatePasswordResponse{}, nil
}

func (u UserService) GetUserDetail(ctx context.Context, req GetUserDetailRequest) (GetUserDetailResponse, error) {
	return GetUserDetailResponse{}, nil
}
