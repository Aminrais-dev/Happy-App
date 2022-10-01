package delivery

import "capstone/happyApp/features/community"

type Request struct {
	Title        string `json:"title" form:"title"`
	Descriptions string `json:"descriptions" form:"descriptions"`
	Logo         string `json:"logo" form:"logo"`
}

type FeedRequst struct {
	Text string `json:"text" form:"text"`
}

type CommentRequst struct {
	Text string `json:"text" form:"text"`
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

func (feed *FeedRequst) ToCore(userid, communityid int) community.CoreFeed {
	return community.CoreFeed{
		UserID:      uint(userid),
		CommunityID: uint(communityid),
		Text:        feed.Text,
	}
}

func (coment *CommentRequst) ToCore(userid, feedid int) community.CoreComment {
	return community.CoreComment{
		UserID: uint(userid),
		FeedID: uint(feedid),
		Text:   coment.Text,
	}
}
