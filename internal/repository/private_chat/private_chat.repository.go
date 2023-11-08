package private_chat

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PrivateChat struct {
	db *gorm.DB
}

func NewPrivateChat(db *gorm.DB) *PrivateChat {
	return &PrivateChat{
		db: db,
	}
}

func (pc PrivateChat) CreatePrivateChat(ctx context.Context, req CreatePrivateChatRequest) (
	CreatePrivateChatResponse, error) {
	pcModel := PrivateChatModel{
		FirstUser:  req.FirstUserId,
		SecondUser: req.SecondUserId,
	}

	if err := pc.db.WithContext(ctx).Create(&pcModel).Error; err != nil {
		return CreatePrivateChatResponse{}, errors.WithStack(err)
	}

	return CreatePrivateChatResponse{Id: pcModel.ID}, nil
}

func (pc PrivateChat) GetUsersPrivateChats(ctx context.Context, req GetUserPrivateChatsResponse) (
	GetUserPrivateChatsResponse, error) {
	return GetUserPrivateChatsResponse{}, nil
}

func (pc PrivateChat) DeletePrivateChat(ctx context.Context, req DeletePrivateChatRequest) error {
	return nil
}
