package repository

import (
	"context"
	"convy/internal/errorext"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

const (
	MongoDbName = "convy-mongo"
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

func (pc PrivateChat) IsExists(ctx context.Context, id uint) (bool, error) {
	privateChat, err := pc.GetById(ctx, id)
	if err != nil {
		return false, err
	}

	if privateChat.ID > 0 {
		return true, nil
	}

	return false, errors.WithStack(errorext.NewDatabaseError("private chat with id %d not found", id))
}

func (pc PrivateChat) GetById(ctx context.Context, id uint) (*PrivateChatModel, error) {
	var privateChat PrivateChatModel

	if err := pc.psqlDb.WithContext(ctx).Where("id = ?", id).Find(&privateChat); err != nil {
		return nil, nil
	}

	return &privateChat, nil
}

func (pc PrivateChat) getCollection(_ context.Context, message PrivateChatMessage) *mongo.Collection {
	return pc.mongoDb.Database(MongoDbName).Collection(message.Collection())
}

func (pc PrivateChat) Message(ctx context.Context, req MessageRequest) error {
	if _, err := pc.IsExists(ctx, req.PrivateChatId); err != nil {
		return err
	}

	_, err := pc.getCollection(nil, req.PrivateChatMessage).InsertOne(ctx, req.PrivateChatMessage)
	if err != nil {
		return err
	}

	return nil
}

func (pc PrivateChat) GetMessages(ctx context.Context, req GetMessagesRequest) ([]PrivateChatMessage, error) {
	return nil, nil
}
