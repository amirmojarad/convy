package service

import (
	"convy/internal/repository"
)

func toRepoUserModel(req CreateUserRequest) repository.UserModel {
	return repository.UserModel{
		Username:       req.Username,
		HashedPassword: req.Password,
		FirstName:      req.FirstName,
		Lastname:       req.LastName,
	}
}

func toSvcUserModel(req repository.UserModel) UserModel {
	return UserModel{
		ID:             req.ID,
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      req.UpdatedAt,
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: req.HashedPassword,
		FirstName:      req.FirstName,
		Lastname:       req.Lastname,
		LastLogin:      req.LastLogin,
	}
}

func toRepoGetUserRequest(req GetUserRequest) repository.GetUserRequest {
	return repository.GetUserRequest{
		Email:    req.Email,
		Username: req.Username,
	}
}
