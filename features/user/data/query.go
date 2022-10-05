package data

import (
	"capstone/happyApp/features/user"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}
}

func (repo *userData) InsertUser(data user.CoreUser) int {

	newData := fromCore(data)
	tx := repo.db.Create(&newData)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return -1
	}

	return 1

}

func (repo *userData) DelUser(id int) int {

	tx := repo.db.Where("id = ? ", id).Delete(&User{})
	if tx.Error != nil {
		return -1
	}

	return 1

}

func (repo *userData) UpdtUser(data user.CoreUser) int {

	newData := fromCore(data)

	tx := repo.db.Model(&User{}).Where("id = ? ", int(data.ID)).Updates(newData)
	if tx.Error != nil {
		return -1
	}

	return int(tx.RowsAffected)

}

func (repo *userData) SelectUser(id int) (user.CoreUser, []user.CommunityProfile, error) {

	var data User
	tx := repo.db.First(&data, "id = ? ", id)
	if tx.Error != nil {
		return user.CoreUser{}, nil, tx.Error
	}

	var comu []myCommunity
	txComu := repo.db.Model(&Community{}).Select("communities.id as id, communities.title as title, communities.logo as logo, join_communities.role as role").Joins("inner join join_communities on join_communities.community_id = communities.id").Where("join_communities.user_id = ? AND join_communities.deleted_at IS NULL", id).Scan(&comu)
	if txComu.Error != nil {
		return user.CoreUser{}, nil, txComu.Error
	}

	return data.toCore(), toList(comu), nil
}
