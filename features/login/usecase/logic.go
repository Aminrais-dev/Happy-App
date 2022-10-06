package usecase

import (
	"capstone/happyApp/features/login"
	"capstone/happyApp/middlewares"
	"errors"

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
		return "please input email and password", errors.New("error email or password")
	}

	results, errEmail := usecase.loginData.LoginUser(email)
	if errEmail != nil {
		return "email not found", errEmail
	}

	errPw := bcrypt.CompareHashAndPassword([]byte(results.Password), []byte(password))
	if errPw != nil {
		return "wrong password", errPw
	}

	token, errToken := middlewares.CreateToken(int(results.ID))

	return token, errToken

}
