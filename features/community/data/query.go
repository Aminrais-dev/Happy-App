package data

import (
	"capstone/happyApp/features/community"

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

func (storage *Storage) SelectList() ([]community.CoreCommunity, string, error) {
	var models []Community
	storage.query.Find(&models)

	listcore := ToCoreList(models)

	for _, v := range listcore {
		Count := storage.query.Model(&JoinCommunity{}).Where("community_id = ?", v.ID).Count(&v.Members)
		if Count.Error != nil {
			return nil, "Terjadi Kesalahan Saat Menghitung Member", Count.Error
		}
	}

	return listcore, "Sukses Mendapatkan Semua Data", nil
}
