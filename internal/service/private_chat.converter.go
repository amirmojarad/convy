package service

import repository "convy/internal/repository/private_chat"

func toRepoCreatePrivateChatRequest(req CreateRequest) repository.CreatePrivateChatRequest {
	return repository.CreatePrivateChatRequest{
		FirstUserId:  req.FirstUserId,
		SecondUserId: req.SecondUserId,
	}
}
