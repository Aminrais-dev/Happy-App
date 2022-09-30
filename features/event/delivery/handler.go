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
	e.GET("/community/:id/event", handler.GetEventListCommunity, middlewares.JWTMiddleware())
	e.GET("/event/:id", handler.GetEventDetailbyId, middlewares.JWTMiddleware())

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

func (delivery *eventDelivery) GetEventListCommunity(c echo.Context) error {

	id := c.Param("id")
	idComu, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}
	search := c.QueryParam("title")
	userId := middlewares.ExtractToken(c)

	data, errGet := delivery.eventUsecase.GetEventComu(search, idComu, userId)
	if errGet != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to get list event in community"))
	} else if data.ID != uint(idComu) {
		return c.JSON(404, helper.FailedResponseHelper("community not found"))
	}

	return c.JSON(200, data)
}

func (delivery *eventDelivery) GetEventDetailbyId(c echo.Context) error {

	id := c.Param("id")
	idEvent, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	userId := middlewares.ExtractToken(c)

	data, errGet := delivery.eventUsecase.GetEventDetail(idEvent, userId)
	if errGet != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to get event detail"))
	} else if data.ID != uint(idEvent) {
		return c.JSON(404, helper.FailedResponseHelper("event not found"))
	}

	return c.JSON(200, data)

}
