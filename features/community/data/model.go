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
	Feeds        []Feed
}

type User struct {
	gorm.Model
	Community []JoinCommunity
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

type JoinCommunity struct {
	gorm.Model
	UserID      uint
	CommunityID uint
	Role        string
}

func ToJoin(userid, communityid int) JoinCommunity {
	return JoinCommunity{
		UserID:      uint(userid),
		CommunityID: uint(communityid),
		Role:        "member",
	}
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

// func ToCoreWithFeed(data Community) community.CoreCommunity {
// 	var feed []community.CoreFeed
// 	for _,v := range data.Feeds{
// 		feed = append(feed, )
// 	}

// 	return community.CoreCommunity{
// 		ID:           data.ID,
// 		Title:        data.Title,
// 		Descriptions: data.Descriptions,
// 		Logo:         data.Logo,
// 	}
// }
