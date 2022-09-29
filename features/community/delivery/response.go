package delivery

import "capstone/happyApp/features/community"

type Respose struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	Logo         string `json:"logo"`
	Members      int64  `json:"members"`
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
