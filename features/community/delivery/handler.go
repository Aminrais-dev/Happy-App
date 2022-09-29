package delivery

import (
	"capstone/happyApp/features/community"
	"capstone/happyApp/middlewares"
	"capstone/happyApp/utils/helper"
	"strconv"

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
	e.GET("/community/members/:communityid", handler.ListMembersCommunity, middlewares.JWTMiddleware())
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

func (user *Delivery) ListMembersCommunity(c echo.Context) error {
	communityid, err := strconv.Atoi(c.Param("communityid"))
	if err != nil {
		c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}
	members, msg, errs := user.From.GetMembers(communityid)
	if errs != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(msg, members))
}
