package private_chat

import "gorm.io/gorm"

type PrivateChatModel struct {
	gorm.Model
	// FirstUser starter of chat
	FirstUser uint
	// SecondUser receiver or another participant in the chat
	SecondUser uint
}

func (p PrivateChatModel) TableName() string {
	return "private_chat"
}

type CreatePrivateChatRequest struct {
	FirstUserId  uint
	SecondUserId uint
}

type CreatePrivateChatResponse struct {
	Id uint
}

type GetUserPrivateChatsRequest struct {
	UserId uint
}

type GetUserPrivateChatsResponse struct {
	Results []PrivateChatModel
}

type DeletePrivateChatRequest struct {
	Id uint
}
