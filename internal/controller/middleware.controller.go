package controller

import (
	"context"
	"convy/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	XAuthHeader = "Authorization"
	UserId      = "user_id"
)

type TokenService interface {
	CreateTokenPair(ctx context.Context, req service.CreateTokenPairRequest) (service.CreateTokenPairResponse, error)
	ValidateToken(ctx context.Context, req service.ValidateTokenRequest) (service.ValidateTokenResponse, error)
	RefreshTokens(ctx context.Context, req service.RefreshTokensRequest) (service.RefreshTokensResponse, error)
}

type Middleware struct {
	tokenService TokenService
}

func NewMiddleware(tokenService TokenService) *Middleware {
	return &Middleware{
		tokenService: tokenService,
	}
}

func (m Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(XAuthHeader)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()

			return
		}

		validateTokenResponse, err := m.tokenService.ValidateToken(nil, service.ValidateTokenRequest{Token: tokenString})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()

			return
		}

		if validateTokenResponse.Valid {
			c.Set(UserId, validateTokenResponse.Claims[UserId])

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})

			c.Abort()
		}
	}
}
