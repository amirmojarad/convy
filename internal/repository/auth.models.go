package repository

import "gorm.io/gorm"

type TokenModel struct {
	gorm.Model
	UserId      uint
	HashedToken string
}

func (m TokenModel) TableName() string {
	return "tokens"
}

type GetRefreshTokenResponse struct {
}

type GetRefreshTokenRequest struct {
}

type SetRefreshTokenRequest struct {
	EncryptedRT string
	UserId      uint
}

type SetRefreshTokenResponse struct {
	Id uint
}
