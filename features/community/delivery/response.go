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
	Name     string            `json:"name"`
	Text     string            `json:"text"`
	Date     string            `json:"date"`
	Comments []ResponseComment `json:"comments"`
}

type ResponseComment struct {
	Name string `json:"name"`
	Text string `json:"text"`
	Date string `json:"date"`
}

type DetailCommunity struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	Logo         string `json:"logo"`
	Members      int64  `json:"members"`
	Feeds        []ResponseFeed
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
