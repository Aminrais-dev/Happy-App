package data

import (
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
	Carts       []Cart
}

func ToModelCart(userid, productid uint) Cart {
	return Cart{
		UserID:    userid,
		ProductID: productid,
	}
}
