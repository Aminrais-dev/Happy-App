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
	DelProduct(idProduct, userId int) int
	UpdtProduct(ProductCore, int) int
}

type UsecaseInterface interface {
	PostProduct(ProductCore, int) int
	DeleteProduct(idProduct, userId int) int
	UpdateProduct(ProductCore, int) int
}
