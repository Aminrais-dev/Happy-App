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

func (service *Service) GetCommunityFeed(communityid int) (community.CoreCommunity, string, error) {
	core, msg, err := service.do.SelectCommunity(communityid)
	return core, msg, err
}

func (service *Service) JoinCommunity(userid, communityid int) (string, error) {
	msg1, err1 := service.do.CheckJoin(userid, communityid)
	if err1 != nil {
		return msg1, err1
	}

	msg2, err2 := service.do.InsertToJoin(userid, communityid)
	return msg2, err2
}

func (service *Service) PostFeed(core community.CoreFeed) (string, error) {
	msg, err := service.do.InsertFeed(core)
	return msg, err
}

func (service *Service) GetDetailFeed(feedid int) (community.CoreFeed, string, error) {
	core, msg, err := service.do.SelectFeed(feedid)
	return core, msg, err
}
