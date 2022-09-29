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
	Community []JoinCommunity
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

func (data *User) toCore() user.CoreUser {
	return user.CoreUser{
		ID:        data.ID,
		Name:      data.Name,
		Username:  data.Username,
		Gender:    data.Gender,
		Email:     data.Email,
		Password:  data.Password,
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
	}
}
