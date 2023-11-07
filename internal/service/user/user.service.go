package user

import (
	"context"
	"convy/conf"
	"convy/internal/errorext"
	"convy/internal/repository"
	"convy/internal/service"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req repository.UserModel) (repository.UserModel, error)
	GetUser(ctx context.Context, req repository.GetUserRequest) (repository.UserModel, error)
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
	fetchedUser, err := u.userRepository.GetUser(ctx, repository.GetUserRequest{Username: req.Username, Email: req.Email})
	if err != nil {
		return CreateUserResponse{}, err
	}

	if fetchedUser.ID > 0 {
		return CreateUserResponse{},
			errorext.NewValidationError("user with email %s already exists", fetchedUser.Email)
	}

	req.Password, err = service.HashPassword(req.Password)
	if err != nil {
		return CreateUserResponse{}, err
	}

	createdUser, err := u.userRepository.CreateUser(ctx, toRepoUserModel(req))
	if err != nil {
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{UserModel: toSvcUserModel(createdUser)}, nil
}

func (u UserService) GetUser(ctx context.Context, req GetUserRequest) (GetUserResponse, error) {
	_, err := service.NewValidation().
		SetUsername(req.Username).
		SetEmail(req.Email).
		SetPassword(req.Password).
		Validate()

	if err != nil {
		u.logger.Error(err)

		return GetUserResponse{}, err
	}

	fetchedUser, err := u.userRepository.GetUser(ctx, toRepoGetUserRequest(req))
	if err != nil {
		return GetUserResponse{}, err
	}

	if !service.CheckPasswordHash(req.Password, fetchedUser.HashedPassword) {
		return GetUserResponse{}, errorext.NewValidationError("password is invalid")
	}

	return GetUserResponse{
		UserModel: toSvcUserModel(fetchedUser),
	}, nil
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
