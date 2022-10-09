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
	Status    string
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

type CommunityProfile struct {
	ID    uint
	Title string
	Logo  string
	Role  string
}

type DataInterface interface {
	InsertUser(CoreUser) int
	DelUser(int) int
	UpdtUser(CoreUser) int
	SelectUser(id int) (CoreUser, []CommunityProfile, error)
	CheckStatus(string, int) string
	UpdtStatus(id int, status string) int
	CheckUsername(string) int
}

type UsecaseInterface interface {
	PostUser(CoreUser) int
	DeleteUser(int) int
	UpdateUser(CoreUser) int
	GetUser(id int) (CoreUser, []CommunityProfile, error)
	UpdateStatus(id int) int
}
