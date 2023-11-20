package private_chat

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type PrivateChat struct {
	psqlDb  *gorm.DB
	mongoDb *mongo.Client
}

func NewPrivateChat(psqlDb *gorm.DB, mongoDb *mongo.Client) *PrivateChat {
	return &PrivateChat{
		psqlDb:  psqlDb,
		mongoDb: mongoDb,
	}
}

func (pc PrivateChat) CreatePrivateChat(ctx context.Context, req CreatePrivateChatRequest) (
	CreatePrivateChatResponse, error) {
	pcModel := PrivateChatModel{
		FirstUser:  req.FirstUserId,
		SecondUser: req.SecondUserId,
	}

	if err := pc.psqlDb.WithContext(ctx).Create(&pcModel).Error; err != nil {
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
