package delivery

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/event"
	"errors"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Request struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"descriptions" form:"descriptions"`
	Price       uint64 `json:"price" form:"price"`
	Date        string `json:"date_event" form:"date_event"`
	Location    string `json:"location" form:"location"`
	CommunityID uint
}

type RequestPayment struct {
	OrderID      string
	GrossAmount  uint64
	Payment_type string `json:"payment_type" form:"payment_type"`
}

func (data *Request) resToCore() event.EventCore {

	var layout = "2006-01-02 15:04:05 MST"
	wib := data.Date + " WIB"
	date, _ := time.Parse(layout, wib)

	return event.EventCore{
		Title:       data.Title,
		Description: data.Description,
		Price:       data.Price,
		Date:        date,
		Location:    data.Location,
		CommunityID: data.CommunityID,
	}
}

func toMidtrans(req RequestPayment) (coreapi.ChargeReq, error) {

	var ReqMidtrans = coreapi.ChargeReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.OrderID,
			GrossAmt: int64(req.GrossAmount),
		},
	}

	if req.Payment_type == config.GOPAY {
		ReqMidtrans.PaymentType = "gopay"

	} else if req.Payment_type == config.BCA_VIRTUAL_ACCOUNT {
		ReqMidtrans.PaymentType = "bank_transfer"
		ReqMidtrans.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		}

	} else if req.Payment_type == config.MANDIRI_VIRTUAL_ACCOUNT {
		ReqMidtrans.PaymentType = "echannel"
		ReqMidtrans.EChannel = &coreapi.EChannelDetail{
			BillInfo1: "Payment:",
			BillInfo2: "Online purchase",
		}

	} else {
		return coreapi.ChargeReq{}, errors.New("method not allowed")
	}

	return ReqMidtrans, nil
}

func typePayment(tipe string) string {

	if tipe == "GOPAY" {
		return config.GOPAY
	} else if tipe == "BCA Virtual Account" {
		return config.BCA_VIRTUAL_ACCOUNT
	} else if tipe == "Mandiri Virtual Account" {
		return config.MANDIRI_VIRTUAL_ACCOUNT
	}

	return ""
}
