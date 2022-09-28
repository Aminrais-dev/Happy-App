package delivery

import (
	"capstone/happyApp/features/user"
	"capstone/happyApp/utils/helper"

	"github.com/labstack/echo/v4"
)

type userDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, data user.UsecaseInterface) {
	handler := &userDelivery{
		userUsecase: data,
	}

	e.POST("/register", handler.CreateUser)
}

func (delivery *userDelivery) CreateUser(c echo.Context) error {

	var reqData Request
	err := c.Bind(&reqData)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error Bind"))
	}

	row := delivery.userUsecase.PostUser(reqData.reqToCore())
	if row == -2 {
		return c.JSON(400, helper.FailedResponseHelper("please input all request"))
	} else if row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("failed sign up"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success sign up"))

}
