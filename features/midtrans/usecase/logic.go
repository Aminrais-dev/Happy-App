package usecase

import "capstone/happyApp/features/midtrans"

type Service struct {
	do midtrans.DataInterface
}

func New(data midtrans.DataInterface) midtrans.UsecaseInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) WeebHookTransaction(orderid, transactionstatus string) (string, error) {
	msg, err := service.do.WeebHookUpdate(orderid, transactionstatus)
	return msg, err
}
