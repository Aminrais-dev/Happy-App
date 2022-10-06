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
	t.Run("Failed Ketiga", func(t *testing.T) {
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan Tidak Di Tangkap", errors.New("Logica Terbalik")).Once()
		DataMock.On("GetUserRole", mock.Anything, mock.Anything).Return("admin", nil).Once()
		DataMock.On("Delete", mock.Anything, mock.Anything).Return(int64(0), "Pesan dari", nil).Once()
		DataMock.On("DeleteCommunity", mock.Anything).Return("Messege", errors.New("A Error")).Once()
		logic := New(DataMock)
		msg, err := logic.Leave(1, 1)
		assert.Equal(t, "Messege", msg)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed Keempat", func(t *testing.T) {
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan Tidak Di Tangkap", errors.New("Logica Terbalik")).Once()
		DataMock.On("GetUserRole", mock.Anything, mock.Anything).Return("admin", nil).Once()
		DataMock.On("Delete", mock.Anything, mock.Anything).Return(int64(5), "Pesan dari Leave", nil).Once()
		DataMock.On("ChangeAdmin", mock.Anything).Return("Nawi", "Gagal Mewariskan", errors.New("Error")).Once()
		logic := New(DataMock)
		msg, err := logic.Leave(1, 1)
		assert.Equal(t, "Gagal Mewariskan", msg)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestUpdateCommunity(t *testing.T) {
	DataMock := new(mocks.CommunityData)
	community := community.CoreCommunity{Title: "Pedang", Descriptions: "Para Penggila Pedang", Logo: config.DEFAULT_COMMUNITY}

	t.Run("Success", func(t *testing.T) {
		DataMock.On("GetUserRole", mock.Anything, mock.Anything).Return("admin", nil).Once()
		DataMock.On("UpdateCommunity", mock.Anything, mock.Anything).Return("Pesan", nil).Once()
		logic := New(DataMock)
		msg, err := logic.UpdateCommunity(1, community)
		assert.Equal(t, msg, "Pesan")
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed Pertama", func(t *testing.T) {
		DataMock.On("GetUserRole", mock.Anything, mock.Anything).Return("member", nil).Once()
		logic := New(DataMock)
		msg, err := logic.UpdateCommunity(1, community)
		assert.Equal(t, msg, "Hanya admin yang bisa mengupdate Community")
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed Kedua", func(t *testing.T) {
		DataMock.On("GetUserRole", mock.Anything, mock.Anything).Return("member", errors.New("A Error")).Once()
		logic := New(DataMock)
		msg, err := logic.UpdateCommunity(1, community)
		assert.Equal(t, msg, "Gagal mendapatkan role user")
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestCommunityFeed(t *testing.T) {
	DataMock := new(mocks.CommunityData)
	community := community.CoreCommunity{ID: 1, Title: "Pedang", Role: "member", Descriptions: "Para Penggila Pedang", Logo: config.DEFAULT_COMMUNITY, Members: 59, Feeds: []community.CoreFeed{
		{ID: 1, Name: "Nawi", Text: "hamm"},
		{ID: 2, Name: "Husen", Text: "himm"},
	}}
	t.Run("Success", func(t *testing.T) {
		DataMock.On("SelectCommunity", mock.Anything, mock.Anything).Return(community, "Pesan", nil)
		logic := New(DataMock)
		community, msg, err := logic.GetCommunityFeed(1, 1)
		assert.Equal(t, community.Feeds[0].Name, "Nawi")
		assert.Equal(t, msg, "Pesan")
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestJoinCommunity(t *testing.T) {
	DataMock := new(mocks.CommunityData)

	t.Run("Success", func(t *testing.T) {
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan", nil).Once()
		DataMock.On("InsertToJoin", mock.Anything, mock.Anything).Return("Berhasil Join Community", nil).Once()
		logic := New(DataMock)
		msg, err := logic.JoinCommunity(1, 1)
		assert.Equal(t, "Berhasil Join Community", msg)
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed", func(t *testing.T) {
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan Dari Check Join", errors.New("Error")).Once()
		logic := New(DataMock)
		msg, err := logic.JoinCommunity(1, 1)
		assert.Equal(t, "Pesan Dari Check Join", msg)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestPostFeed(t *testing.T) {
	DataMock := new(mocks.CommunityData)
	feed := community.CoreFeed{Text: "Ini adalah Feed", CommunityID: 1, UserID: 1}

	t.Run("Success", func(t *testing.T) {
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan", errors.New("Error")).Once()
		DataMock.On("InsertFeed", mock.Anything).Return("Pesan Dari Post", nil).Once()
		logic := New(DataMock)
		msg, err := logic.PostFeed(feed)
		assert.Equal(t, "Pesan Dari Post", msg)
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed", func(t *testing.T) {
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan", nil).Once()
		logic := New(DataMock)
		msg, err := logic.PostFeed(feed)
		assert.Equal(t, "Hanya Anggota dari Community Yang bisa Post Feed", msg)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestGetFeed(t *testing.T) {
	DataMock := new(mocks.CommunityData)
	feed := community.CoreFeed{ID: 1, Name: "Nawi", Text: "Ini adalah Feed", CommunityID: 1, UserID: 1, Comments: []community.CoreComment{
		{Name: "Husen", Text: "Sebuah Comment"}, {Name: "Husen", Text: "Sebuah Comment"}, {Name: "Husen", Text: "Sebuah Comment"},
	}}
	t.Run("Success", func(t *testing.T) {
		DataMock.On("SelectFeed", mock.Anything).Return(feed, "Success", nil)
		logic := New(DataMock)
		feed, msg, err := logic.GetDetailFeed(1)
		assert.Equal(t, "Ini adalah Feed", feed.Text)
		assert.Equal(t, "Success", msg)
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestAddComment(t *testing.T) {
	DataMock := new(mocks.CommunityData)

	t.Run("Success", func(t *testing.T) {
		DataMock.On("SelectCommunityIdWithFeed", mock.Anything).Return(1, nil).Once()
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan", errors.New("Anda Sudah Bergabung")).Once()
		DataMock.On("InsertComment", mock.Anything).Return("Pesan Keberhasilan", nil)
		logic := New(DataMock)
		msg, err := logic.AddComment(community.CoreComment{Text: "Ini Comment"})
		assert.Equal(t, "Pesan Keberhasilan", msg)
		assert.NoError(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed Pertama", func(t *testing.T) {
		DataMock.On("SelectCommunityIdWithFeed", mock.Anything).Return(1, errors.New("Error")).Once()
		logic := New(DataMock)
		msg, err := logic.AddComment(community.CoreComment{Text: "Ini Comment"})
		assert.Equal(t, "Gagal Mendapatkan Comunity Id", msg)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
	t.Run("Failed Kedua", func(t *testing.T) {
		DataMock.On("SelectCommunityIdWithFeed", mock.Anything).Return(1, nil).Once()
		DataMock.On("CheckJoin", mock.Anything, mock.Anything).Return("Pesan", nil).Once()
		logic := New(DataMock)
		msg, err := logic.AddComment(community.CoreComment{Text: "Ini Comment"})
		assert.Equal(t, "Hanya Anggota dari Community yang bisa melakukan Comment", msg)
		assert.Error(t, err)
		DataMock.AssertExpectations(t)
	})
}

func TestGetListCommunityWithParam(t *testing.T) {
	DataMock := new(mocks.CommunityData)
	community := []community.CoreCommunity{{Title: "Pedang", Descriptions: "Para Penggila Pedang", Logo: config.DEFAULT_COMMUNITY}, {Title: "Genshin", Descriptions: "Para Penggila Pedang", Logo: config.DEFAULT_COMMUNITY}}

	t.Run("Success", func(t *testing.T) {
		DataMock.On("SelectListCommunityWithParam", mock.Anything).Return(community, "Sebuah Pesan", nil).Once()
		logic := New(DataMock)
		list, msg, err := logic.GetListCommunityWithParam("da")
		assert.Equal(t, list[1].Title, "Genshin")
		assert.Equal(t, msg, "Sebuah Pesan")
		assert.Nil(t, err)
		DataMock.AssertExpectations(t)
	})
}
