package data

import (
	"capstone/happyApp/features/event"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string
	Description string
	Price       uint64
	Date        time.Time
	Location    string
	CommunityID uint
	Community   Community
}

type JoinEvent struct {
	gorm.Model
	UserID          uint
	EventID         uint
	Type_payment    string
	Status_payment  string
	Virtual_account string
}

type User struct {
	gorm.Model
}

type Community struct {
	gorm.Model
	Logo  string
	Title string
}

type JoinCommunity struct {
	gorm.Model
	UserID      uint
	CommunityID uint
	Role        string
}

type temp struct {
	ID          uint
	Logo        string
	Title       string
	Description string
	Count       int64
}

type tempDetail struct {
	ID            uint
	Title         string
	Description   string
	Penyelenggara string
	Date          time.Time
	Price         uint64
	Location      string
}

type tempRespon struct {
	ID           uint
	Logo         string
	Title        string
	Descriptions string
	Date         time.Time
	Price        int64
}

type member struct {
	Member uint8
}

func fromCore(data event.EventCore) Event {
	return Event{
		Title:       data.Title,
		Description: data.Description,
		Price:       data.Price,
		Date:        data.Date,
		Location:    data.Location,
		CommunityID: data.CommunityID,
	}
}

func toRes(data []tempRespon) []event.Response {

	var dataRespon []event.Response
	for _, v := range data {
		dataRespon = append(dataRespon, event.Response{
			ID:           v.ID,
			Logo:         v.Logo,
			Title:        v.Title,
			Descriptions: v.Descriptions,
			Date:         v.Date,
			Price:        v.Price,
		})
	}

	return dataRespon
}

func resEventComu(data []event.Response, dataComu temp, role string) event.CommunityEvent {
	return event.CommunityEvent{
		ID:          dataComu.ID,
		Role:        role,
		Logo:        dataComu.Logo,
		Title:       dataComu.Title,
		Description: dataComu.Description,
		Count:       dataComu.Count,
		Event:       data,
	}
}

func resEventDetail(data tempDetail, member member, role string) event.EventDetail {
	return event.EventDetail{
		ID:            data.ID,
		Title:         data.Title,
		Description:   data.Description,
		Status:        role,
		Penyelenggara: data.Penyelenggara,
		Date:          data.Date,
		Partisipasi:   member.Member,
		Price:         data.Price,
		Location:      data.Location,
	}
}
