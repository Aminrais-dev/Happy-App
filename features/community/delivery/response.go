package delivery

import (
	"capstone/happyApp/features/community"
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
	Date     time.Time         `json:"date"`
	Comments []ResponseComment `json:"comments"`
}

type ResponseComment struct {
	Name string    `json:"name"`
	Text string    `json:"text"`
	Date time.Time `json:"date"`
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
