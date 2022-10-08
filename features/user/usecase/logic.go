package usecase

import (
	"capstone/happyApp/features/user"
	"capstone/happyApp/utils/helper"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type usecaseInterface struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &usecaseInterface{
		userData: data,
	}
}

func (usecase *usecaseInterface) PostUser(data user.CoreUser) int {

	if data.Name == "" || data.Email == "" || data.Gender == "" || data.Password == "" || data.Username == "" {
		return -2
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	data.Password = string(hashPass)

	status := usecase.userData.CheckStatus(data.Email, 0)
	data.Status = status

	row, name := usecase.userData.InsertUser(data)
	if row > 0 {

		bodyEmail := helper.BodyEmail{
			Name: name,
			Url:  "https://tugas.website/user/verify/" + strconv.Itoa(row),
		}

		helper.SendEmailVerify(data.Email, "Email Verification", bodyEmail)

	}
	return row

}

func (usecase *usecaseInterface) DeleteUser(id int) int {

	row := usecase.userData.DelUser(id)
	return row

}

func (usecase *usecaseInterface) UpdateUser(data user.CoreUser) int {

	if data.Password != "" {
		hashPass, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		data.Password = string(hashPass)

	}
	row := usecase.userData.UpdtUser(data)
	return row

}

func (usecase *usecaseInterface) GetUser(id int) (user.CoreUser, []user.CommunityProfile, error) {

	data, comu, err := usecase.userData.SelectUser(id)
	if err != nil {
		return user.CoreUser{}, nil, err
	}

	return data, comu, nil

}

func (usecase *usecaseInterface) UpdateStatus(id int) int {

	status := usecase.userData.CheckStatus("", id)

	row := usecase.userData.UpdtStatus(id, status)

	return row

}
