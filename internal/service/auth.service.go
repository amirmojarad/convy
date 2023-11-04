package service

import (
	"context"
	"convy/conf"
	"convy/internal/repository"
)

type TokenRepository interface {
	SetRefreshToken(ctx context.Context, req repository.SetRefreshTokenRequest) (repository.SetRefreshTokenResponse, error)
	GetRefreshToken(ctx context.Context, req repository.GetRefreshTokenRequest) (repository.GetRefreshTokenResponse, error)
}

type Token struct {
	cfg             *conf.AppConfig
	tokenRepository TokenRepository
}

func NewToken(cfg *conf.AppConfig, tokenRepository TokenRepository) *Token {
	return &Token{
		cfg:             cfg,
		tokenRepository: tokenRepository,
	}
}

func (t Token) SetRefreshToken(ctx context.Context, req SetRefreshTokenRequest) (SetRefreshTokenResponse, error) {
	encrypt, err := Encrypt([]byte(req.RefreshToken), t.cfg.Secrets.EncryptionKey)
	if err != nil {
		return SetRefreshTokenResponse{}, err
	}

	response, err := t.tokenRepository.SetRefreshToken(ctx, repository.SetRefreshTokenRequest{
		EncryptedRT: encrypt,
		UserId:      req.UserId,
	})

	if err != nil {
		return SetRefreshTokenResponse{}, err
	}

	return SetRefreshTokenResponse(response), nil
}

func (t Token) GetRefreshToken(ctx context.Context, req GetRefreshTokenRequest) (GetRefreshTokenResponse, error) {
	return GetRefreshTokenResponse{}, nil
}
