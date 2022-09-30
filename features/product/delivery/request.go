package delivery

import "capstone/happyApp/features/product"

type Request struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"descriptions" form:"descriptions"`
	Photo       string `json:"photo" form:"photo"`
	Stock       uint64 `json:"stock" form:"stock"`
	Price       uint64 `json:"price" form:"price"`
}

func (req *Request) resToCore(id int) product.ProductCore {

	return product.ProductCore{
		Name:        req.Name,
		Description: req.Description,
		Photo:       req.Photo,
		Stock:       req.Stock,
		Price:       req.Price,
		CommunityID: uint(id),
	}
}
