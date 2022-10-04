package delivery

import (
	"capstone/happyApp/features/cart"

	"github.com/midtrans/midtrans-go/coreapi"
)

type ResposeCommunity struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	Logo         string `json:"logo"`
	Members      int64  `json:"members"`
}

type ResponseCart struct {
	ID           uint   `json:"cartid"`
	ProductID    uint   `json:"productid"`
	Name         string `json:"name"`
	Descriptions string `json:"descriptions"`
	Photo        string `json:"photo"`
	Price        int    `json:"price"`
}

type VAnumbers struct {
	BankTransfer string `json:"bank"`
	VAnumber     string `json:"va_number"`
}

type Actions struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Url    string `json:"url"`
}

type ResponseChargeMandiri struct {
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
	GroosAmt          string `json:"gross_amount"`
	BillKey           string `json:"bill_key"`
	BillerCode        string `json:"biller_code"`
}

type ResponseChargeBCA struct {
	TransactionTime   string    `json:"transaction_time"`
	TransactionStatus string    `json:"transaction_status"`
	PaymentType       string    `json:"payment_type"`
	OrderID           string    `json:"order_id"`
	GroosAmt          string    `json:"gross_amount"`
	VAnumbers         VAnumbers `json:"va_numbers"`
}

type ResponseChargeGopay struct {
	TransactionTime   string    `json:"transaction_time"`
	TransactionStatus string    `json:"transaction_status"`
	PaymentType       string    `json:"payment_type"`
	OrderID           string    `json:"order_id"`
	GroosAmt          string    `json:"gross_amount"`
	Actions           []Actions `json:"actions"`
}

type ChargeResponse struct {
	TransactionTime   string    `json:"transaction_time"`
	TransactionStatus string    `json:"transaction_status"`
	PaymentType       string    `json:"payment_type"`
	VAnumbers         VAnumbers `json:"va_numbers"`
	OrderID           string    `json:"order_id"`
	GroosAmt          string    `json:"gross_amount"`
}

func CoreToResCommunity(data cart.CoreCommunity) ResposeCommunity {
	return ResposeCommunity{
		ID:           data.ID,
		Title:        data.Title,
		Descriptions: data.Descriptions,
		Logo:         data.Logo,
		Members:      data.Members,
	}
}

func CoreToResponseCart(data cart.CoreCart) ResponseCart {
	return ResponseCart{
		ID:           data.ID,
		ProductID:    data.ProductID,
		Name:         data.Name,
		Descriptions: data.Descriptions,
		Photo:        data.Photo,
		Price:        data.Price,
	}
}

func CoreToResponseCartList(data []cart.CoreCart) []ResponseCart {
	var list []ResponseCart
	for _, v := range data {
		list = append(list, CoreToResponseCart(v))
	}

	return list
}

func ToResponseBCA(data coreapi.ChargeResponse) ResponseChargeBCA {
	return ResponseChargeBCA{
		TransactionTime:   data.TransactionTime,
		TransactionStatus: data.TransactionStatus,
		PaymentType:       data.PaymentType,
		OrderID:           data.OrderID,
		GroosAmt:          data.GrossAmount,
		VAnumbers: VAnumbers{
			BankTransfer: data.VaNumbers[0].Bank,
			VAnumber:     data.VaNumbers[0].VANumber,
		},
	}
}

func ToResponseMandiri(data coreapi.ChargeResponse) ResponseChargeMandiri {
	return ResponseChargeMandiri{
		TransactionTime:   data.TransactionTime,
		TransactionStatus: data.TransactionStatus,
		PaymentType:       data.PaymentType,
		OrderID:           data.OrderID,
		GroosAmt:          data.GrossAmount,
		BillKey:           data.BillKey,
		BillerCode:        data.BillerCode,
	}
}

func ToResponseGopay(data coreapi.ChargeResponse) ResponseChargeGopay {
	var actions []Actions
	for _, v := range data.Actions {
		actions = append(actions, Actions{
			Name:   v.Name,
			Method: v.Method,
			Url:    v.URL,
		})
	}
	return ResponseChargeGopay{
		TransactionTime:   data.TransactionID,
		TransactionStatus: data.TransactionStatus,
		PaymentType:       data.PaymentType,
		OrderID:           data.OrderID,
		GroosAmt:          data.GrossAmount,
		Actions:           actions,
	}
}

func ToChargeMidtrans(data coreapi.ChargeResponse) ChargeResponse {
	return ChargeResponse{
		TransactionTime:   data.TransactionTime,
		TransactionStatus: data.TransactionStatus,
		PaymentType:       data.PaymentType,
		VAnumbers: VAnumbers{
			BankTransfer: data.VaNumbers[0].Bank,
			VAnumber:     data.VaNumbers[0].VANumber,
		},
		OrderID:  data.OrderID,
		GroosAmt: data.GrossAmount,
	}
}
