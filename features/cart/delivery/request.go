package delivery

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/cart"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Request struct {
	Productid int `json:"productid" form:"productid"`
}

type RequestHistory struct {
	CartID       []int  `json:"cartid" form:"cartid"`
	Street       string `json:"street" form:"street"`
	City         string `json:"city" form:"city"`
	Province     string `json:"province" form:"province"`
	Type_Payment string `json:"type_payment" form:"type_payment"`
}

func (data *RequestHistory) ToCore() cart.CoreHistory {
	return cart.CoreHistory{
		Carts:        data.CartID,
		Street:       data.Street,
		City:         data.City,
		Province:     data.Province,
		Type_Payment: data.Type_Payment,
	}
}

func ToCoreMandiri(data coreapi.ChargeReq) coreapi.ChargeReq {
	return coreapi.ChargeReq{
		PaymentType: "echannel",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  data.TransactionDetails.OrderID,
			GrossAmt: data.TransactionDetails.GrossAmt,
		},
		EChannel: &coreapi.EChannelDetail{
			BillInfo1: data.EChannel.BillInfo1,
			BillInfo2: data.EChannel.BillInfo2,
		},
	}
}

func ToCoreBCA(data coreapi.ChargeReq) coreapi.ChargeReq {
	return coreapi.ChargeReq{
		PaymentType: "bank_transfer",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  data.TransactionDetails.OrderID,
			GrossAmt: data.TransactionDetails.GrossAmt,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: data.BankTransfer.Bank,
		},
	}
}
func ToCoreGopay(data coreapi.ChargeReq) coreapi.ChargeReq {
	return coreapi.ChargeReq{
		PaymentType: "gopay",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  data.TransactionDetails.OrderID,
			GrossAmt: data.TransactionDetails.GrossAmt,
		},
	}
}

func ToCoreMidtransBank(dataReq coreapi.ChargeReq) coreapi.ChargeReq {

	return coreapi.ChargeReq{
		PaymentType: "bank_transfer",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  dataReq.TransactionDetails.OrderID,
			GrossAmt: dataReq.TransactionDetails.GrossAmt,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: dataReq.BankTransfer.Bank,
		},
	}
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
