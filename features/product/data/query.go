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

func (repo *productData) SelectProduct(idProduct, userId int) (product.Comu, product.ProductCore, error) {

	var ProductComu temp
	tx := repo.db.Model(&Community{}).Select("communities.id as id, communities.logo as logo, communities.title as title, communities.descriptions as description, count(join_communities.user_id) as count").Joins("inner join join_communities on join_communities.community_id = communities.id").Joins("inner join products on products.community_id = communities.id").Where("products.id = ? AND join_communities.deleted_at IS NULL", idProduct).Group("communities.id").Scan(&ProductComu)
	if tx.Error != nil {
		return product.Comu{}, product.ProductCore{}, tx.Error
	}

	var dataProduct Product
	repo.db.First(&dataProduct, "id = ? ", idProduct)

	var role JoinCommunity
	repo.db.First(&role, "user_id = ? AND community_id = ? ", userId, ProductComu.ID)

	return ProductComu.dataComu(role.Role), dataProduct.toCore(), nil

}

func (repo *productData) SelectProductComu(idComu, userId int) (product.Comu, []product.ProductCore, error) {

	var ProductComu temp
	tx := repo.db.Model(&Community{}).Select("communities.id as id, communities.logo as logo, communities.title as title, communities.descriptions as description, count(join_communities.user_id) as count").Joins("inner join join_communities on join_communities.community_id = communities.id").Where("communities.id = ? AND join_communities.deleted_at IS NULL", idComu).Group("communities.id").Scan(&ProductComu)
	if tx.Error != nil {
		return product.Comu{}, nil, tx.Error
	}

	var dataProduct []Product
	repo.db.Find(&dataProduct, "community_id = ? ", idComu)

	var role JoinCommunity
	repo.db.First(&role, "user_id = ? AND community_id = ? ", userId, ProductComu.ID)

	return ProductComu.dataComu(role.Role), toListProduct(dataProduct), nil

}
