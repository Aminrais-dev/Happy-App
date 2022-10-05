package delivery

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/midtrans"
	"capstone/happyApp/utils/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	From midtrans.UsecaseInterface
}

func New(e *echo.Echo, data midtrans.UsecaseInterface) {
	handler := &Delivery{
		From: data,
	}
	e.POST("/midtrans/webhook", handler.UpdateAData)
}

func (user *Delivery) UpdateAData(c echo.Context) error {
	midreq := MidtransHookRequest{}
	errb := c.Bind(&midreq)
	if errb != nil {
		return c.JSON(400, helper.FailedResponseHelper("Gagal Bind"))
	}

	table := strings.Split(midreq.OrderID, "-")
	if table[0] == config.TRANSACTION {
		msg, err := user.From.WeebHookTransaction(midreq.OrderID, midreq.TransactionStatus)
		if err != nil {
			return c.JSON(400, helper.FailedResponseHelper(msg))
		}

		return c.JSON(200, helper.SuccessResponseHelper(msg))

	} else if table[0] == config.EVENT {
		msg, err := user.From.WeebHookJoinEvent(midreq.OrderID, midreq.TransactionStatus)
		if err != nil {
			c.JSON(400, helper.FailedResponseHelper(msg))
		}

		return c.JSON(200, helper.SuccessResponseHelper(msg))

	} else {
		return c.JSON(400, helper.FailedResponseHelper("Order Id Has A Different Format"))
	}
}
