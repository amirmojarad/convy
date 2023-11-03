package repository

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	gorm.Model
	Username       string
	HashedPassword string
	FirstName      string
	Email          string
	Lastname       string
	LastLogin      time.Time
}

func (UserModel) TableName() string {
	return "users "
}

type GetUserRequest struct {
	Email    string
	Username string
}
