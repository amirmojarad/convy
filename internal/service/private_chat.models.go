package service

type PrivateChatMessage struct {
	SenderId      uint
	ReceiverId    uint
	PrivateChatId uint
	Message       string
}

type CreateRequest struct {
	// FirstUserId user that request create the chat
	FirstUserId uint
	// SecondUserId another user that participate in chat
	SecondUserId uint
}

type CreateResponse struct {
	Id uint
}

type AddMessageRequest struct {
	PrivateChatId  uint
	Message        string
	SenderUserId   uint
	ReceiverUserId uint
}

type AddMessageResponse struct {
	Id      uint
	Message string
}

type MessageRequest struct {
	PrivateChatMessage
}

type GetMessagesRequest struct {
	PrivateChatId uint
}

type GetMessagesResponse struct {
}
