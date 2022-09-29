package community

type CoreCommunity struct {
	ID           uint
	Title        string
	Descriptions string
	Logo         string
	Members      int64
}

type DataInterface interface {
	SelectList() ([]CoreCommunity, string, error)
}

type UsecaseInterface interface {
	GetListCommunity() ([]CoreCommunity, string, error)
}
