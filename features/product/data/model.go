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
	Carts       []Cart
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

type temp struct {
	ID          uint
	Logo        string
	Title       string
	Description string
	Count       int64
}

type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint
}

func (data *temp) dataComu(role string) product.Comu {
	return product.Comu{
		ID:          data.ID,
		Logo:        data.Logo,
		Title:       data.Title,
		Description: data.Description,
		Count:       data.Count,
		Role:        role,
	}
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

func (data *Product) toCore() product.ProductCore {
	return product.ProductCore{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Photo:       data.Photo,
		Stock:       data.Stock,
		Price:       data.Price,
	}
}

func toListProduct(data []Product) []product.ProductCore {

	var dataProduct []product.ProductCore
	for key := range data {
		dataProduct = append(dataProduct, data[key].toCore())
	}

	return dataProduct
}
