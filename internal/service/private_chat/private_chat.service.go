package private_chat

import (
	"context"
	"convy/conf"
	"github.com/sirupsen/logrus"
)

type PrivateChatRepository struct {
}

type PrivateChatMessageRepository struct {
}

type PrivateChat struct {
	cfg           *conf.AppConfig
	logger        *logrus.Entry
	prRepository  PrivateChatRepository
	prmRepository PrivateChatMessageRepository
}

func NewPrivateChat(cfg *conf.AppConfig,
	logger *logrus.Entry,
	prRepository PrivateChatRepository,
	prmRepository PrivateChatMessageRepository) *PrivateChat {
	return &PrivateChat{
		cfg:           cfg,
		logger:        logger,
		prmRepository: prmRepository,
		prRepository:  prRepository,
	}
}

func (pc PrivateChat) CreatePrivateChat(ctx context.Context, req CreateRequest) (
	CreateResponse, error) {
	return CreateResponse{}, nil
}

func (pc PrivateChat) GetAllMessages(ctx context.Context, req GetAllMessagesRequest) (GetAllMessagesResponse, error) {
	return GetAllMessagesResponse{}, nil
}

func (pc PrivateChat) AddMessage(ctx context.Context, req AddMessageRequest) (AddMessageResponse, error) {
	return AddMessageResponse{}, nil
}
