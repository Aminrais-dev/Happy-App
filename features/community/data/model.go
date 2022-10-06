package data

import (
	"capstone/happyApp/features/community"
	event "capstone/happyApp/features/event/data"

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
	Name      string
	Username  string `gorm:"unique"`
	Gender    string
	Email     string `gorm:"unique"`
	Password  string
	Photo     string
	Community []JoinCommunity
	Feeds     []Feed
	Comments  []Comment
	Event     []event.JoinEvent
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

type Product struct {
	gorm.Model
	Name        string
	Description string
	Photo       string
	Stock       uint64
	Price       uint64
	CommunityID uint
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

func ToModelFeed(data community.CoreFeed) Feed {
	return Feed{
		UserID:      data.UserID,
		CommunityID: data.CommunityID,
		Text:        data.Text,
	}
}

func ToCoreFeed(data Feed, name string) community.CoreFeed {
	return community.CoreFeed{
		ID:   data.ID,
		Name: name,
		Text: data.Text,
		Date: data.CreatedAt,
	}
}

func ToCoreWithFeed(data Community, sum int64, feeds []community.CoreFeed, role string) community.CoreCommunity {
	return community.CoreCommunity{
		ID:           data.ID,
		Role:         role,
		Title:        data.Title,
		Descriptions: data.Descriptions,
		Logo:         data.Logo,
		Members:      sum,
		Feeds:        feeds,
	}
}

func ToCoreComment(data Comment, name string) community.CoreComment {
	return community.CoreComment{
		Name: name,
		Text: data.Text,
		Date: data.CreatedAt,
	}
}

func ToCoreWithComment(data Feed, name string, comment []community.CoreComment) community.CoreFeed {
	return community.CoreFeed{
		ID:       data.ID,
		Name:     name,
		Text:     data.Text,
		Date:     data.CreatedAt,
		Comments: comment,
	}
}

func ToModelComment(data community.CoreComment) Comment {
	return Comment{
		Text:   data.Text,
		UserID: data.UserID,
		FeedID: data.FeedID,
	}
}
