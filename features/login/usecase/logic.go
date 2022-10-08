package usecase

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/login"
	"capstone/happyApp/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	loginData login.DataInterface
}

func New(data login.DataInterface) login.UsecaseInterface {
	return &loginUsecase{
		loginData: data,
	}
}

func (usecase *loginUsecase) LoginAuthorized(email, password string) (string, error) {

	if email == "" || password == "" {
		return "please input email and password", nil
	}

	results, errEmail := usecase.loginData.LoginUser(email)

	if errEmail != nil {
		return "email not found", nil
	}

	errPw := bcrypt.CompareHashAndPassword([]byte(results.Password), []byte(password))
	if errPw != nil {
		return "wrong password", nil
	}

	if results.Status != config.VERIFY {
		return "please confirm your account in gmail", nil
	}

	token, errToken := middlewares.CreateToken(int(results.ID))

	return token, errToken

}
