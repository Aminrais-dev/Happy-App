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

type Comu struct {
	ID          uint
	Role        string
	Logo        string
	Title       string
	Description string
	Count       int64
}

type DataInterface interface {
	InsertProduct(ProductCore, int) int
	DelProduct(idProduct, userId int) int
	UpdtProduct(ProductCore, int) int
	SelectProduct(idProduct, userId int) (Comu, ProductCore, error)
}

type UsecaseInterface interface {
	PostProduct(ProductCore, int) int
	DeleteProduct(idProduct, userId int) int
	UpdateProduct(ProductCore, int) int
	GetProduct(idProduct, userId int) (Comu, ProductCore, error)
}
