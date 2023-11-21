package repository

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u User) CreateUser(ctx context.Context, req UserModel) (UserModel, error) {
	var userModel UserModel

	if err := u.db.WithContext(ctx).Model(&UserModel{}).Create(&userModel).Error; err != nil {
		return UserModel{}, err
	}

	return userModel, nil
}

func (u User) GetUser(ctx context.Context, req GetUserRequest) (UserModel, error) {
	var userModel UserModel

	query := u.db.WithContext(ctx).Model(&UserModel{})

	if len(req.Email) != 0 {
		query = query.Where("email = ?", req.Email)
	}

	if len(req.Username) != 0 {
		query = query.Where("username = ?", req.Username)
	}

	if req.Id > 0 {
		query = query.Where("id = ?", req.Id)
	}

	if err := query.Find(&userModel).Error; err != nil {
		return UserModel{}, err
	}

	return userModel, nil
}
