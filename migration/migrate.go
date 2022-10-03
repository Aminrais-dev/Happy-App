package migration

import (
	cartModel "capstone/happyApp/features/cart/data"
	communityModel "capstone/happyApp/features/community/data"
	eventModel "capstone/happyApp/features/event/data"
	productModel "capstone/happyApp/features/product/data"
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
	db.AutoMigrate(productModel.Product{})
	db.AutoMigrate(cartModel.Cart{})
}

// do nothing
