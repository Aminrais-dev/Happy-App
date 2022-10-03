package event

import (
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
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

type DataInterface interface {
	InsertEvent(EventCore, int) int
	SelectEvent(string) ([]Response, error)
	SelectEventComu(search string, idComu, userId int) (CommunityEvent, error)
	SelectEventDetail(idEvent, userId int) (EventDetail, error)
	SelectAmountEvent(idEvent int) uint64
	CreatePayment(reqMidtrans coreapi.ChargeReq, userId, EventId int, method string) (*coreapi.ChargeResponse, error)
}

type UsecaseInterface interface {
	PostEvent(EventCore, int) int
	GetEvent(string) ([]Response, error)
	GetEventComu(search string, idComu, userId int) (CommunityEvent, error)
	GetEventDetail(idEvent, userId int) (EventDetail, error)
	GetAmountEvent(idEvent int) uint64
	CreatePaymentMidtrans(reqMidtrans coreapi.ChargeReq, userId, EventId int, method string) (*coreapi.ChargeResponse, error)
}
