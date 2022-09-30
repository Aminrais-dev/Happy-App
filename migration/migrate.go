package migration

import (
	communityModel "capstone/happyApp/features/community/data"
	eventModel "capstone/happyApp/features/event/data"
	userModel "capstone/happyApp/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(userModel.User{})
	db.AutoMigrate(communityModel.Community{})
	db.AutoMigrate(communityModel.JoinCommunity{})
	db.AutoMigrate(eventModel.Event{})
	db.AutoMigrate(eventModel.JoinEvent{})
	db.AutoMigrate(communityModel.Feed{})
	db.AutoMigrate(communityModel.Comment{})
}
