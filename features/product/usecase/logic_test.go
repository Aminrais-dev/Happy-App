package usecase

import (
	"capstone/happyApp/features/product"
	"capstone/happyApp/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProductComu(t *testing.T) {

	productMock := new(mocks.DataProduct)
	var returnData2 []product.ProductCore
	returnData2 = append(returnData2, product.ProductCore{ID: 1, Name: "baju baru", Description: "baju kain woll terbaik", Photo: "https://photo", Stock: 10, Price: 30000, CommunityID: 1})
	returnData := product.Comu{ID: 1, Title: "comunity keren", Role: "member", Logo: "https://logo", Description: "community untuk jadi keren", Count: 12}

	comuId := uint(1)
	userId := 1

	t.Run("Get product community success", func(t *testing.T) {
		productMock.On("SelectProductComu", int(comuId), userId).Return(returnData, returnData2, nil).Once()

		useCase := New(productMock)
		res, res2, _ := useCase.GetProductComu(int(comuId), userId)
		assert.Equal(t, res.ID, comuId)
		assert.Equal(t, res2[0].CommunityID, comuId)
		productMock.AssertExpectations(t)

	})

	t.Run("Get product community failed", func(t *testing.T) {

		var kosong []product.ProductCore
		kosong = append(kosong, product.ProductCore{})
		productMock.On("SelectProductComu", int(comuId), userId).Return(product.Comu{}, kosong, errors.New("error")).Once()

		useCase := New(productMock)
		res, res2, err := useCase.GetProductComu(int(comuId), userId)
		assert.Error(t, err)
		assert.NotEqual(t, res.ID, comuId)
		assert.NotEqual(t, res2[0].CommunityID, comuId)
		productMock.AssertExpectations(t)

	})

}

func TestGetProductId(t *testing.T) {

	productMock := new(mocks.DataProduct)
	returnData := product.Comu{ID: 1, Title: "comunity keren", Role: "member", Logo: "https://logo", Description: "community untuk jadi keren", Count: 12}
	returnData2 := product.ProductCore{ID: 1, Name: "baju baru", Description: "baju kain woll terbaik", Photo: "https://photo", Stock: 10, Price: 30000, CommunityID: 1}

	productId := uint(1)
	userId := 1

	t.Run("Get product success", func(t *testing.T) {
		productMock.On("SelectProduct", int(productId), userId).Return(returnData, returnData2, nil).Once()

		useCase := New(productMock)
		res, res2, _ := useCase.GetProduct(int(productId), userId)
		assert.Equal(t, res.ID, productId)
		assert.Equal(t, res2.CommunityID, productId)
		productMock.AssertExpectations(t)

	})

	t.Run("Get product failed", func(t *testing.T) {

		productMock.On("SelectProduct", int(productId), userId).Return(product.Comu{}, product.ProductCore{}, errors.New("error")).Once()

		useCase := New(productMock)
		res, res2, err := useCase.GetProduct(int(productId), userId)
		assert.Error(t, err)
		assert.NotEqual(t, res.ID, productId)
		assert.NotEqual(t, res2.CommunityID, productId)
		productMock.AssertExpectations(t)

	})

}

func TestUpdateProduct(t *testing.T) {

	productMock := new(mocks.DataProduct)
	input := product.ProductCore{Name: "baju baru", Description: "baju kain woll terbaik", Photo: "https://photo", Stock: 10, Price: 30000}
	userId := 1

	t.Run("update succes", func(t *testing.T) {

		productMock.On("UpdtProduct", input, userId).Return(1, nil).Once()

		useCase := New(productMock)
		res := useCase.UpdateProduct(input, userId)
		assert.Equal(t, 1, res)
		productMock.AssertExpectations(t)

	})

	t.Run("update failed", func(t *testing.T) {

		productMock.On("UpdtProduct", input, userId).Return(-1).Once()

		useCase := New(productMock)
		res := useCase.UpdateProduct(input, userId)
		assert.Equal(t, -1, res)
		productMock.AssertExpectations(t)

	})

	t.Run("failed not have access", func(t *testing.T) {

		productMock.On("UpdtProduct", input, userId).Return(-2).Once()

		useCase := New(productMock)
		res := useCase.UpdateProduct(input, userId)
		assert.Equal(t, -2, res)
		productMock.AssertExpectations(t)

	})

}

func TestPostProduct(t *testing.T) {

	productMock := new(mocks.DataProduct)
	input := product.ProductCore{Name: "baju baru", Description: "baju kain woll terbaik", Photo: "https://photo", Stock: 10, Price: 30000}
	userId := 1

	t.Run("create success", func(t *testing.T) {

		productMock.On("InsertProduct", mock.Anything, mock.Anything).Return(1).Once()

		useCase := New(productMock)
		res := useCase.PostProduct(input, userId)
		assert.Equal(t, 1, res)
		productMock.AssertExpectations(t)
	})

	t.Run("create failed", func(t *testing.T) {

		productMock.On("InsertProduct", mock.Anything, mock.Anything).Return(-1).Once()

		useCase := New(productMock)
		res := useCase.PostProduct(input, userId)
		assert.Equal(t, -1, res)
		productMock.AssertExpectations(t)

	})

	t.Run("failed not have access", func(t *testing.T) {

		productMock.On("InsertProduct", mock.Anything, mock.Anything).Return(-2).Once()

		useCase := New(productMock)
		res := useCase.PostProduct(input, userId)
		assert.Equal(t, -2, res)
		productMock.AssertExpectations(t)

	})

	t.Run("create failed because input name not filled", func(t *testing.T) {

		input.Name = ""
		useCase := New(productMock)
		res := useCase.PostProduct(input, userId)
		assert.Equal(t, -3, res)
		productMock.AssertExpectations(t)

	})

}

func TestDelete(t *testing.T) {

	productMock := new(mocks.DataProduct)
	userId := 1
	productId := 1

	t.Run("delete succes", func(t *testing.T) {

		productMock.On("DelProduct", productId, userId).Return(1).Once()

		useCase := New(productMock)
		res := useCase.DeleteProduct(productId, userId)
		assert.Equal(t, 1, res)
		productMock.AssertExpectations(t)

	})

	t.Run("delete failed", func(t *testing.T) {

		productMock.On("DelProduct", productId, userId).Return(-1).Once()

		useCase := New(productMock)
		res := useCase.DeleteProduct(productId, userId)
		assert.Equal(t, -1, res)
		productMock.AssertExpectations(t)

	})

	t.Run("failed not have access", func(t *testing.T) {

		productMock.On("DelProduct", productId, userId).Return(-2).Once()

		useCase := New(productMock)
		res := useCase.DeleteProduct(productId, userId)
		assert.Equal(t, -2, res)
		productMock.AssertExpectations(t)

	})

}
