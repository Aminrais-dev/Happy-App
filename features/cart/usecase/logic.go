package usecase

import (
	"capstone/happyApp/features/cart"
)

type Service struct {
	do cart.DataInterface
}

func New(data cart.DataInterface) cart.UsecaseInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) AddToCart(userid, productid int) (string, error) {
	msg, err := service.do.InsertIntoCart(userid, productid)
	return msg, err
}

func (service *Service) GetCartList(userid, communityid int) (cart.CoreCommunity, []cart.CoreCart, string, error) {
	corecommunity, msg1, ers := service.do.GetCommunity(communityid)
	if ers != nil {
		return cart.CoreCommunity{}, nil, msg1, ers
	}

	listcart, msg, err := service.do.SelectCartList(userid, communityid)
	return corecommunity, listcart, msg, err
}
