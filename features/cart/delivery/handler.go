package delivery

import (
	"capstone/happyApp/features/cart"
	"capstone/happyApp/middlewares"
	"capstone/happyApp/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	From cart.UsecaseInterface
}

func New(e *echo.Echo, data cart.UsecaseInterface) {
	handler := &Delivery{
		From: data,
	}

	e.POST("/cart/:productid", handler.AddToCart, middlewares.JWTMiddleware())

}

func (user *Delivery) AddToCart(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	productid, err := strconv.Atoi(c.Param("communityid"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	msg, ers := user.From.AddToCart(userid, productid)
	if ers != nil {
		return c.JSON(40, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}
