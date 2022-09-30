package delivery

import "capstone/happyApp/features/product"

type Response struct {
	ID          uint        `json:"id"`
	Role        string      `json:"role"`
	Logo        string      `json:"logo"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Count       int64       `json:"members"`
	Product     interface{} `json:"product"`
}

type dataProductRespon struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"descriptions,omitempty"`
	Photo       string `json:"photo"`
	Price       uint64 `json:"price"`
	Stock       uint64 `json:"stock,omitempty"`
}

func ResponseDetail(dataComu product.Comu, dataProduct product.ProductCore) Response {

	return Response{
		ID:          dataComu.ID,
		Role:        dataComu.Role,
		Logo:        dataComu.Logo,
		Title:       dataComu.Title,
		Description: dataComu.Description,
		Count:       dataComu.Count,
		Product:     toRes(dataProduct),
	}
}

func ResponseDetailList(dataComu product.Comu, dataProduct []product.ProductCore) Response {

	return Response{
		ID:          dataComu.ID,
		Role:        dataComu.Role,
		Logo:        dataComu.Logo,
		Title:       dataComu.Title,
		Description: dataComu.Description,
		Count:       dataComu.Count,
		Product:     toResList(dataProduct),
	}
}

func toRes(dataProduct product.ProductCore) dataProductRespon {
	return dataProductRespon{
		ID:          dataProduct.ID,
		Name:        dataProduct.Name,
		Description: dataProduct.Description,
		Photo:       dataProduct.Photo,
		Price:       dataProduct.Price,
		Stock:       dataProduct.Stock,
	}
}

func toResList(dataProduct []product.ProductCore) []dataProductRespon {

	var dataRespon []dataProductRespon
	for _, v := range dataProduct {
		dataRespon = append(dataRespon, dataProductRespon{
			ID:    v.ID,
			Name:  v.Name,
			Price: v.Price,
			Photo: v.Photo,
		})
	}

	return dataRespon
}
