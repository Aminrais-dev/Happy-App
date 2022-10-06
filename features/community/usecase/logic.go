package usecase

import (
	"capstone/happyApp/features/community"
	"errors"
	"fmt"
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
	_, err1 := service.do.CheckJoin(userid, communityid)
	if err1 == nil {
		return "Hanya Member Dari Community Yang Bisa Leave", errors.New("Anda Belum Masuk Community")
	}

	role, errs := service.do.GetUserRole(userid, communityid)
	if errs != nil {
		return "Gagal mendapatkan role user", errs
	}

	members, msg, err := service.do.Delete(userid, communityid)
	if members == 0 {
		msgdc, errdc := service.do.DeleteCommunity(communityid)
		if errdc != nil {
			return msgdc, errdc
		}
		msg += " " + msgdc
	} else if role == "admin" {
		newadmin, msgcr, errcr := service.do.ChangeAdmin(communityid)
		if errcr != nil {
			return msgcr, errcr
		}
		msg = fmt.Sprintf("Community akan di wariskan ke %s", newadmin)
	}

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

func (service *Service) GetCommunityFeed(userid, communityid int) (community.CoreCommunity, string, error) {
	core, msg, err := service.do.SelectCommunity(userid, communityid)
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
	_, err1 := service.do.CheckJoin(int(core.UserID), int(core.CommunityID))
	if err1 == nil {
		return "Hanya Anggota dari Community Yang bisa Post Feed", errors.New("Anda Belum Masuk Community")
	}

	msg, err := service.do.InsertFeed(core)
	return msg, err
}

func (service *Service) GetDetailFeed(feedid int) (community.CoreFeed, string, error) {
	core, msg, err := service.do.SelectFeed(feedid)
	return core, msg, err
}

func (service *Service) AddComment(core community.CoreComment) (string, error) {
	communityid, errf := service.do.SelectCommunityIdWithFeed(int(core.FeedID))
	if errf != nil {
		return "Gagal Mendapatkan Comunity Id", errf
	}
	_, err1 := service.do.CheckJoin(int(core.UserID), int(communityid))
	if err1 == nil {
		return "Hanya Anggota dari Community yang bisa melakukan Comment", errors.New("Anda Belum Masuk Community")
	}

	msg, err := service.do.InsertComment(core)
	return msg, err
}

func (service *Service) GetListCommunityWithParam(param string) ([]community.CoreCommunity, string, error) {
	list, msg, err := service.do.SelectListCommunityWithParam(param)
	return list, msg, err
}
