package data

import (
	"capstone/happyApp/features/cart"
	community "capstone/happyApp/features/community/data"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID       uint
	ProductID    uint
	Quantity     int `gorm:"default:1"`
	Transactions []TransactionCart
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

type Transaction struct {
	gorm.Model
	Street           string
	City             string
	Province         string
	Type_Payment     string
	Status_Payment   string `gorm:"default:unpaid"`
	Midtrans_Virtual string
	Carts            []TransactionCart
}

type TransactionCart struct {
	gorm.Model
	TransactionID uint
	CartID        uint
	Price         uint
}

type Payment struct {
	gorm.Model
	UserID  uint
	OrderID string
	Groos   int
}

func ToModelCart(userid, productid uint) Cart {
	return Cart{
		UserID:    userid,
		ProductID: productid,
	}
}

func ToCoreCommunity(data community.Community, sum int64) cart.CoreCommunity {
	return cart.CoreCommunity{
		ID:           data.ID,
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

func ToModelTransaction(data cart.CoreHistory) Transaction {
	return Transaction{
		Street:       data.Street,
		City:         data.City,
		Province:     data.Province,
		Type_Payment: data.Type_Payment,
	}
}

func ToModelTransactionCart(transid, cartid int) TransactionCart {
	return TransactionCart{
		TransactionID: uint(transid),
		CartID:        uint(cartid),
	}
}
