package event

import "time"

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
	ID           uint      `json:"id"`
	Logo         string    `json:"logo"`
	Title        string    `json:"title"`
	Members      uint8     `json:"members"`
	Descriptions string    `json:"descriptions"`
	Date         time.Time `json:"date"`
	Price        int64     `json:"price"`
}

type CommunityEvent struct {
	ID          uint       `json:"id"`
	Role        string     `json:"role"`
	Logo        string     `json:"logo"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Count       int64      `json:"members"`
	Event       []Response `json:"event"`
}

type EventDetail struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Status        string    `json:"status"`
	Description   string    `json:"descriptions"`
	Penyelenggara string    `json:"penyelenggara"`
	Date          time.Time `json:"date_event"`
	Partisipasi   uint8     `json:"partisipasi"`
	Price         uint64    `json:"price"`
	Location      string    `json:"location"`
}

type DataInterface interface {
	InsertEvent(EventCore, int) int
	SelectEvent(string) ([]Response, error)
	SelectEventComu(search string, idComu, userId int) (CommunityEvent, error)
	SelectEventDetail(idEvent, userId int) (EventDetail, error)
}

type UsecaseInterface interface {
	PostEvent(EventCore, int) int
	GetEvent(string) ([]Response, error)
	GetEventComu(search string, idComu, userId int) (CommunityEvent, error)
	GetEventDetail(idEvent, userId int) (EventDetail, error)
}
