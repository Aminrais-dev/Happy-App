package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userData "capstone/happyApp/features/user/data"
	userDelivery "capstone/happyApp/features/user/delivery"
	userUsecase "capstone/happyApp/features/user/usecase"

	loginData "capstone/happyApp/features/login/data"
	loginDelivery "capstone/happyApp/features/login/delivery"
	loginUsecase "capstone/happyApp/features/login/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)

}
