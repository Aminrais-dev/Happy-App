package user

import "time"

type CoreUser struct {
	ID        uint
	Name      string
	Username  string
	Gender    string
	Email     string
	Password  string
	Photo     string
	Community []JoinCommunityCore
}

type CommunityCore struct {
	ID           uint
	Title        string
	Descriptions string
	Logo         string
	Member       []JoinCommunityCore
}

type JoinCommunityCore struct {
	ID          uint
	UserID      uint
	CommunityID uint
	Role        string
	Created_at  time.Time
}

type DataInterface interface {
	InsertUser(CoreUser) int
	DelUser(int) int
	UpdtUser(CoreUser) int
}

type UsecaseInterface interface {
	PostUser(CoreUser) int
	DeleteUser(int) int
	UpdateUser(CoreUser) int
}
