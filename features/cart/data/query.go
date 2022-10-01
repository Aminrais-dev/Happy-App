package data

import (
	"capstone/happyApp/features/cart"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) cart.DataInterface {
	return &Storage{
		query: db,
	}
}

func (storage *Storage) InsertIntoCart(userid, productid int) (string, error) {
	cart := ToModelCart(uint(userid), uint(productid))

	tx := storage.query.Create(&cart)
	if tx.Error != nil {
		return "Gagl Menambahkan Ke Cart", tx.Error
	}

	return "Sukses Menambahkan Ke Cart", nil
}
