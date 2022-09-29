package data

import (
	"capstone/happyApp/features/user"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}
}

func (repo *userData) InsertUser(data user.CoreUser) int {

	newData := fromCore(data)
	tx := repo.db.Create(&newData)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return -1
	}

	return 1

}

func (repo *userData) DelUser(id int) int {

	tx := repo.db.Where("id = ? ", id).Delete(&User{})
	if tx.Error != nil {
		return -1
	}

	return 1

}

func (repo *userData) UpdtUser(data user.CoreUser) int {

	newData := fromCore(data)

	tx := repo.db.Model(&User{}).Where("id = ? ", int(data.ID)).Updates(newData)
	if tx.Error != nil {
		return -1
	}

	return int(tx.RowsAffected)

}
