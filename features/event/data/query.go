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

type temp struct {
	ID          uint
	Logo        string
	Title       string
	Description string
	Count       int64
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

	var dataEvent []event.Response
	if search == "" {
		tx := repo.db.Model(&Community{}).Select("events.id as id, communities.logo as logo, events.title as title, count(join_events.id) as members, events.description as descriptions, events.date as date, events.price as price").Joins("inner join events on events.community_id = communities.id").Joins("inner join join_events on join_events.event_id = events.id").Group("events.id").Scan(&dataEvent)
		if tx.Error != nil {
			return nil, tx.Error
		}

		return dataEvent, nil
	} else {
		tx := repo.db.Model(&Community{}).Select("events.id as id, communities.logo as logo, events.title as title, count(join_events.id) as members, events.description as descriptions, events.date as date, events.price as price").Joins("inner join events on events.community_id = communities.id").Joins("inner join join_events on join_events.event_id = events.id").Where("events.title like ?", ("%" + search + "%")).Group("events.id").Scan(&dataEvent)
		if tx.Error != nil {
			return nil, tx.Error
		}

		return dataEvent, nil
	}

}

func (repo *eventData) SelectEventComu(search string, idComu, userId int) (event.CommunityEvent, error) {

	var dataEvent []event.Response
	if search == "" {
		tx := repo.db.Model(&Community{}).Select("events.id as id, communities.logo as logo, events.title as title, count(join_events.id) as members, events.description as descriptions, events.date as date, events.price as price").Joins("inner join events on events.community_id = communities.id").Joins("inner join join_events on join_events.event_id = events.id").Group("events.id").Scan(&dataEvent)
		if tx.Error != nil {
			return event.CommunityEvent{}, tx.Error
		}
	} else {
		tx := repo.db.Model(&Community{}).Select("events.id as id, communities.logo as logo, events.title as title, count(join_events.id) as members, events.description as descriptions, events.date as date, events.price as price").Joins("inner join events on events.community_id = communities.id").Joins("inner join join_events on join_events.event_id = events.id").Where("events.title like ?", ("%" + search + "%")).Group("events.id").Scan(&dataEvent)
		if tx.Error != nil {
			return event.CommunityEvent{}, tx.Error
		}
	}

	var EventComu temp
	tx := repo.db.Model(&Community{}).Select("communities.id as id, communities.logo as logo, communities.title as title, communities.descriptions as description, count(join_communities.id) as count").Joins("inner join join_communities on join_communities.community_id = communities.id").Where("communities.id = ? ", idComu).Group("communities.id").Scan(&EventComu)
	if tx.Error != nil {
		return event.CommunityEvent{}, tx.Error
	}
	var role JoinCommunity
	repo.db.First(&role, "user_id = ?", userId)

	var dataReturn = event.CommunityEvent{
		ID:          EventComu.ID,
		Role:        role.Role,
		Logo:        EventComu.Logo,
		Title:       EventComu.Title,
		Description: EventComu.Description,
		Count:       EventComu.Count,
		Event:       dataEvent,
	}

	return dataReturn, nil

}
