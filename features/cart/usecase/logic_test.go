package usecase

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/cart"
	"capstone/happyApp/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddToCart(t *testing.T) {
	DataMock := new(mocks.CartData)

	t.Run("Sukses", func(t *testing.T) {
		DataMock.On("CheckCommunity", mock.Anything).Return(1, "pesan", nil).Once()
		DataMock.On("CheckMember", mock.Anything, mock.Anything).Return(1, "pesan", nil).Once()
		DataMock.On("InsertIntoCart", mock.Anything, mock.Anything).Return("pesan", nil).Once()
		logic := New(DataMock)
		msg, err := logic.AddToCart(1, 1)
		assert.Equal(t, msg, "pesan")
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed 1", func(t *testing.T) {
		DataMock.On("CheckCommunity", mock.Anything).Return(1, "pesan", errors.New("Error")).Once()
		logic := New(DataMock)
		msg, err := logic.AddToCart(1, 1)
		assert.Equal(t, msg, "pesan")
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed 2", func(t *testing.T) {
		DataMock.On("CheckCommunity", mock.Anything).Return(1, "pesan", nil).Once()
		DataMock.On("CheckMember", mock.Anything, mock.Anything).Return(1, "pesan", errors.New("error")).Once()
		logic := New(DataMock)
		msg, err := logic.AddToCart(1, 1)
		assert.Equal(t, msg, "pesan")
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed 3", func(t *testing.T) {
		DataMock.On("CheckCommunity", mock.Anything).Return(1, "pesan", nil).Once()
		DataMock.On("CheckMember", mock.Anything, mock.Anything).Return(0, "pesan", nil).Once()
		logic := New(DataMock)
		msg, err := logic.AddToCart(1, 1)
		assert.Equal(t, msg, "Anda Bukan Anggota Community Dari Pemilik Product")
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestGetCartList(t *testing.T) {
	DataMock := new(mocks.CartData)
	community := cart.CoreCommunity{ID: 1, Title: "Genshin", Members: 5, Descriptions: "Ini", Logo: config.DEFAULT_COMMUNITY}
	cart := []cart.CoreCart{{Name: "Ayaka", Descriptions: "Female", Photo: config.DEFAULT_PROFILE, Price: 10000}}

	t.Run("Success", func(t *testing.T) {
		DataMock.On("GetCommunity", mock.Anything).Return(community, "Pesan", nil).Once()
		DataMock.On("SelectCartList", mock.Anything, mock.Anything).Return(cart, "Pesan", nil).Once()
		logic := New(DataMock)
		com, car, msg, err := logic.GetCartList(1, 1)
		assert.Equal(t, com.Title, "Genshin")
		assert.Equal(t, car[0].Name, "Ayaka")
		assert.Equal(t, msg, "Pesan")
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed", func(t *testing.T) {
		DataMock.On("GetCommunity", mock.Anything).Return(community, "Pesan", errors.New("Error"))
		logic := New(DataMock)
		_, _, msg, err := logic.GetCartList(1, 1)
		assert.Equal(t, msg, "Pesan")
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestDeleteFromCart(t *testing.T) {
	DataMock := new(mocks.CartData)
	t.Run("Success", func(t *testing.T) {
		DataMock.On("DeleteFromCart", mock.Anything, mock.Anything).Return("Pesan", nil).Once()
		logic := New(DataMock)
		msg, err := logic.DeleteFromCart(1, 1)
		assert.Equal(t, "Pesan", msg)
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestInsertIntoTransaction(t *testing.T) {
	DataMock := new(mocks.CartData)

	t.Run("Success", func(t *testing.T) {
		DataMock.On("CheckStock", mock.Anything).Return([]int{1, 2, 3, 4}, "Pesan", nil).Once()
		DataMock.On("UpdateStock", mock.Anything).Return("Pesan", nil).Once()
		DataMock.On("InsertIntoTransaction", mock.Anything).Return(1, "Pesan", nil).Once()
		DataMock.On("GetTotalTransaction", mock.Anything).Return(100000, "Pesan", nil).Once()
		DataMock.On("DeleteCart", mock.Anything).Return("Pesan", nil).Once()
		logic := New(DataMock)
		id, gross, msg, err := logic.InsertIntoTransaction(cart.CoreHistory{})
		assert.Equal(t, "Pesan", msg)
		assert.Equal(t, 100000, gross)
		assert.Equal(t, 1, id)
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed 1", func(t *testing.T) {
		DataMock.On("CheckStock", mock.Anything).Return([]int{1, 2, 3, 4}, "Pesan", errors.New("Error")).Once()
		logic := New(DataMock)
		id, gross, msg, err := logic.InsertIntoTransaction(cart.CoreHistory{})
		assert.Equal(t, "Pesan", msg)
		assert.Equal(t, 0, gross)
		assert.Equal(t, 0, id)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed 2", func(t *testing.T) {
		DataMock.On("CheckStock", mock.Anything).Return([]int{1, 2, 3, 4}, "Pesan", nil).Once()
		DataMock.On("UpdateStock", mock.Anything).Return("Pesan", errors.New("Error")).Once()
		logic := New(DataMock)
		id, gross, msg, err := logic.InsertIntoTransaction(cart.CoreHistory{})
		assert.Equal(t, "Pesan", msg)
		assert.Equal(t, 0, gross)
		assert.Equal(t, 0, id)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed 3", func(t *testing.T) {
		DataMock.On("CheckStock", mock.Anything).Return([]int{1, 2, 3, 4}, "Pesan", nil).Once()
		DataMock.On("UpdateStock", mock.Anything).Return("Pesan", nil).Once()
		DataMock.On("InsertIntoTransaction", mock.Anything).Return(0, "Pesan", errors.New("Error")).Once()
		logic := New(DataMock)
		id, gross, msg, err := logic.InsertIntoTransaction(cart.CoreHistory{})
		assert.Equal(t, "Pesan", msg)
		assert.Equal(t, 0, gross)
		assert.Equal(t, 0, id)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed 4", func(t *testing.T) {
		DataMock.On("CheckStock", mock.Anything).Return([]int{1, 2, 3, 4}, "Pesan", nil).Once()
		DataMock.On("UpdateStock", mock.Anything).Return("Pesan", nil).Once()
		DataMock.On("InsertIntoTransaction", mock.Anything).Return(0, "Pesan", nil).Once()
		DataMock.On("GetTotalTransaction", mock.Anything).Return(0, "Pesan", errors.New("Error")).Once()
		logic := New(DataMock)
		id, gross, msg, err := logic.InsertIntoTransaction(cart.CoreHistory{})
		assert.Equal(t, "Pesan", msg)
		assert.Equal(t, 0, gross)
		assert.Equal(t, 0, id)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
}

// func TestGetCharge(t *testing.T) {
// 	DataMock := new(mocks.CartData)

// }
