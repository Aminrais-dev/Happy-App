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
	Members      int64     `json:"members"`
	Descriptions string    `json:"descriptions"`
	Date         time.Time `json:"date"`
	Price        int64     `json:"price"`
}

type DataInterface interface {
	InsertEvent(EventCore, int) int
	SelectEvent(string) ([]Response, error)
}

type UsecaseInterface interface {
	PostEvent(EventCore, int) int
	GetEvent(string) ([]Response, error)
}
