package data

import (
	"capstone/happyApp/features/product"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.DataInterface {
	return &productData{
		db: db,
	}
}

func (repo *productData) InsertProduct(data product.ProductCore, userId int) int {

	var check JoinCommunity
	repo.db.First(&check, "community_id = ? AND user_id = ? ", data.CommunityID, userId)

	if check.Role != "admin" {
		return -2
	}

	newData := fromCore(data)
	tx := repo.db.Create(&newData)
	if tx.Error != nil {
		return -1
	}

	return 1

}
