package controller

import (
	"convy/conf"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PrivateChatService interface {
}

type PrivateChat struct {
	logger *logrus.Entry
	cfg    *conf.AppConfig
	svc    PrivateChatService
}

func NewPrivateChat(cfg *conf.AppConfig, logger *logrus.Entry, svc PrivateChatService) *PrivateChat {
	return &PrivateChat{
		logger: logger,
		cfg:    cfg,
		svc:    svc,
	}
}

func (pc PrivateChat) CreatePrivateChat(ctx *gin.Context) {

}

func (pc PrivateChat) GetUserPrivateChats(ctx *gin.Context) {

}

func (pc PrivateChat) DeletePrivateChat(ctx *gin.Context) {
	
}
