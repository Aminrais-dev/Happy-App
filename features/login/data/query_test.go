package data

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/login"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dbTest = config.InitDBTest()

func TestLoginUser(t *testing.T) {

	dbTest.AutoMigrate(&User{})
	repo := New(dbTest)

	t.Run("Get Data Login Success", func(t *testing.T) {
		input := login.Core{Email: "amin@gmail.com", Password: "amin"}
		data, err := repo.LoginUser(input.Email)
		assert.Nil(t, err)
		assert.Equal(t, input.Email, data.Email)
	})

	t.Run("Get Data Login Failed", func(t *testing.T) {
		input := login.Core{Email: "amin12@mail.id", Password: "amin"}
		data, err := repo.LoginUser(input.Email)
		assert.NotNil(t, err)
		assert.NotEqual(t, input.Email, data.Email)
	})

}
