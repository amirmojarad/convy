package controller

import (
	"context"
	"convy/conf"
	"convy/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserFollowService interface {
	UnFollow(ctx context.Context, req service.UnFollowRequest) (service.UnFollowResponse, error)
	Follow(ctx context.Context, req service.FollowRequest) (service.FollowResponse, error)
}

type UserFollow struct {
	logger *logrus.Entry
	cfg    *conf.AppConfig
	svc    UserFollowService
}

func NewUserFollow(cfg *conf.AppConfig, logger *logrus.Entry, userSvc UserFollowService) *UserFollow {
	return &UserFollow{
		logger: logger,
		cfg:    cfg,
		svc:    userSvc,
	}
}

func (f UserFollow) Follow(ctx *gin.Context) {
	var req FollowRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		WriteBindingErrorResponse(ctx, err)

		return
	}

	response, err := f.svc.Follow(ctx, service.FollowRequest(req))
	if err != nil {
		return
	}

	ctx.JSON(http.StatusCreated, FollowResponse(response))
}

func (f UserFollow) UnFollow(ctx *gin.Context) {
	var req UnFollowRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		WriteBindingErrorResponse(ctx, err)

		return
	}

	response, err := f.svc.UnFollow(ctx, service.UnFollowRequest(req))
	if err != nil {
		return
	}

	ctx.JSON(http.StatusCreated, FollowResponse(response))
}
