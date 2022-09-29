package usecase

import (
	"capstone/happyApp/features/community"
	"errors"
)

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

func (service *Service) UpdateCommunity(userid int, core community.CoreCommunity) (string, error) {
	role, err := service.do.GetUserRole(userid, int(core.ID))
	if err != nil {
		return "Gagal mendapatkan role user", err
	} else if role != "admin" {
		return "Hanya admin yang bisa mengupdate Community", errors.New("Dont have access")
	}

	msg, errs := service.do.UpdateCommunity(int(core.ID), core)
	return msg, errs
}
