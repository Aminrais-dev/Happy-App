package data

import (
	"capstone/happyApp/features/cart"
	community "capstone/happyApp/features/community/data"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Photo       string
	Stock       uint64
	Price       uint64
	CommunityID uint
	Carts       Cart
}

type JoinCommunity struct {
	gorm.Model
	UserID      uint
	CommunityID uint
	Role        string
}

func ToModelCart(userid, productid uint) Cart {
	return Cart{
		UserID:    userid,
		ProductID: productid,
	}
}

func ToCoreCommunity(data community.Community, sum int64) cart.CoreCommunity {
	return cart.CoreCommunity{
		Title:        data.Title,
		Descriptions: data.Descriptions,
		Logo:         data.Logo,
		Members:      sum,
	}
}

func ProductToCart(data Product, cartId uint) cart.CoreCart {
	return cart.CoreCart{
		ID:           cartId,
		ProductID:    data.ID,
		Name:         data.Name,
		Descriptions: data.Description,
		Photo:        data.Photo,
		Price:        int(data.Price),
	}
}
