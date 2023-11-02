package repository

import (
	"convy/internal/service"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	gorm.Model
	Username       string
	HashedPassword string
	FirstName      string
	Lastname       string
	LastLogin      time.Time
}

func (UserModel) TableName() string {
	return "users "
}

func (u UserModel) castToUserModelSvc() service.UserModel {
	return service.UserModel{
		ID:             u.ID,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
		Username:       u.Username,
		HashedPassword: u.HashedPassword,
		FirstName:      u.FirstName,
		Lastname:       u.Lastname,
		LastLogin:      u.LastLogin,
	}
}

func NewFromUserModelSvc(userModelSvc service.UserModel) *UserModel {
	return &UserModel{
		Username:       userModelSvc.Username,
		HashedPassword: userModelSvc.HashedPassword,
		FirstName:      userModelSvc.FirstName,
		Lastname:       userModelSvc.Lastname,
		LastLogin:      userModelSvc.LastLogin,
	}
}
