package data

import (
	"capstone/happyApp/features/event"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
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
	UserID           uint
	EventID          uint
	Order_id         string
	Type_payment     string
	Payment_method   string
	Status_payment   string
	Midtrans_virtual string
	GrossAmount      string
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

func EventList(data []tempRespon) []event.Response {

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

func EventListComu(data []event.Response, dataComu temp, role string) event.CommunityEvent {
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

func EventDetails(data tempDetail, role string) event.EventDetail {
	return event.EventDetail{
		ID:            data.ID,
		Title:         data.Title,
		Description:   data.Description,
		Status:        role,
		Penyelenggara: data.Penyelenggara,
		Date:          data.Date,
		Price:         data.Price,
		Location:      data.Location,
	}
}

func toModelJoinEvent(data *coreapi.ChargeResponse, userId, idEvent int, method string) JoinEvent {

	var midtransVirtual string
	if data.VaNumbers != nil {
		midtransVirtual = data.VaNumbers[0].VANumber
	} else if data.BillKey != "" {
		midtransVirtual = data.BillKey
	} else if data.Actions != nil {
		midtransVirtual = data.Actions[0].URL
	}

	return JoinEvent{
		UserID:           uint(userId),
		EventID:          uint(idEvent),
		Type_payment:     data.PaymentType,
		Payment_method:   method,
		Order_id:         data.OrderID,
		Status_payment:   data.TransactionStatus,
		Midtrans_virtual: midtransVirtual,
		GrossAmount:      data.GrossAmount,
	}
}
