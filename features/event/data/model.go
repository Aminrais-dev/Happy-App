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
