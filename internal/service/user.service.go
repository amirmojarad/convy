package service

import (
	"context"
	"convy/conf"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req UserModel) (UserModel, error)
	GetUser(ctx context.Context, req GetUserRequest) (UserModel, error)
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
	fetchedUser, err := u.userRepository.GetUser(ctx, GetUserRequest{Username: req.Username, Email: req.Email})
	if err != nil {
		return CreateUserResponse{}, err
	}

	if fetchedUser.ID > 0 {
		return CreateUserResponse{}, err
	}

	req.Password, err = HashPassword(req.Password)
	if err != nil {
		return CreateUserResponse{}, err
	}

	createdUser, err := u.userRepository.CreateUser(ctx, toUserModel(req))
	if err != nil {
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{UserModel: createdUser}, nil
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
