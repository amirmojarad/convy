package service

import (
	"context"
	"convy/conf"
	"convy/internal/errorext"
	"convy/internal/logger"
	"convy/internal/repository"
	"convy/internal/repository/mocks"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	t.Parallel()
	repo := mocks.NewUserRepository(t)

	cfg := &conf.AppConfig{}
	log := logger.GetLogger().WithField("name", "service-user_test")

	testCases := []struct {
		name             string
		expectedError    error
		repoActions      func(userRepository *mocks.UserRepository, req CreateUserRequest, expectedError error)
		req              CreateUserRequest
		expectedResponse CreateUserResponse
	}{
		{
			name:             "when req is empty",
			expectedError:    sql.ErrNoRows,
			req:              CreateUserRequest{},
			expectedResponse: CreateUserResponse{},
			repoActions: func(repo *mocks.UserRepository, req CreateUserRequest, expectedError error) {
				repo.On("GetUser",
					context.Background(),
					repository.GetUserRequest{
						Email:    req.Email,
						Username: req.Username,
					}).
					Return(
						repository.UserModel{},
						sql.ErrNoRows,
					).Once()
			},
		},
		{
			name:          "when user with given credentials already exists",
			expectedError: errorext.NewValidationError("user with email %s already exists", "test@test.com"),
			req: CreateUserRequest{
				Email:    "test@test.com",
				Username: "testUsername",
			},
			repoActions: func(repo *mocks.UserRepository, req CreateUserRequest, expectedError error) {
				repo.On("GetUser",
					context.Background(),
					repository.GetUserRequest{
						Email:    req.Email,
						Username: req.Username,
					}).
					Return(
						repository.UserModel{
							Model: gorm.Model{
								ID: 1,
							},
							Username: req.Username,
							Email:    req.Email,
						},
						nil,
					).Once()
			},
		},
		{
			name: "when "
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.repoActions(repo, tt.req, tt.expectedError)
			givenContext := context.Background()

			svc := NewUser(cfg, log, repo)
			response, err := svc.CreateUser(givenContext, tt.req)
			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), tt.expectedError.Error())
			assert.Equal(t, uint(0), response.ID)
		})
	}
}
