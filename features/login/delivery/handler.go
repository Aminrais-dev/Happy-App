package delivery

import (
	"capstone/happyApp/features/login"
	"capstone/happyApp/utils/helper"

	"github.com/labstack/echo/v4"
)

type loginDelivery struct {
	loginUsecase login.UsecaseInterface
}

func New(e *echo.Echo, usecase login.UsecaseInterface) {

	handler := loginDelivery{
		loginUsecase: usecase,
	}

	e.POST("/login", handler.loginUser)

}

func (delivery *loginDelivery) loginUser(c echo.Context) error {

	var req Request
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("wrong request"))
	}

	str, err := delivery.loginUsecase.LoginAuthorized(req.Email, req.Password)
	if str == "please input email and password" || str == "email not found" || str == "wrong password" {
		return c.JSON(400, helper.FailedResponseHelper(str))
	} else if err != nil {
		return c.JSON(500, helper.FailedResponseHelper(str))
	} else {
		return c.JSON(200, map[string]interface{}{
			"access_token": str,
			"massage":      "login success",
			"status":       "success",
		})
	}

}
