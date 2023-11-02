package controller

import "time"

type UserModel struct {
	ID             uint      `json:"id,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Username       string    `json:"username,omitempty"`
	Email          string    `json:"email,omitempty"`
	HashedPassword string    `json:"hashed_password,omitempty"`
	FirstName      string    `json:"first_name,omitempty"`
	Lastname       string    `json:"lastname,omitempty"`
	LastLogin      time.Time `json:"last_login"`
}

type SignupRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type SignupResponse struct {
	UserModel
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserModel
}
