package repository

import (
	"context"
	"convy/internal/service"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u User) CreateUser(ctx context.Context, req service.UserModel) (service.UserModel, error) {
	userModel := NewFromUserModelSvc(req)

	if err := u.db.WithContext(ctx).Model(&UserModel{}).Create(&userModel).Error; err != nil {
		return service.UserModel{}, err
	}

	return userModel.castToUserModelSvc(), nil
}

func (u User) GetUser(ctx context.Context, req service.GetUserRequest) (service.UserModel, error) {
	var userModel UserModel

	query := u.db.WithContext(ctx).Model(&UserModel{})

	if len(req.Email) != 0 {
		query = query.Where("email = ?", req.Email)
	}

	if len(req.Username) != 0 {
		query = query.Where("username = ?", req.Username)
	}

	if err := query.Find(&userModel).Error; err != nil {
		return service.UserModel{}, err
	}

	return userModel.castToUserModelSvc(), nil
}
