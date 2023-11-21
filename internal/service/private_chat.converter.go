package service

import (
	"convy/internal/repository"
)

func toRepoCreatePrivateChatRequest(req CreateRequest) repository.CreatePrivateChatRequest {
	return repository.CreatePrivateChatRequest{
		FirstUserId:  req.FirstUserId,
		SecondUserId: req.SecondUserId,
	}
}
