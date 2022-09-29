package usecase

import "capstone/happyApp/features/community"

type Service struct {
	do community.DataInterface
}

func New(data community.DataInterface) community.UsecaseInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) GetListCommunity() ([]community.CoreCommunity, string, error) {
	listcore, msg, err := service.do.SelectList()
	return listcore, msg, err
}
