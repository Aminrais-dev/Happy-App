package usecase

import (
	"capstone/happyApp/features/user"

	"golang.org/x/crypto/bcrypt"
)

type usecaseInterface struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &usecaseInterface{
		userData: data,
	}
}

func (usecase *usecaseInterface) PostUser(data user.CoreUser) int {

	if data.Name == "" || data.Email == "" || data.Gender == "" || data.Password == "" || data.Username == "" {
		return -2
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	data.Password = string(hashPass)

	row := usecase.userData.InsertUser(data)
	return row

}
