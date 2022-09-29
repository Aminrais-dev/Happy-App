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

type DataInterface interface {
	InsertEvent(EventCore, int) int
}

type UsecaseInterface interface {
	PostEvent(EventCore, int) int
}
