package usecase

import (
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

func (usecase *loginUsecase) LoginAuthorized(email, password string) string {

	if email == "" || password == "" {
		return "please input email and password"
	}

	results, errEmail := usecase.loginData.LoginUser(email)
	if errEmail != nil {
		return "email not found"
	}

	errPw := bcrypt.CompareHashAndPassword([]byte(results.Password), []byte(password))
	if errPw != nil {
		return "wrong password"
	}

	token, errToken := middlewares.CreateToken(int(results.ID))

	if errToken != nil {
		return "error to created token"
	}

	return token

}
