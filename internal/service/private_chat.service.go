package service

import (
	"context"
	"convy/conf"
	repository "convy/internal/repository/private_chat"
	"github.com/sirupsen/logrus"
)

type PrivateChatRepository interface {
	CreatePrivateChat(ctx context.Context, req repository.CreatePrivateChatRequest) (repository.CreatePrivateChatResponse, error)
	GetUsersPrivateChats(ctx context.Context, req repository.GetUserPrivateChatsResponse) (repository.GetUserPrivateChatsResponse, error)
	DeletePrivateChat(ctx context.Context, req repository.DeletePrivateChatRequest) error
}

type PrivateChat struct {
	cfg          *conf.AppConfig
	logger       *logrus.Entry
	prRepository PrivateChatRepository
}

func NewPrivateChat(cfg *conf.AppConfig,
	logger *logrus.Entry,
	prRepository PrivateChatRepository) *PrivateChat {
	return &PrivateChat{
		cfg:          cfg,
		logger:       logger,
		prRepository: prRepository,
	}
}

func (pc PrivateChat) CreatePrivateChat(ctx context.Context, req CreateRequest) (
	CreateResponse, error) {
	if _, err := NewValidation().SetIds(req.SecondUserId, req.FirstUserId).Validate(); err != nil {
		return CreateResponse{}, nil
	}

	privateChat, err := pc.prRepository.CreatePrivateChat(ctx, toRepoCreatePrivateChatRequest(req))
	if err != nil {
		return CreateResponse{}, err
	}

	return CreateResponse{
		Id: privateChat.Id,
	}, nil
}
