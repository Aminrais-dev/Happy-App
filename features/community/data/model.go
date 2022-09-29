package data

import (
	"capstone/happyApp/features/community"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Title        string
	Descriptions string
	Logo         string
	Member       []JoinCommunity
}

type User struct {
	gorm.Model
	Community []JoinCommunity
}

type JoinCommunity struct {
	gorm.Model
	UserID      uint
	CommunityID uint
	Role        string
}

func GetLeader(userid, communityid int) JoinCommunity {
	return JoinCommunity{
		UserID:      uint(userid),
		CommunityID: uint(communityid),
		Role:        "admin",
	}
}

func ToModel(data community.CoreCommunity) Community {
	return Community{
		Title:        data.Title,
		Descriptions: data.Descriptions,
		Logo:         data.Logo,
	}
}

func ToCore(data Community) community.CoreCommunity {
	return community.CoreCommunity{
		ID:           data.ID,
		Title:        data.Title,
		Descriptions: data.Descriptions,
		Logo:         data.Logo,
	}
}

func ToCoreList(data []Community) []community.CoreCommunity {
	var list []community.CoreCommunity
	for _, v := range data {
		list = append(list, ToCore(v))
	}

	return list
}
