package private_chat

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
