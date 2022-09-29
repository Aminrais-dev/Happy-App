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

func (service *Service) AddNewCommunity(userid int, core community.CoreCommunity) (string, error) {
	msg, err := service.do.Insert(userid, core)
	return msg, err
}

func (service *Service) GetListCommunity() ([]community.CoreCommunity, string, error) {
	listcore, msg, err := service.do.SelectList()
	return listcore, msg, err
}

func (service *Service) GetMembers(communityid int) ([]string, string, error) {
	members, msg, err := service.do.SelectMembers(communityid)
	return members, msg, err
}

func (service *Service) Leave(userid, communityid int) (string, error) {
	msg, err := service.do.Delete(userid, communityid)
	return msg, err
}
