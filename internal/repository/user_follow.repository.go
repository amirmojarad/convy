package repository

import (
	"context"
	"gorm.io/gorm"
)

type UserFollow struct {
	db *gorm.DB
}

func NewUserFollow(db *gorm.DB) *UserFollow {
	return &UserFollow{db: db}
}

func (f UserFollow) Follow(ctx context.Context, followerId, followingId uint) error {
	return f.db.Model(&UserFollowModel{}).Create(&UserFollowModel{
		Follower:  followerId,
		Following: followingId,
	}).Error
}

func (f UserFollow) UnFollow(ctx context.Context, followerId, followingId uint) error {
	f.db.Model(&UserFollowModel{}).Delete(&UserFollowModel{
		Follower:  followerId,
		Following: followingId,
	})
	return nil
}
