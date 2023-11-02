package controller

import "convy/internal/service"

func toViewSignupResponse(resp service.CreateUserResponse) SignupResponse {
	return SignupResponse{
		UserModel: UserModel(resp.UserModel),
	}
}

func toSvcGetUserRequest(req LoginRequest) service.GetUserRequest {
	return service.GetUserRequest{
		Username: req.Username,
		Password: req.Password,
	}
}

func toViewLoginResponse(resp service.GetUserResponse) LoginResponse {
	return LoginResponse{
		UserModel: UserModel(resp.UserModel),
	}
}
