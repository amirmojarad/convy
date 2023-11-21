package service

import (
	"context"
	"convy/conf"
	"convy/internal/repository"
	"github.com/sirupsen/logrus"
)

type PrivateChatRepository interface {
	CreatePrivateChat(ctx context.Context, req repository.CreatePrivateChatRequest) (repository.CreatePrivateChatResponse, error)
	GetUsersPrivateChats(ctx context.Context, req repository.GetUserPrivateChatsResponse) (repository.GetUserPrivateChatsResponse, error)
	DeletePrivateChat(ctx context.Context, req repository.DeletePrivateChatRequest) error
	Message(ctx context.Context, req repository.MessageRequest) error
}

type PrivateChat struct {
	cfg            *conf.AppConfig
	logger         *logrus.Entry
	prRepository   PrivateChatRepository
	userRepository UserRepository
}

func NewPrivateChat(cfg *conf.AppConfig,
	logger *logrus.Entry,
	prRepository PrivateChatRepository,
	userRepository UserRepository,
) *PrivateChat {
	return &PrivateChat{
		cfg:            cfg,
		logger:         logger,
		prRepository:   prRepository,
		userRepository: userRepository,
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

func (pc PrivateChat) Message(ctx context.Context, req MessageRequest) error {
	if _, err := pc.userRepository.GetUser(ctx, repository.GetUserRequest{Id: req.SenderId}); err != nil {
		pc.logger.Error(err)

		return err
	}

	if _, err := pc.userRepository.GetUser(ctx, repository.GetUserRequest{Id: req.ReceiverId}); err != nil {
		pc.logger.Error(err)

		return err
	}

	return pc.prRepository.Message(ctx, repository.MessageRequest{
		PrivateChatMessage: repository.PrivateChatMessage(req.PrivateChatMessage),
	})
}
