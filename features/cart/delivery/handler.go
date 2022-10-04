package delivery

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/cart"
	"capstone/happyApp/middlewares"
	"capstone/happyApp/utils/helper"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Delivery struct {
	From cart.UsecaseInterface
}

var transaction coreapi.Client

func New(e *echo.Echo, data cart.UsecaseInterface) {
	handler := &Delivery{
		From: data,
	}

	e.POST("/cart", handler.AddToCart, middlewares.JWTMiddleware())
	e.GET("/cart", handler.GetCart, middlewares.JWTMiddleware())
	e.DELETE("/cart/:cartid", handler.DeleteCart, middlewares.JWTMiddleware())
	e.POST("/checkout", handler.BuyInCart, middlewares.JWTMiddleware())

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

func (user *Delivery) BuyInCart(c echo.Context) error {
	midtrans.ServerKey = config.MidtransServerKey()
	transaction.New(midtrans.ServerKey, midtrans.Sandbox)

	userid := middlewares.ExtractToken(c)
	var buy RequestHistory
	erb := c.Bind(&buy)
	if erb != nil {
		return c.JSON(400, helper.FailedResponseHelper("Gagal Bind Data"))
	}

	if buy.City == "" || buy.Street == "" || buy.Province == "" || buy.CartID == nil || buy.Type_Payment == "" {
		return c.JSON(400, helper.FailedResponseHelper("Semua data harus di isi"))
	}

	if buy.Type_Payment == config.BCA_VIRTUAL_ACCOUNT || buy.Type_Payment == config.GOPAY || buy.Type_Payment == config.MANDIRI_VIRTUAL_ACCOUNT {
		// insert to transaction and get some data
		transid, gross, msg, err := user.From.InsertIntoTransaction(buy.ToCore())
		if err != nil {
			return c.JSON(400, helper.FailedResponseHelper(msg))
		}

		chargeresponse, OrderID, errtransfer := user.From.GetCharge(transid, gross, buy.Type_Payment, config.TRANSACTION)
		if errtransfer != nil {
			return c.JSON(400, helper.FailedResponseHelper(OrderID))
		}

		chargecore, msgch, errformatcharge := user.From.ChargeRequest(chargeresponse, buy.Type_Payment)
		if errformatcharge != nil {
			return c.JSON(400, helper.FailedResponseHelper(msgch))
		}

		midtransresp, errcharge := coreapi.ChargeTransaction(&chargecore)
		if errcharge != nil {
			return c.JSON(400, helper.FailedResponseHelper(errcharge.GetMessage()))
		}

		payment := cart.CorePayment{
			UserID:  uint(userid),
			OrderID: OrderID,
			Groos:   gross,
		}

		msgpayi, errpayid := user.From.InsertIntoPayment(payment)
		if errpayid != nil {
			return c.JSON(400, helper.FailedResponseHelper(msgpayi))
		}

		if buy.Type_Payment == config.BCA_VIRTUAL_ACCOUNT {
			return c.JSON(200, ToResponseBCA(*midtransresp))
		} else if buy.Type_Payment == config.MANDIRI_VIRTUAL_ACCOUNT {
			return c.JSON(200, ToResponseMandiri(*midtransresp))
		} else if buy.Type_Payment == config.GOPAY {
			return c.JSON(200, ToResponseGopay(*midtransresp))
		} else {
			return c.JSON(200, ToChargeMidtrans(*midtransresp))
		}
	} else {
		return c.JSON(400, helper.FailedResponseHelper(fmt.Sprintf("Kami Hanya Menyediakan %s,%s,%s Saja", config.BCA_VIRTUAL_ACCOUNT, config.MANDIRI_VIRTUAL_ACCOUNT, config.GOPAY)))
	}
}
