package data

import (
	"capstone/happyApp/features/event"

	"gorm.io/gorm"
)

type eventData struct {
	db *gorm.DB
}

func New(db *gorm.DB) event.DataInterface {
	return &eventData{
		db: db,
	}
}

func (repo *eventData) InsertEvent(data event.EventCore, id int) int {

	var check JoinCommunity
	tx := repo.db.First(&check, "user_id = ? AND community_id = ?", id, data.CommunityID)
	if tx.Error != nil {
		return -2
	}

	if check.Role != "admin" {
		return -2
	}

	dataCreate := fromCore(data)
	txCreate := repo.db.Create(&dataCreate)
	if txCreate.Error != nil {
		return -1
	}

	return 1

}

func (repo *eventData) SelectEvent(search string) ([]event.Response, error) {

	var dataEvent []tempRespon
	if search == "" {
		tx := repo.db.Model(&Community{}).Select("events.id as id, communities.logo as logo, events.title as title, events.description as descriptions, events.date as date, events.price as price").Joins("inner join events on events.community_id = communities.id").Group("events.id").Scan(&dataEvent)
		if tx.Error != nil {
			return nil, tx.Error
		}

		var dataRes = toRes(dataEvent)
		for key := range dataRes {
			repo.db.Model(&JoinEvent{}).Select("count(join_events.id) as member").Where("event_id = ? ", dataRes[key].ID).Scan(&dataRes[key].Members)
		}

		return dataRes, nil
	} else {
		tx := repo.db.Model(&Community{}).Select("events.id as id, communities.logo as logo, events.title as title, events.description as descriptions, events.date as date, events.price as price").Joins("inner join events on events.community_id = communities.id").Where("events.title like ?", ("%" + search + "%")).Group("events.id").Scan(&dataEvent)
		if tx.Error != nil {
			return nil, tx.Error
		}

		var dataRes = toRes(dataEvent)
		for key := range dataRes {
			repo.db.Model(&JoinEvent{}).Select("count(join_events.id) as member").Where("event_id = ? ", dataRes[key].ID).Scan(&dataRes[key].Members)
		}

		return dataRes, nil
	}

}

func (repo *eventData) SelectEventComu(search string, idComu, userId int) (event.CommunityEvent, error) {

	var dataEvent []tempRespon
	var dataRes []event.Response
	if search == "" {
		tx := repo.db.Model(&Community{}).Select("events.id as id, communities.logo as logo, events.title as title, events.description as descriptions, events.date as date, events.price as price").Joins("inner join events on events.community_id = communities.id").Where("events.community_id = ? ", idComu).Group("events.id").Scan(&dataEvent)
		if tx.Error != nil {
			return event.CommunityEvent{}, tx.Error
		}

		dataRes = toRes(dataEvent)
		for key := range dataRes {
			repo.db.Model(&JoinEvent{}).Select("count(join_events.id) as member").Where("event_id = ? ", dataRes[key].ID).Scan(&dataRes[key].Members)
		}

	} else {
		tx := repo.db.Model(&Community{}).Select("events.id as id, communities.logo as logo, events.title as title, events.description as descriptions, events.date as date, events.price as price").Joins("inner join events on events.community_id = communities.id").Where("events.title like ? AND events.community_id = ? ", ("%" + search + "%"), idComu).Group("events.id").Scan(&dataEvent)
		if tx.Error != nil {
			return event.CommunityEvent{}, tx.Error
		}

		dataRes = toRes(dataEvent)
		for key := range dataRes {
			repo.db.Model(&JoinEvent{}).Select("count(join_events.id) as member").Where("event_id = ? ", dataRes[key].ID).Scan(&dataRes[key].Members)
		}

	}

	var EventComu temp
	tx := repo.db.Model(&Community{}).Select("communities.id as id, communities.logo as logo, communities.title as title, communities.descriptions as description, count(join_communities.user_id) as count").Joins("inner join join_communities on join_communities.community_id = communities.id").Where("communities.id = ? AND join_communities.deleted_at IS NULL ", idComu).Group("communities.id").Scan(&EventComu)
	if tx.Error != nil {
		return event.CommunityEvent{}, tx.Error
	}
	var role JoinCommunity
	repo.db.First(&role, "user_id = ? AND community_id = ? ", userId, idComu)

	var dataReturn = resEventComu(dataRes, EventComu, role.Role)

	return dataReturn, nil

}

func (repo *eventData) SelectEventDetail(idEvent, userId int) (event.EventDetail, error) {

	var data tempDetail
	repo.db.Model(&Community{}).Select("events.id as id, events.title as title, events.description as description, communities.title as penyelenggara, events.date as date, events.price as price, events.location as location").Joins("inner join events on events.community_id = communities.id").Where("events.id = ?", idEvent).Group("events.id").Scan(&data)

	var member member
	repo.db.Model(&Community{}).Select("count(join_events.id) as member").Where("join_events.event_id = ? ", idEvent).Scan(&member)

	var role JoinEvent
	repo.db.First(&role, "user_id = ? AND event_id = ? ", userId, idEvent)

	var status = "join"
	if role.UserID != uint(userId) {
		status = "not join"
	}

	return resEventDetail(data, member, status), nil

}
