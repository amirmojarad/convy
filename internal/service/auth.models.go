package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
	AtExpires    time.Time
	RtExpires    time.Time
}

type CreateTokenPairRequest struct {
	UserClaims
}

type CreateTokenPairResponse struct {
	Tokens
}

type ValidateTokenResponse struct {
	Token  *jwt.Token
	Valid  bool
	Claims jwt.MapClaims
}

type ValidateTokenRequest struct {
	Token string
}

type RefreshTokensResponse struct {
	Tokens
}

type RefreshTokensRequest struct {
	RefreshToken string
	UserId       uint
}
