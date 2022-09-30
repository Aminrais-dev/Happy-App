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

func (repo *productData) DelProduct(idProduct, userId int) int {

	var check string
	repo.db.Model(&Community{}).Select("join_communities.role").Joins("inner join products on products.community_id = communities.id").Joins("inner join join_communities on join_communities.community_id = communities.id").Where("products.id = ? AND join_communities.user_id = ? ", idProduct, userId).Scan(&check)
	if check != "admin" {
		return -2
	}

	tx := repo.db.Delete(&Product{}, "id = ? ", idProduct)
	if tx.Error != nil {
		return -1
	}

	return 1

}

func (repo *productData) UpdtProduct(data product.ProductCore, userId int) int {

	var check string
	repo.db.Model(&Community{}).Select("join_communities.role").Joins("inner join products on products.community_id = communities.id").Joins("inner join join_communities on join_communities.community_id = communities.id").Where("products.id = ? AND join_communities.user_id = ? ", data.ID, userId).Scan(&check)
	if check != "admin" {
		return -2
	}

	var newData = fromCore(data)
	tx := repo.db.Model(&Product{}).Where("id = ? ", int(data.ID)).Updates(newData)
	if tx.Error != nil {
		return -1
	}

	return 1

}
