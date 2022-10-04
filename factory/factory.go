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

	communityData "capstone/happyApp/features/community/data"
	communityDelivery "capstone/happyApp/features/community/delivery"
	communityUsecase "capstone/happyApp/features/community/usecase"

	eventData "capstone/happyApp/features/event/data"
	eventDelivery "capstone/happyApp/features/event/delivery"
	eventUsecase "capstone/happyApp/features/event/usecase"

	productData "capstone/happyApp/features/product/data"
	productDelivery "capstone/happyApp/features/product/delivery"
	productUsecase "capstone/happyApp/features/product/usecase"

	cartData "capstone/happyApp/features/cart/data"
	cartDelivery "capstone/happyApp/features/cart/delivery"
	cartUsecase "capstone/happyApp/features/cart/usecase"

	midtransData "capstone/happyApp/features/midtrans/data"
	midtransDelivery "capstone/happyApp/features/midtrans/delivery"
	midtransUsecase "capstone/happyApp/features/midtrans/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	productDataFactory := productData.New(db)
	productUsecaseFactory := productUsecase.New(productDataFactory)
	productDelivery.New(e, productUsecaseFactory)

	eventDataFactory := eventData.New(db)
	eventUsecaseFactory := eventUsecase.New(eventDataFactory)
	eventDelivery.New(e, eventUsecaseFactory)

	communityDataFactory := communityData.New(db)
	communityUsecaseFactory := communityUsecase.New(communityDataFactory)
	communityDelivery.New(e, communityUsecaseFactory)

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)

	cartDataFactory := cartData.New(db)
	cartUsecaseFactory := cartUsecase.New(cartDataFactory)
	cartDelivery.New(e, cartUsecaseFactory)

	midtransDataFactory := midtransData.New(db)
	midtransUsecaseFactory := midtransUsecase.New(midtransDataFactory)
	midtransDelivery.New(e, midtransUsecaseFactory)
}
