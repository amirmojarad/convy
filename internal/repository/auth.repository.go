package repository

import (
	"context"
	"gorm.io/gorm"
)

type Token struct {
	db *gorm.DB
}

func (t Token) SetRefreshToken(ctx context.Context, req SetRefreshTokenRequest) (SetRefreshTokenResponse, error) {
	return SetRefreshTokenResponse{}, nil
}

func (t Token) GetRefreshToken(ctx context.Context, req GetRefreshTokenRequest) (GetRefreshTokenResponse, error) {
	return GetRefreshTokenResponse{}, nil
}
