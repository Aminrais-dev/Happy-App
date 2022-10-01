package usecase

import "capstone/happyApp/features/cart"

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
