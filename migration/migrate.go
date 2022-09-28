package migration

import (
	communityModel "capstone/happyApp/features/community/data"
	userModel "capstone/happyApp/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(userModel.User{})
	db.AutoMigrate(communityModel.Community{})
	db.AutoMigrate(communityModel.JoinCommunity{})
}
