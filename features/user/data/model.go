package data

import (
	"capstone/happyApp/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Username  string
	Gender    string
	Email     string
	Password  string
	Photo     string
	Status    string
	Community []JoinCommunity
	Feeds     []Feed
	Comments  []Comment
	Event     []JoinEvent
}

type Community struct {
	gorm.Model
	Title        string
	Descriptions string
	Logo         string
	Member       []JoinCommunity
}

type JoinCommunity struct {
	gorm.Model
	UserID      uint
	CommunityID uint
	Role        string
}

type JoinEvent struct {
	gorm.Model
	UserID uint
}

type myCommunity struct {
	ID    uint
	Title string
	Logo  string
	Role  string
}

type Feed struct {
	gorm.Model
	Text        string
	UserID      uint
	CommunityID uint
	Comments    []Comment
}

type Comment struct {
	gorm.Model
	Text   string
	FeedID uint
	UserID uint
}

func (data *User) toCore() user.CoreUser {
	return user.CoreUser{
		ID:        data.ID,
		Name:      data.Name,
		Username:  data.Username,
		Gender:    data.Gender,
		Email:     data.Email,
		Password:  data.Password,
		Photo:     data.Photo,
		Community: toCoreJoinList(data.Community),
	}
}

func (data *JoinCommunity) toCoreJoin() user.JoinCommunityCore {
	return user.JoinCommunityCore{
		ID:          data.ID,
		UserID:      data.UserID,
		CommunityID: data.CommunityID,
		Role:        data.Role,
		Created_at:  data.CreatedAt,
	}
}

func toCoreJoinList(data []JoinCommunity) []user.JoinCommunityCore {

	var resData []user.JoinCommunityCore
	for key := range data {
		resData = append(resData, data[key].toCoreJoin())
	}

	return resData
}

func fromCore(data user.CoreUser) User {
	return User{
		Name:     data.Name,
		Username: data.Username,
		Gender:   data.Gender,
		Email:    data.Email,
		Password: data.Password,
		Photo:    data.Photo,
		Status:   data.Status,
	}
}

func (data *myCommunity) toComuCore() user.CommunityProfile {
	return user.CommunityProfile{
		ID:    data.ID,
		Title: data.Title,
		Logo:  data.Logo,
		Role:  data.Role,
	}
}

func toList(data []myCommunity) []user.CommunityProfile {

	var dataCore []user.CommunityProfile
	for key := range data {
		dataCore = append(dataCore, data[key].toComuCore())
	}
	return dataCore

}
