package controller

import (
	"convy/internal/service/user"
)

func toViewSignupResponse(resp user.CreateUserResponse) SignupResponse {
	return SignupResponse{
		UserModel: UserModel(resp.UserModel),
	}
}

func toSvcGetUserRequest(req LoginRequest) user.GetUserRequest {
	return user.GetUserRequest{
		Username: req.Username,
		Password: req.Password,
	}
}

func toViewLoginResponse(resp user.GetUserResponse) LoginResponse {
	return LoginResponse{
		UserModel: UserModel(resp.UserModel),
	}
}
