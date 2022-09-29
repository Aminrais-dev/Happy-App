package delivery

import (
	"capstone/happyApp/features/event"
	"capstone/happyApp/middlewares"
	"capstone/happyApp/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type eventDelivery struct {
	eventUsecase event.UsecaseInterface
}

func New(e *echo.Echo, data event.UsecaseInterface) {
	handler := &eventDelivery{
		eventUsecase: data,
	}

	e.POST("community/:id/event", handler.PostEventCommunity, middlewares.JWTMiddleware())
	e.GET("/event", handler.GetEventList)

}

func (delivery *eventDelivery) PostEventCommunity(c echo.Context) error {

	id := c.Param("id")
	idComu, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	var Req Request
	errBind := c.Bind(&Req)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	Req.CommunityID = uint(idComu)
	idToken := middlewares.ExtractToken(c)

	row := delivery.eventUsecase.PostEvent(Req.resToCore(), idToken)
	if row == -3 {
		return c.JSON(400, helper.FailedResponseHelper("please input all request"))
	} else if row == -2 {
		return c.JSON(400, helper.FailedResponseHelper("not have access in community"))
	} else if row == -1 {
		return c.JSON(400, helper.FailedResponseHelper("failed to create event"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success create event"))

}

func (delivery *eventDelivery) GetEventList(c echo.Context) error {

	search := c.QueryParam("title")

	data, err := delivery.eventUsecase.GetEvent(search)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to get list event"))
	}

	return c.JSON(200, map[string]interface{}{
		"event":   data,
		"massage": "success get list event",
	})
}
