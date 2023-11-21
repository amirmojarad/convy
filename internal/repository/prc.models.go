package repository

import "gorm.io/gorm"

type PCMessageModel struct {
	gorm.Model
	SenderId      uint
	ReceiverId    uint
	Message       string
	PrivateChatId uint
}

func (p PCMessageModel) TableName() string {
	return "pr_chat_messages"
}

type PrivateChatMessage struct {
	SenderId      uint   `bson:"sender_id,omitempty"`
	ReceiverId    uint   `bson:"receiver_id,omitempty"`
	PrivateChatId uint   `bson:"private_chat_id,omitempty"`
	Message       string `bson:"message,omitempty"`
}

func (pcm PrivateChatMessage) Collection() string {
	return "private_chat_messages"
}

type MessageRequest struct {
	PrivateChatMessage
}

type GetMessagesRequest struct {
	PrivateChatId uint
}

type GetMessagesResponse struct {
}
