package controller

import (
	"context"
	"convy/conf"
	"convy/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	GetUserDetail(ctx context.Context, req service.GetUserDetailRequest) (service.GetUserDetailResponse, error)
	UpdatePassword(ctx context.Context, req service.UpdatePasswordRequest) (service.UpdatePasswordResponse, error)
	UpdateUserInformation(ctx context.Context, req service.UpdateUserInformationRequest) (
		service.UpdateUserInformationResponse, error)
	GetUser(ctx context.Context, req service.GetUserRequest) (service.GetUserResponse, error)
	CreateUser(ctx context.Context, req service.CreateUserRequest) (service.CreateUserResponse, error)
}

type User struct {
	logger *logrus.Entry
	cfg    *conf.AppConfig
	svc    UserService
}

func NewUser(cfg *conf.AppConfig, logger *logrus.Entry, userSvc UserService) *User {
	return &User{
		logger: logger,
		cfg:    cfg,
		svc:    userSvc,
	}
}

func (u User) Signup(ctx *gin.Context) {

}

func (u User) Login(ctx *gin.Context) {

}
