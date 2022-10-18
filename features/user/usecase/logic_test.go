package usecase

import (
	"capstone/happyApp/config"
	"capstone/happyApp/mocks"
	"errors"
	"testing"

	"capstone/happyApp/features/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProfile(t *testing.T) {

	userMock := new(mocks.DataUser)
	var returnData2 []user.CommunityProfile
	returnData2 = append(returnData2, user.CommunityProfile{ID: 1, Title: "community saya", Logo: "https://logo", Role: "member"})
	returnData := user.CoreUser{ID: 1, Name: "amin", Username: "aminrais89", Gender: "Male", Email: "amin@a.com", Password: "$2a$10$3qSIi7BiTknraN3A9tRX/eoI4N9yuln/oWI8Ft9KcrZNF3ec6jIHK", Photo: "https://photo_profile"}

	token := 1

	t.Run("Get profile success", func(t *testing.T) {
		userMock.On("SelectUser", token).Return(returnData, returnData2, nil).Once()

		useCase := New(userMock)
		res, res2, _ := useCase.GetUser(token)
		assert.Equal(t, token, int(res.ID))
		assert.NotEqual(t, "", res2[0].Role)
		userMock.AssertExpectations(t)

	})

	t.Run("Get profile failed", func(t *testing.T) {

		var kosong []user.CommunityProfile
		kosong = append(kosong, user.CommunityProfile{})
		userMock.On("SelectUser", token).Return(user.CoreUser{}, kosong, errors.New("error")).Once()

		useCase := New(userMock)
		param := 1
		res, _, err := useCase.GetUser(token)
		assert.Error(t, err)
		assert.Equal(t, "", kosong[0].Role)
		assert.NotEqual(t, param, int(res.ID))
		userMock.AssertExpectations(t)

	})

}

func TestUpdateUser(t *testing.T) {

	userMock := new(mocks.DataUser)
	input := user.CoreUser{ID: 1, Name: "amin", Username: "aminrais89", Gender: "Male", Email: "amin@a.com"}

	t.Run("update succes", func(t *testing.T) {

		userMock.On("CheckUsername", mock.Anything).Return(1).Once()
		userMock.On("UpdtUser", input).Return(1).Once()

		useCase := New(userMock)
		res := useCase.UpdateUser(input)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("update failed", func(t *testing.T) {

		userMock.On("CheckUsername", mock.Anything).Return(1).Once()
		userMock.On("UpdtUser", input).Return(-1).Once()

		useCase := New(userMock)
		res := useCase.UpdateUser(input)
		assert.Equal(t, -1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("update failed username", func(t *testing.T) {

		userMock.On("CheckUsername", mock.Anything).Return(-4).Once()

		useCase := New(userMock)
		res := useCase.UpdateUser(input)
		assert.Equal(t, -4, res)
		userMock.AssertExpectations(t)

	})

	t.Run("update success", func(t *testing.T) {

		userMock.On("CheckUsername", mock.Anything).Return(1).Once()
		userMock.On("UpdtUser", mock.Anything).Return(1).Once()

		input.Password = "123"
		useCase := New(userMock)
		res := useCase.UpdateUser(input)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)

	})
}

func TestPostData(t *testing.T) {

	userMock := new(mocks.DataUser)
	input := user.CoreUser{Name: "amin", Username: "aminrais89", Gender: "Male", Email: "muhammadamin.rais13@gmail.com", Password: "12345"}

	t.Run("create success", func(t *testing.T) {

		userMock.On("CheckStatus", mock.Anything, mock.AnythingOfType("int")).Return(config.DEFAULT_PROFILE).Once()
		userMock.On("CheckUsername", mock.Anything).Return(1).Once()
		userMock.On("InsertUser", mock.Anything).Return(1).Once()

		useCase := New(userMock)
		res := useCase.PostUser(input)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)
	})

	t.Run("create failed", func(t *testing.T) {

		userMock.On("CheckStatus", mock.Anything, mock.AnythingOfType("int")).Return(config.DEFAULT_PROFILE).Once()
		userMock.On("CheckUsername", mock.Anything).Return(1).Once()
		userMock.On("InsertUser", mock.Anything).Return(-1).Once()

		useCase := New(userMock)
		res := useCase.PostUser(input)
		assert.Equal(t, -1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("create failed status verify", func(t *testing.T) {

		userMock.On("CheckStatus", mock.Anything, mock.AnythingOfType("int")).Return(config.VERIFY).Once()
		userMock.On("CheckUsername", mock.Anything).Return(1).Once()
		userMock.On("InsertUser", mock.Anything).Return(-3).Once()

		useCase := New(userMock)
		res := useCase.PostUser(input)
		assert.Equal(t, -3, res)
		userMock.AssertExpectations(t)

	})

	t.Run("create failed status verify", func(t *testing.T) {

		userMock.On("CheckStatus", mock.Anything, mock.AnythingOfType("int")).Return(config.DEFAULT_PROFILE).Once()
		userMock.On("CheckUsername", mock.Anything).Return(-4).Once()

		useCase := New(userMock)
		res := useCase.PostUser(input)
		assert.Equal(t, -4, res)
		userMock.AssertExpectations(t)

	})

	t.Run("create failed because input name not filled", func(t *testing.T) {

		input.Name = ""
		input.Password = ""
		useCase := New(userMock)
		res := useCase.PostUser(input)
		assert.Equal(t, -2, res)
		userMock.AssertExpectations(t)

	})

}

func TestDelete(t *testing.T) {

	userMock := new(mocks.DataUser)
	param := 1

	t.Run("delete succes", func(t *testing.T) {

		userMock.On("DelUser", param).Return(1).Once()

		useCase := New(userMock)
		res := useCase.DeleteUser(param)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("delete failed", func(t *testing.T) {

		userMock.On("DelUser", param).Return(-1).Once()

		useCase := New(userMock)
		res := useCase.DeleteUser(param)
		assert.Equal(t, -1, res)
		userMock.AssertExpectations(t)

	})

}

func TestUpdateStatus(t *testing.T) {

	userMock := new(mocks.DataUser)
	param := 1

	t.Run("Update status success", func(t *testing.T) {

		userMock.On("CheckStatus", mock.Anything, mock.Anything).Return(config.DEFAULT_STATUS).Once()
		userMock.On("UpdtStatus", param, mock.Anything).Return(1).Once()

		useCase := New(userMock)
		res := useCase.UpdateStatus(param)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("Update status failed", func(t *testing.T) {

		userMock.On("CheckStatus", mock.Anything, mock.Anything).Return(config.VERIFY).Once()
		userMock.On("UpdtStatus", param, mock.Anything).Return(-2).Once()

		useCase := New(userMock)
		res := useCase.UpdateStatus(param)
		assert.Equal(t, -2, res)
		userMock.AssertExpectations(t)

	})

	t.Run("Update status failed", func(t *testing.T) {

		userMock.On("CheckStatus", mock.Anything, mock.Anything).Return(config.DEFAULT_STATUS).Once()
		userMock.On("UpdtStatus", param, mock.Anything).Return(-1).Once()

		useCase := New(userMock)
		res := useCase.UpdateStatus(param)
		assert.Equal(t, -1, res)
		userMock.AssertExpectations(t)

	})

}
