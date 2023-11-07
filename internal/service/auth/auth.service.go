package auth

import (
	"context"
	"convy/conf"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	cfg *conf.AppConfig
}

func NewToken(cfg *conf.AppConfig) *Token {
	return &Token{
		cfg: cfg,
	}
}

// getExpiresDays returns RefreshToken Expires Days, AccessToken Expires Days
func (t Token) getExpiresDaysDurations() (time.Duration, time.Duration) {
	return time.Hour * time.Duration(t.cfg.TokenDetails.RtExpiresDays) * 24,
		time.Hour * time.Duration(t.cfg.TokenDetails.AtExpiresDays) * 24
}

func (t Token) toUserClaims(userId uint, expiresTime time.Time) UserClaims {
	return UserClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    t.cfg.TokenDetails.Issuer,
		},
	}
}

func (t Token) CreateTokenPair(_ context.Context, req CreateTokenPairRequest) (CreateTokenPairResponse, error) {
	var response CreateTokenPairResponse
	rt, at := t.getExpiresDaysDurations()

	response.AtExpires = time.Now().Add(at)
	response.RtExpires = time.Now().Add(rt)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, t.toUserClaims(req.UserId, response.AtExpires))
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, t.toUserClaims(req.UserId, response.RtExpires))

	if acToken, err := accessToken.SignedString([]byte(t.cfg.TokenDetails.Secret)); err != nil {
		return CreateTokenPairResponse{}, err
	} else {
		response.AccessToken = acToken
	}

	if rtToken, err := refreshToken.SignedString([]byte(t.cfg.TokenDetails.Secret)); err != nil {
		return CreateTokenPairResponse{}, err
	} else {
		response.RefreshToken = rtToken
	}

	return response, nil
}

func (t Token) ValidateToken(_ context.Context, req ValidateTokenRequest) (ValidateTokenResponse, error) {
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(t.cfg.TokenDetails.Secret), nil
	})
	if err != nil {
		return ValidateTokenResponse{}, err
	}

	return ValidateTokenResponse{
		Token:  token,
		Valid:  true,
		Claims: token.Claims.(jwt.MapClaims),
	}, nil
}

func (t Token) RefreshTokens(_ context.Context, req RefreshTokensRequest) (RefreshTokensResponse, error) {
	validateTokenResponse, err := t.ValidateToken(nil, ValidateTokenRequest{Token: req.RefreshToken})
	if err != nil {
		return RefreshTokensResponse{}, err
	}

	token := validateTokenResponse.Token

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if claims["user_id"] != req.UserId {
			return RefreshTokensResponse{}, errors.New("invalid refresh token")
		}
	}

	createTokensResponse, err := t.CreateTokenPair(nil, CreateTokenPairRequest{UserClaims{
		UserId:         req.UserId,
		StandardClaims: jwt.StandardClaims{},
	}})
	if err != nil {
		return RefreshTokensResponse{}, err
	}

	return RefreshTokensResponse(createTokensResponse), nil
}
