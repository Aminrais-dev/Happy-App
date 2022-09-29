package delivery

import (
	"capstone/happyApp/features/community"
	"capstone/happyApp/utils/helper"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	From community.UsecaseInterface
}

func New(e *echo.Echo, data community.UsecaseInterface) {
	handler := &Delivery{
		From: data,
	}
	e.GET("/community", handler.ListCommunity)
	// e.POST("/community", handler.AddCommunity, middlewares.JWTMiddleware())
}

// func (user *Delivery) AddCommunity(c echo.Context) error {
// 	userid := middlewares.ExtractToken(c)
// 	var reqcom Request
// 	errb := c.Bind(&reqcom)
// 	if errb != nil {
// 		return c.JSON(400, helper.FailedResponseHelper("Gagal Bind Data"))
// 	}
// }

func (user *Delivery) ListCommunity(c echo.Context) error {
	listcore, msg, err := user.From.GetListCommunity()
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(msg, ToResponseList(listcore)))
}
