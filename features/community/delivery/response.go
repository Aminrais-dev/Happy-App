package delivery

import (
	"capstone/happyApp/features/community"
	"fmt"
	"time"
)

type Respose struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	Logo         string `json:"logo"`
	Members      int64  `json:"members"`
}

type ResponseFeed struct {
	ID       uint              `json:"id"`
	Name     string            `json:"name"`
	Text     string            `json:"text"`
	Date     string            `json:"date"`
	Comments []ResponseComment `json:"comments"`
}

type ResponseFeedNoComment struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
	Date string `json:"date"`
}

type ResponseComment struct {
	Name string `json:"name"`
	Text string `json:"text"`
	Date string `json:"date"`
}

type DetailCommunity struct {
	ID           uint   `json:"id"`
	Role         string `json:"role"`
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	Logo         string `json:"logo"`
	Members      int64  `json:"members"`
	Feeds        []ResponseFeed
}
type DetailCommunityNoComment struct {
	ID           uint   `json:"id"`
	Role         string `json:"role"`
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	Logo         string `json:"logo"`
	Members      int64  `json:"members"`
	Feeds        []ResponseFeedNoComment
}

func ToResponse(core community.CoreCommunity) Respose {
	return Respose{
		ID:           core.ID,
		Logo:         core.Logo,
		Title:        core.Title,
		Members:      core.Members,
		Descriptions: core.Descriptions,
	}
}

func ToResponseList(core []community.CoreCommunity) []Respose {
	var list []Respose
	for _, v := range core {
		list = append(list, ToResponse(v))
	}

	return list
}

func ToFeedResponse(data community.CoreFeed) ResponseFeed {
	return ResponseFeed{
		ID:   data.ID,
		Name: data.Name,
		Text: data.Text,
		Date: GetDateHour(data.Date),
	}
}

func ToFeedResponseList(data []community.CoreFeed) []ResponseFeed {
	var list []ResponseFeed
	for _, v := range data {
		list = append(list, ToFeedResponse(v))
	}

	return list
}

func ResponseWithFeed(core community.CoreCommunity) DetailCommunity {
	return DetailCommunity{
		ID:           core.ID,
		Title:        core.Title,
		Role:         core.Role,
		Descriptions: core.Descriptions,
		Logo:         core.Logo,
		Members:      core.Members,
		Feeds:        ToFeedResponseList(core.Feeds),
	}
}

func GetDateHour(data time.Time) string {
	time := fmt.Sprintf("%v", data)
	tanggal := time[:10]
	jam := time[11:19]
	return tanggal + " " + jam
}

func ToResponseComment(data community.CoreComment) ResponseComment {
	return ResponseComment{
		Name: data.Name,
		Text: data.Text,
		Date: GetDateHour(data.Date),
	}
}

func ToResponseCommentList(data []community.CoreComment) []ResponseComment {
	var list []ResponseComment
	for _, v := range data {
		list = append(list, ToResponseComment(v))
	}

	return list
}

func ResponseFeedWithComment(core community.CoreFeed) ResponseFeed {
	return ResponseFeed{
		ID:       core.ID,
		Name:     core.Name,
		Text:     core.Text,
		Date:     GetDateHour(core.Date),
		Comments: ToResponseCommentList(core.Comments),
	}
}

func ToFeedResponseNoComment(data community.CoreFeed) ResponseFeedNoComment {
	return ResponseFeedNoComment{
		ID:   data.ID,
		Name: data.Name,
		Text: data.Text,
		Date: GetDateHour(data.Date),
	}
}

func ToFeedResponseListNoComment(data []community.CoreFeed) []ResponseFeedNoComment {
	var list []ResponseFeedNoComment
	for _, v := range data {
		list = append(list, ToFeedResponseNoComment(v))
	}

	return list
}

func ResponseWithFeedNoComment(core community.CoreCommunity) DetailCommunityNoComment {
	return DetailCommunityNoComment{
		ID:           core.ID,
		Title:        core.Title,
		Role:         core.Role,
		Descriptions: core.Descriptions,
		Logo:         core.Logo,
		Members:      core.Members,
		Feeds:        ToFeedResponseListNoComment(core.Feeds),
	}
}
