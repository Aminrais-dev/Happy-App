package delivery

import (
	"capstone/happyApp/features/event"
	"time"
)

type Request struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"descriptions" form:"descriptions"`
	Price       uint64 `json:"price" form:"price"`
	Date        string `json:"date_event" form:"date_event"`
	Location    string `json:"location" form:"location"`
	CommunityID uint
}

func (data *Request) resToCore() event.EventCore {

	var layout = "2006-01-02 15:04:05"
	date, _ := time.Parse(layout, data.Date)

	return event.EventCore{
		Title:       data.Title,
		Description: data.Description,
		Price:       data.Price,
		Date:        date,
		Location:    data.Location,
		CommunityID: data.CommunityID,
	}
}
