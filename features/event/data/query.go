package data

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/event"
	"errors"
	"fmt"

	"github.com/midtrans/midtrans-go/coreapi"
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

		var dataRes = EventList(dataEvent)

		return dataRes, nil
	} else {
		tx := repo.db.Model(&Community{}).Select("events.id as id, communities.logo as logo, events.title as title, events.description as descriptions, events.date as date, events.price as price").Joins("inner join events on events.community_id = communities.id").Where("events.title like ?", ("%" + search + "%")).Group("events.id").Scan(&dataEvent)
		if tx.Error != nil {
			return nil, tx.Error
		}

		var dataRes = EventList(dataEvent)

		return dataRes, nil
	}

}

func (repo *eventData) SelectEventComu(idComu, userId int) (event.CommunityEvent, error) {

	var dataEvent []event.Response
	txGet := repo.db.Model(&Community{}).Select("events.id as id, communities.logo as logo, events.title as title, events.description as descriptions, events.date as date, events.price as price").Joins("inner join events on events.community_id = communities.id").Where("communities.id = ? ", idComu).Group("events.id").Scan(&dataEvent)
	if txGet.Error != nil {
		return event.CommunityEvent{}, txGet.Error
	}

	var EventComu temp
	tx := repo.db.Model(&Community{}).Select("communities.id as id, communities.logo as logo, communities.title as title, communities.descriptions as description, count(join_communities.user_id) as count").Joins("inner join join_communities on join_communities.community_id = communities.id").Where("communities.id = ? AND join_communities.deleted_at IS NULL ", idComu).Group("communities.id").Scan(&EventComu)
	if tx.Error != nil {
		return event.CommunityEvent{}, tx.Error
	}
	fmt.Println(EventComu)
	var role JoinCommunity
	repo.db.First(&role, "user_id = ? AND community_id = ? ", userId, idComu)

	var dataReturn = EventListComu(dataEvent, EventComu, role.Role)

	return dataReturn, nil

}

func (repo *eventData) SelectEventDetail(idEvent, userId int) (event.EventDetail, error) {

	var data tempDetail
	tx := repo.db.Model(&Community{}).Select("events.id as id, events.title as title, events.description as description, communities.title as penyelenggara, events.date as date, events.price as price, events.location as location").Joins("inner join events on events.community_id = communities.id").Where("events.id = ?", idEvent).Group("events.id").Scan(&data)
	if tx.Error != nil {
		return event.EventDetail{}, tx.Error
	}
	var role JoinEvent
	repo.db.First(&role, "user_id = ? AND event_id = ? ", userId, idEvent)

	var status = "join"
	if role.UserID != uint(userId) {
		status = "not join"
	}

	var dataRes = EventDetails(data, status)

	repo.db.Model(&JoinEvent{}).Select("count(join_events.id) as member").Where("event_id = ? AND status_payment = ? ", idEvent, config.SUCCESS).Scan(&dataRes.Partisipasi)

	return dataRes, nil

}

func (repo *eventData) SelectAmountEvent(idEvent int) uint64 {

	var event Event
	tx := repo.db.First(&event, "id = ? ", idEvent)
	if tx.Error != nil {
		return 00
	}

	return event.Price

}

func (repo *eventData) CreatePayment(reqMidtrans coreapi.ChargeReq, userId, idEvent int, method string) (*coreapi.ChargeResponse, error) {

	var check JoinEvent
	paid := "paid"
	repo.db.First(&check, "user_id = ? AND event_id = ? AND status_payment = ?", userId, idEvent, paid)
	if check.UserID == uint(userId) {
		return nil, errors.New("sudah join dalam event")
	}

	chargeResponse, errCreate := coreapi.ChargeTransaction(&reqMidtrans)
	if errCreate != nil {
		return nil, errCreate
	}

	data := toModelJoinEvent(chargeResponse, userId, idEvent, method)
	tx := repo.db.Create(&data)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	return chargeResponse, nil

}

func (repo *eventData) GetMembers(data []event.Response) []event.Response {

	for key := range data {
		repo.db.Model(&JoinEvent{}).Select("count(join_events.id) as member").Where("event_id = ? AND status_payment = ? ", data[key].ID, config.SUCCESS).Scan(&data[key].Members)
	}

	return data

}
