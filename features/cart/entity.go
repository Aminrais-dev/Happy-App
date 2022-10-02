package cart

type CoreCart struct {
	ID           uint
	UserID       uint
	ProductID    uint
	Name         string
	Descriptions string
	Photo        string
	Price        int
}

type CoreCommunity struct {
	ID           uint
	Title        string
	Descriptions string
	Logo         string
	Members      int64
}

type DataInterface interface {
	InsertIntoCart(userid, productid int) (string, error)
	GetCommunity(communityid int) (CoreCommunity, string, error)
	SelectCartList(userid, communityid int) ([]CoreCart, string, error)
	DeleteFromCart(cartid int) (string, error)
}

type UsecaseInterface interface {
	AddToCart(userid, productid int) (string, error)
	GetCartList(userid, communityid int) (CoreCommunity, []CoreCart, string, error)
	DeleteFromCart(cartid int) (string, error)
}
