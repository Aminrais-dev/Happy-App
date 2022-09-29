package data

import (
	"capstone/happyApp/features/community"
	user "capstone/happyApp/features/user/data"

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

	for _, v := range listcore {
		Count := storage.query.Model(&JoinCommunity{}).Where("community_id = ?", v.ID).Count(&v.Members)
		if Count.Error != nil {
			return nil, "Terjadi Kesalahan Saat Menghitung Member", Count.Error
		}
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
		txx := storage.query.Find(&temp, "id = ?", v.ID)
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
