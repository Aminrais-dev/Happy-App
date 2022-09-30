package data

import (
	"capstone/happyApp/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Photo       string
	Stock       uint64
	Price       uint64
	CommunityID uint
}

type Community struct {
	gorm.Model
	Title        string
	Descriptions string
	Logo         string
	Member       []JoinCommunity
}

type JoinCommunity struct {
	gorm.Model
	UserID      uint
	CommunityID uint
	Role        string
}

func fromCore(data product.ProductCore) Product {
	return Product{
		Name:        data.Name,
		Description: data.Description,
		Photo:       data.Photo,
		Stock:       data.Stock,
		Price:       data.Price,
		CommunityID: data.CommunityID,
	}
}
