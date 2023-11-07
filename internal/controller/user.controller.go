package controller

import (
	"context"
	"convy/conf"
	"convy/internal/service/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserService interface {
	GetUserDetail(ctx context.Context, req user.GetUserDetailRequest) (user.GetUserDetailResponse, error)
	UpdatePassword(ctx context.Context, req user.UpdatePasswordRequest) (user.UpdatePasswordResponse, error)
	UpdateUserInformation(ctx context.Context, req user.UpdateUserInformationRequest) (
		user.UpdateUserInformationResponse, error)
	GetUser(ctx context.Context, req user.GetUserRequest) (user.GetUserResponse, error)
	CreateUser(ctx context.Context, req user.CreateUserRequest) (user.CreateUserResponse, error)
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
	var req SignupRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		WriteBindingErrorResponse(ctx, err)

		return
	}

	response, err := u.svc.CreateUser(ctx.Request.Context(), user.CreateUserRequest(req))
	if err != nil {
		WriteErrorResponse(ctx, err, u.logger)

		return
	}

	ctx.JSON(http.StatusCreated, toViewSignupResponse(response))
}

func (u User) Login(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		WriteBindingErrorResponse(ctx, err)

		return
	}

	response, err := u.svc.GetUser(ctx, toSvcGetUserRequest(req))
	if err != nil {
		WriteErrorResponse(ctx, err, u.logger)

		return
	}

	ctx.JSON(http.StatusOK, toViewLoginResponse(response))
}
