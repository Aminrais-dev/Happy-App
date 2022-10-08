package data

import (
	"capstone/happyApp/features/cart"
	community "capstone/happyApp/features/community/data"
	event "capstone/happyApp/features/event/data"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Title        string
	Descriptions string
	Logo         string
	Member       []JoinCommunity
	Feeds        []Feed
	Products     []Product
}
type User struct {
	gorm.Model
	Name      string
	Username  string `gorm:"unique"`
	Gender    string
	Email     string `gorm:"unique"`
	Password  string
	Photo     string
	Status    string
	Community []JoinCommunity
	Feeds     []Feed
	Comments  []Comment
	Event     []event.JoinEvent
	Carts     []Cart
}

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
	OrderID          string
	Gross            string
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
type Feed struct {
	gorm.Model
	Text        string
	UserID      uint
	CommunityID uint
	Comments    []Comment
}
type Comment struct {
	gorm.Model
	Text   string
	FeedID uint
	UserID uint
}

type History struct {
	ProductID uint
	UserID    uint
	Name      string
	Photo     string
	Price     uint
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

func ToCoreProductResponse(data History, name string) cart.CoreProductResponse {
	return cart.CoreProductResponse{
		ID:    data.ProductID,
		Name:  data.Name,
		Photo: data.Photo,
		Price: int(data.Price),
		Buyer: name,
	}
}
