package user

import (
	"time"
)

type UserModel struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Username       string
	Email          string
	HashedPassword string
	FirstName      string
	Lastname       string
	LastLogin      time.Time
}

type CreateUserRequest struct {
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
}

type CreateUserResponse struct {
	UserModel
}

type GetUserRequest struct {
	Email    string
	Username string
	Password string
}

type GetUserResponse struct {
	UserModel
}

type UpdateUserInformationRequest struct {
}

type UpdateUserInformationResponse struct {
}

type UpdatePasswordRequest struct {
}

type UpdatePasswordResponse struct {
}

type GetUserDetailRequest struct {
}

type GetUserDetailResponse struct {
}
