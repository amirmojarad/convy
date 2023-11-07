package auth

import (
	"convy/conf"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	cfg = conf.AppConfig{
		TokenDetails: struct {
			Issuer        string
			Secret        string
			AtExpiresDays int
			RtExpiresDays int
		}{
			Issuer:        "issuer",
			Secret:        "secret",
			AtExpiresDays: 1,
			RtExpiresDays: 2,
		},
	}
	timeFormat = "2006-04-06"
	userId     = "user_id"
)

func TestToken_CreateTokenPair(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		req           CreateTokenPairRequest
		expectedError error
		AtExpireDate  time.Time
		RtExpireDate  time.Time
	}{
		{
			name:          "when create tokens successfully, no errors!",
			req:           CreateTokenPairRequest{UserClaims{UserId: 1}},
			expectedError: nil,
			AtExpireDate:  time.Now().Add(time.Hour * time.Duration(cfg.TokenDetails.AtExpiresDays) * 24),
			RtExpireDate:  time.Now().Add(time.Hour * time.Duration(cfg.TokenDetails.RtExpiresDays) * 24),
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tokenSvc := NewToken(&cfg)

			response, err := tokenSvc.CreateTokenPair(nil, tt.req)

			assert.Equal(t, err, tt.expectedError)
			assert.Equal(t, tt.AtExpireDate.Format(timeFormat), response.AtExpires.Format(timeFormat))
			assert.Equal(t, tt.RtExpireDate.Format(timeFormat), response.RtExpires.Format(timeFormat))
		})
	}

}

func TestToken_ValidateToken(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		res           ValidateTokenResponse
		createToken   func() CreateTokenPairResponse
		expectedError error
	}{
		{
			name:          "when token is invalid",
			res:           ValidateTokenResponse{Valid: false, Claims: jwt.MapClaims{userId: 1}},
			expectedError: jwt.ErrSignatureInvalid,
			createToken: func() CreateTokenPairResponse {
				ctResponse, err := NewToken(&conf.AppConfig{TokenDetails: struct {
					Issuer        string
					Secret        string
					AtExpiresDays int
					RtExpiresDays int
				}{
					Issuer:        "invalid_issuer",
					Secret:        "invalid_secret",
					AtExpiresDays: 45,
					RtExpiresDays: 89,
				}}).CreateTokenPair(nil, CreateTokenPairRequest{UserClaims{UserId: 5}})
				assert.Nil(t, err)

				return ctResponse
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tokenSvc := NewToken(&cfg)

			ctResponse := tt.createToken()
			response, err := tokenSvc.ValidateToken(nil, ValidateTokenRequest{Token: ctResponse.AccessToken})

			assert.Equal(t, tt.expectedError.Error(), err.Error())
			assert.Equal(t, tt.res.Valid, response.Valid)
			assert.NotEqual(t, float64(tt.res.Claims[userId].(int)), response.Claims[userId])
			assert.Nil(t, response.Token)
		})
	}
}
