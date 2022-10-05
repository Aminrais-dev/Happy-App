package usecase

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/community"
	"capstone/happyApp/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMakeCommunity(t *testing.T) {
	DataMock := new(mocks.CommunityData)
	community := community.CoreCommunity{Title: "Pedang", Descriptions: "Para Penggila Pedang", Logo: config.DEFAULT_COMMUNITY}

	t.Run("Sukses", func(t *testing.T) {
		DataMock.On("Insert", mock.Anything, mock.Anything).Return("pesan", nil).Once()
		logic := New(DataMock)
		msg, err := logic.AddNewCommunity(1, community)
		assert.Equal(t, msg, "pesan")
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestGetList(t *testing.T) {
	DataMock := new(mocks.CommunityData)
	community := []community.CoreCommunity{{Title: "Pedang", Descriptions: "Para Penggila Pedang", Logo: config.DEFAULT_COMMUNITY}, {Title: "Genshin", Descriptions: "Para Penggila Pedang", Logo: config.DEFAULT_COMMUNITY}}

	t.Run("Success", func(t *testing.T) {
		DataMock.On("SelectList").Return(community, "Sebuah Pesan", nil).Once()
		logic := New(DataMock)
		list, msg, err := logic.GetListCommunity()
		assert.Equal(t, list[1].Title, "Genshin")
		assert.Equal(t, msg, "Sebuah Pesan")
		assert.Nil(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestMembers(t *testing.T) {
	DataMock := new(mocks.CommunityData)
	members := []string{"Muhammad", "Nawi", "Husen"}

	t.Run("Success", func(t *testing.T) {
		DataMock.On("SelectMembers", mock.Anything).Return(members, "Pesan", nil)
		logic := New(DataMock)
		members, msg, err := logic.GetMembers(1)
		assert.NoError(t, err)
		assert.Equal(t, len(members), 3)
		assert.Equal(t, msg, "Pesan")
		DataMock.AssertExpectations(t)
	})
}

func TestLeave(t *testing.T) {
	DataMock := new(mocks.CommunityData)

	t.Run("Sucess", func(t *testing.T) {
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan Tidak Di Tangkap", errors.New("Logica Terbalik")).Once()
		DataMock.On("GetUserRole", mock.Anything, mock.Anything).Return("role", nil).Once()
		DataMock.On("Delete", mock.Anything, mock.Anything).Return(int64(0), "Pesan dari Leave", nil).Once()
		DataMock.On("DeleteCommunity", mock.Anything).Return("Pesan dari Community", nil).Once()
		logic := New(DataMock)
		msg, err := logic.Leave(1, 1)
		assert.Equal(t, "Pesan dari Leave Pesan dari Community", msg)
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Success Kedua", func(t *testing.T) {
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan Tidak Di Tangkap", errors.New("Logica Terbalik")).Once()
		DataMock.On("GetUserRole", mock.Anything, mock.Anything).Return("admin", nil).Once()
		DataMock.On("Delete", mock.Anything, mock.Anything).Return(int64(2), "Pesan dari Leave", nil).Once()
		DataMock.On("ChangeAdmin", mock.Anything).Return("Nawi", "Pesan", nil).Once()
		logic := New(DataMock)
		msg, err := logic.Leave(1, 1)
		assert.Equal(t, "Community akan di wariskan ke Nawi", msg)
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed", func(t *testing.T) {
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan Tidak Di Tangkap", nil).Once()
		logic := New(DataMock)
		msg, err := logic.Leave(1, 1)
		assert.Equal(t, "Hanya Member Dari Community Yang Bisa Leave", msg)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed Kedua", func(t *testing.T) {
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan Tidak Di Tangkap", errors.New("Logica Terbalik")).Once()
		DataMock.On("GetUserRole", mock.Anything, mock.Anything).Return("admin", errors.New("ada error")).Once()
		logic := New(DataMock)
		msg, err := logic.Leave(1, 1)
		assert.Equal(t, "Gagal mendapatkan role user", msg)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed Kedua", func(t *testing.T) {
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan Tidak Di Tangkap", errors.New("Logica Terbalik")).Once()
		DataMock.On("GetUserRole", mock.Anything, mock.Anything).Return("admin", nil).Once()
		DataMock.On("Delete", mock.Anything, mock.Anything).Return(int64(0), "Pesan dari", nil).Once()
		DataMock.On("DeleteCommunity", mock.Anything).Return("Delete Community", nil).Once()
		logic := New(DataMock)
		msg, err := logic.Leave(1, 1)
		assert.Equal(t, "Pesan dari Delete Community", msg)
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
}
