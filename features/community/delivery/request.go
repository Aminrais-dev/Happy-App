package delivery

import "capstone/happyApp/features/community"

type Request struct {
	Title        string `json:"title" form:"title"`
	Descriptions string `json:"descriptions" form:"descriptions"`
	Logo         string `json:"logo" form:"logo"`
}

func (req *Request) ToCore() community.CoreCommunity {
	return community.CoreCommunity{
		Title:        req.Title,
		Descriptions: req.Descriptions,
		Logo:         req.Logo,
	}
}

func (req *Request) ToCoreWithId(communityid int) community.CoreCommunity {
	return community.CoreCommunity{
		ID:           uint(communityid),
		Title:        req.Title,
		Descriptions: req.Descriptions,
		Logo:         req.Logo,
	}
}
