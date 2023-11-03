package repository

import "gorm.io/gorm"

type UserFollowModel struct {
	gorm.Model
	Follower  uint
	Following uint
}

func (u UserFollowModel) TableName() string {
	return "user_follow"
}
