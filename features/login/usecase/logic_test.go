package usecase

import (
	"capstone/happyApp/features/login"
	"capstone/happyApp/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	repo := new(mocks.DataLogin)
	dataInput := login.Core{ID: 1, Email: "amin@mail.id", Password: "$2a$10$3qSIi7BiTknraN3A9tRX/eoI4N9yuln/oWI8Ft9KcrZNF3ec6jIHK"}

	t.Run("success password.", func(t *testing.T) {
		repo.On("LoginUser", mock.Anything, mock.Anything).Return(dataInput, nil).Once()

		usecase := New(repo)
		result, err := usecase.LoginAuthorized("amin@mail.id", "12345")
		assert.NotEqual(t, "please input email and password", result)
		assert.NotEqual(t, "email not found", result)
		assert.NotEqual(t, "wrong password", result)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("empty passwor", func(t *testing.T) {
		usecase := New(repo)
		result, _ := usecase.LoginAuthorized("amin@mail.id", "")
		assert.Equal(t, "please input email and password", result)
		repo.AssertExpectations(t)
	})

	t.Run("Empty email.", func(t *testing.T) {
		usecase := New(repo)
		result, _ := usecase.LoginAuthorized("", "amin123")
		assert.Equal(t, "please input email and password", result)
		repo.AssertExpectations(t)
	})

	t.Run("Email not found", func(t *testing.T) {
		repo.On("LoginUser", mock.Anything).Return(login.Core{}, errors.New("error")).Once()

		usecase := New(repo)
		result, _ := usecase.LoginAuthorized("ridho@mail.uk", "888ridho")
		assert.Equal(t, "email not found", result)
		repo.AssertExpectations(t)
	})

	t.Run("Wrong Password.", func(t *testing.T) {
		repo.On("LoginUser", mock.Anything).Return(login.Core{}, nil).Once()

		usecase := New(repo)
		result, _ := usecase.LoginAuthorized("amin@mail.id", "amin123")
		assert.Equal(t, "wrong password", result)
		repo.AssertExpectations(t)
	})

}
