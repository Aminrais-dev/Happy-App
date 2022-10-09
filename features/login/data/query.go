package data

import (
	"capstone/happyApp/features/login"

	"gorm.io/gorm"
)

type loginData struct {
	db *gorm.DB
}

func New(db *gorm.DB) login.DataInterface {
	return &loginData{
		db: db,
	}
}

func (repo *loginData) LoginUser(email string) (login.Core, error) {

	var data User
	txEmail := repo.db.Where("email = ?", email).First(&data)
	if txEmail.Error != nil {
		return login.Core{}, txEmail.Error
	}

	var dataUser = toCore(data)

	return dataUser, nil

}
