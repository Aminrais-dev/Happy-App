package data

import (
	"capstone/happyApp/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
}

func toCore(user User) login.Core {

	var core = login.Core{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}

	return core
}
