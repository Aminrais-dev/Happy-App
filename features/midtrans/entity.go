package midtrans

import (
	"time"
)

type DropData struct {
	Date       time.Time
	Name       string
	Email      string
	TitleEvent string
}

type DataInterface interface {
	WeebHookUpdateTransaction(orderid, status string) (string, error)
	WeebHookUpdateJoinEvent(orderid, status string) (DropData, error)
}

type UsecaseInterface interface {
	WeebHookTransaction(orderid, status string) (string, error)
	WeebHookJoinEvent(orderid, status string) (string, error)
}
