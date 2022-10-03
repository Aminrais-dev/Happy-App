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

	e.POST("/cart", handler.AddToCart, middlewares.JWTMiddleware())
	e.GET("/cart", handler.GetCart, middlewares.JWTMiddleware())
	e.DELETE("/cart/:cartid", handler.DeleteCart, middlewares.JWTMiddleware())

}

func (user *Delivery) AddToCart(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	var req Request
	errbind := c.Bind(&req)
	if errbind != nil {
		return c.JSON(400, helper.FailedResponseHelper("Gagal Bind Data"))
	}
	productid := req.Productid

	msg, ers := user.From.AddToCart(userid, productid)
	if ers != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}

func (user *Delivery) GetCart(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	communityid, err := strconv.Atoi(c.QueryParam("communityid"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	corecommonity, listcart, msg, errs := user.From.GetCartList(userid, communityid)
	if errs != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessCartResponseHelper(msg, CoreToResCommunity(corecommonity), CoreToResponseCartList(listcart)))
}

func (user *Delivery) DeleteCart(c echo.Context) error {
	// userid := middlewares.ExtractToken(c)
	cartid, err := strconv.Atoi(c.Param("cartid"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	msg, ers := user.From.DeleteFromCart(cartid)
	if ers != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}
