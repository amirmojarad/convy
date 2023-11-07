package user_follow

import (
	"context"
	"convy/conf"
	"github.com/sirupsen/logrus"
)

type UserFollowRepository interface {
	Follow(ctx context.Context, followerId, followingId uint) error
	UnFollow(ctx context.Context, followerId, followingId uint) error
}

type UserFollow struct {
	cfg                  *conf.AppConfig
	logger               *logrus.Entry
	userFollowRepository UserFollowRepository
}

func NewUserFollow(cfg *conf.AppConfig, logger *logrus.Entry, userFollowRepository UserFollowRepository) *UserFollow {
	return &UserFollow{
		cfg:                  cfg,
		logger:               logger,
		userFollowRepository: userFollowRepository,
	}
}

func (f UserFollow) Follow(ctx context.Context, req FollowRequest) (FollowResponse, error) {
	err := f.userFollowRepository.Follow(ctx, req.FollowerId, req.FollowingId)
	if err != nil {
		return FollowResponse{}, err
	}
	return FollowResponse{}, nil
}

func (f UserFollow) UnFollow(ctx context.Context, req UnFollowRequest) (UnFollowResponse, error) {
	if err := f.userFollowRepository.UnFollow(ctx, req.FollowerId, req.FollowingId); err != nil {
		return UnFollowResponse{}, err
	}
	return UnFollowResponse{}, nil
}
