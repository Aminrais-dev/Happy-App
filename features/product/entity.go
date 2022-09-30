package product

type ProductCore struct {
	ID          uint
	Name        string
	Description string
	Photo       string
	Stock       uint64
	Price       uint64
	CommunityID uint
}

type DataInterface interface {
	InsertProduct(ProductCore, int) int
}

type UsecaseInterface interface {
	PostProduct(ProductCore, int) int
}
