package usecase

import "capstone/happyApp/features/product"

type productUsecase struct {
	productData product.DataInterface
}

func New(data product.DataInterface) product.UsecaseInterface {
	return &productUsecase{
		productData: data,
	}
}

func (usecase *productUsecase) PostProduct(data product.ProductCore, userId int) int {

	if data.Description == "" || data.Name == "" || data.Stock == 0 {
		return -3
	}

	row := usecase.productData.InsertProduct(data, userId)
	return row

}
