package service

import (
	"context"
	"convy/conf"
	"convy/internal/errorext"
	"convy/internal/logger"
	"convy/internal/repository"
	"convy/internal/repository/mocks"
	"database/sql"
	"github.com/pkg/errors"
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

func TestUserService_GetUser(t *testing.T) {
	t.Parallel()
	repo := mocks.NewUserRepository(t)

	cfg := &conf.AppConfig{}
	log := logger.GetLogger().WithField("name", "service-user_test")

	testCases := []struct {
		name             string
		expectedError    error
		repoActions      func(userRepository *mocks.UserRepository, req GetUserRequest, expectedError error)
		req              GetUserRequest
		expectedResponse GetUserResponse
	}{
		{
			name:          "when credentials are invalid",
			expectedError: errors.New("email is invalid"),
			repoActions: func(userRepository *mocks.UserRepository, req GetUserRequest, expectedError error) {

			},
			req: GetUserRequest{
				Email:    "email@email.com",
				Username: "usernameeeee",
				Password: "@TestPass123",
			},
			expectedResponse: GetUserResponse{},
		},
		{
			name:             "when req is empty",
			expectedError:    errors.New("email is invalid"),
			req:              GetUserRequest{},
			expectedResponse: GetUserResponse{},
			repoActions: func(repo *mocks.UserRepository, req GetUserRequest, expectedError error) {

			},
		},
		{
			name:          "when user with given credentials already exists",
			expectedError: errorext.NewValidationError("user with email %s already exists", "test@test.com"),
			req: GetUserRequest{
				Email:    "test@test.com",
				Username: "usernameeeee",
				Password: "testPass123",
			},
			repoActions: func(repo *mocks.UserRepository, req GetUserRequest, expectedError error) {
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
						expectedError,
					).Once()
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.repoActions(repo, tt.req, tt.expectedError)
			givenContext := context.Background()

			svc := NewUser(cfg, log, repo)
			response, err := svc.GetUser(givenContext, tt.req)
			assert.NotNil(t, err)
			assert.Equal(t, tt.expectedError.Error(), err.Error())
			assert.Equal(t, uint(0), response.ID)
		})
	}
}
