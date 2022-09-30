package usecase

import (
	"capstone/happyApp/features/product"
)

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

func (usecase *productUsecase) DeleteProduct(idProduct, userId int) int {

	row := usecase.productData.DelProduct(idProduct, userId)
	return row

}

func (usecase *productUsecase) UpdateProduct(data product.ProductCore, userId int) int {

	row := usecase.productData.UpdtProduct(data, userId)
	return row

}

func (usecase *productUsecase) GetProduct(idProduct, userId int) (product.Comu, product.ProductCore, error) {

	dataComu, dataProduct, err := usecase.productData.SelectProduct(idProduct, userId)
	if err != nil {
		return dataComu, dataProduct, err
	}

	return dataComu, dataProduct, nil

}
