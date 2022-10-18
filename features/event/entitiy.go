package event

import (
	"time"
)

type EventCore struct {
	ID          uint
	Title       string
	Description string
	Price       uint64
	Date        time.Time
	Location    string
	CommunityID uint
	Community   CommunityCore
}

type CommunityCore struct {
	Logo  string
	Title string
}

type Response struct {
	ID           uint
	Logo         string
	Title        string
	Members      uint8
	Descriptions string
	Date         time.Time
	Price        int64
}

type CommunityEvent struct {
	ID          uint
	Role        string
	Logo        string
	Title       string
	Description string
	Count       int64
	Event       []Response
}

type EventDetail struct {
	ID            uint
	Title         string
	Status        string
	Description   string
	Penyelenggara string
	Date          time.Time
	Partisipasi   uint8
	Price         uint64
	Location      string
}

type JoinEventCore struct {
	ID               uint
	UserID           uint
	EventID          uint
	Order_id         string
	Type_payment     string
	Payment_method   string
	Status_payment   string
	Midtrans_virtual string
	GrossAmount      string
}

type DataInterface interface {
	InsertEvent(EventCore, int) int
	SelectEvent(string) ([]Response, error)
	SelectEventComu(idComu, userId int) (CommunityEvent, error)
	SelectEventDetail(idEvent, userId int) (EventDetail, error)
	SelectAmountEvent(idEvent int) uint64
	CheckJoin(userId, EventId int) error
	InsertTransaction(JoinEventCore) error
	GetMembers([]Response) []Response
}

type UsecaseInterface interface {
	PostEvent(EventCore, int) int
	GetEvent(string) ([]Response, error)
	GetEventComu(idComu, userId int) (CommunityEvent, error)
	GetEventDetail(idEvent, userId int) (EventDetail, error)
	GetAmountEvent(idEvent int) uint64
	CheckStatus(userId, EventId int) error
	PostTransaction(JoinEventCore) error
}
