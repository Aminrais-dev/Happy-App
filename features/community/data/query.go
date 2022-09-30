package data

import (
	"capstone/happyApp/features/community"
	user "capstone/happyApp/features/user/data"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) community.DataInterface {
	return &Storage{
		query: db,
	}
}

func (storage *Storage) Insert(userid int, core community.CoreCommunity) (string, error) {
	model := ToModel(core)
	tx := storage.query.Create(&model)
	if tx.Error != nil {
		return "Masalah Pada Insert Database", tx.Error
	}
	model2 := GetLeader(userid, int(model.ID))
	tx2 := storage.query.Create(&model2)
	if tx2.Error != nil {
		return "Masalah saat bergabung ke community", tx.Error
	}

	return "Suksus Bergabung ke Community", nil
}

func (storage *Storage) SelectList() ([]community.CoreCommunity, string, error) {
	var models []Community
	tx := storage.query.Find(&models)
	if tx.Error != nil {
		return nil, "Terjadi Kesalahan", tx.Error
	}

	listcore := ToCoreList(models)

	for k, v := range listcore {
		Count := storage.query.Model(&JoinCommunity{}).Where("community_id = ?", v.ID).Count(&v.Members)
		if Count.Error != nil {
			return nil, "Terjadi Kesalahan Saat Menghitung Member", Count.Error
		}
		listcore[k].Members = v.Members
	}

	return listcore, "Sukses Mendapatkan Semua Data", nil
}

func (storage *Storage) SelectMembers(communityid int) ([]string, string, error) {
	var models []JoinCommunity
	tx := storage.query.Find(&models, "community_id = ?", communityid)
	if tx.Error != nil {
		return nil, "Terjadi Kesalahan pada pengambilan member", tx.Error
	}

	var names []string
	for _, v := range models {
		var temp user.User
		txx := storage.query.Find(&temp, "id = ?", v.UserID)
		if txx.Error != nil {
			return nil, "Terjadi Kesalahan pada pengambilan nama", txx.Error
		}
		names = append(names, temp.Name)
	}

	return names, "Sekses Mendapatkan Semua Nama", nil
}

func (storage *Storage) Delete(userid, communityid int) (string, error) {
	tx := storage.query.Unscoped().Delete(&JoinCommunity{}, "user_id = ? and community_id = ?", userid, communityid)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return "Terjadi Kesalahan", tx.Error
	}
	return "Sukses Keluar dari Community", nil
}

func (storage *Storage) GetUserRole(userid, communityid int) (string, error) {
	var model JoinCommunity
	tx := storage.query.Find(&model, "user_id = ? and community_id = ?", userid, communityid)
	if tx.Error != nil {
		return "Gagal mendapatkan role User", tx.Error
	}

	return model.Role, nil
}

func (storage *Storage) UpdateCommunity(communityid int, core community.CoreCommunity) (string, error) {
	update := ToModel(core)
	tx := storage.query.Model(&update).Where("id = ?", communityid).Updates(update)
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Gagal Update", tx.Error
	}

	return "Sukses Update", nil

}

func (storage *Storage) CheckJoin(userid, communityid int) (string, error) {

	tx := storage.query.Find(&JoinCommunity{}, "user_id = ? and community_id = ?", userid, communityid)
	if tx.Error != nil {
		return "Gagal Chek Join", tx.Error
	} else if tx.RowsAffected == 1 {
		return "Anda Sudah ada di community", errors.New("Anda telah Bergabung")
	}

	return "Bisa Join", nil
}

func (storage *Storage) InsertToJoin(userid, communityid int) (string, error) {
	model := ToJoin(userid, communityid)
	tx := storage.query.Create(&model)
	if tx.Error != nil {
		return "Gagal Bergabung ke Community", tx.Error
	}

	return "Sukses Begabung Ke Community", nil
}

func (storage *Storage) SelectCommunity(communityid int) (community.CoreCommunity, string, error) {
	var get Community
	tx := storage.query.Preload("Feeds").Find(&get, "id = ?", communityid)
	if tx.Error != nil {
		return community.CoreCommunity{}, "Gagal Ke database", tx.Error
	}
	fmt.Println(get)

	return community.CoreCommunity{}, "Sukses mendapat Data", nil
}
