package delivery

import (
	"capstone/happyApp/features/event"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type Payment struct {
	OrderID           string          `json:"order_id"`
	TransactionID     string          `json:"transaction_id"`
	PaymentMethod     string          `json:"payment_method"`
	BillNumber        string          `json:"bill_number,omitempty"`
	Bank              string          `json:"bank,omitempty"`
	GrossAmount       string          `json:"gross_amount"`
	TransactionTime   string          `json:"transaction_time"`
	TransactionStatus string          `json:"transaction_status"`
	Actions           []gopayResponse `json:"actions,omitempty"`
	Bill_key          string          `json:"bill_key,omitempty"`
	Biller_code       string          `json:"biller_code,omitempty"`
}

type gopayResponse struct {
	Name   string `json:"name" form:"name"`
	Method string `json:"method" form:"method"`
	Url    string `json:"url" form:"url"`
}

type ResponseListEvent struct {
	ID           uint      `json:"id"`
	Logo         string    `json:"logo"`
	Title        string    `json:"title"`
	Members      uint8     `json:"members"`
	Descriptions string    `json:"descriptions"`
	Date         time.Time `json:"date"`
	Price        int64     `json:"price"`
}

type ResponseCommunityEvent struct {
	ID          uint                `json:"id"`
	Role        string              `json:"role"`
	Logo        string              `json:"logo"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Count       int64               `json:"members"`
	Event       []ResponseListEvent `json:"event"`
}

type ResponseEventDetail struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Status        string    `json:"status"`
	Description   string    `json:"descriptions"`
	Penyelenggara string    `json:"penyelenggara"`
	Date          time.Time `json:"date_event"`
	Partisipasi   uint8     `json:"partisipasi"`
	Price         uint64    `json:"price"`
	Location      string    `json:"location"`
}

var layout = "2023-01-01 00:00:01"

func FromMidtransToPayment(resMidtrans *coreapi.ChargeResponse, payment_type string) Payment {

	var returnGopay []gopayResponse
	if resMidtrans.Actions != nil {
		for key := range resMidtrans.Actions {
			returnGopay = append(returnGopay, gopayResponse{
				Name:   resMidtrans.Actions[key].Name,
				Method: resMidtrans.Actions[key].Method,
				Url:    resMidtrans.Actions[key].URL,
			})
		}
	}

	if resMidtrans.VaNumbers == nil {
		resMidtrans.VaNumbers = append(resMidtrans.VaNumbers, coreapi.VANumber{
			Bank:     "",
			VANumber: "",
		})
	}

	return Payment{
		OrderID:           resMidtrans.OrderID,
		TransactionID:     resMidtrans.TransactionID,
		PaymentMethod:     resMidtrans.PaymentType,
		BillNumber:        resMidtrans.VaNumbers[0].VANumber,
		Bank:              resMidtrans.VaNumbers[0].Bank,
		GrossAmount:       resMidtrans.GrossAmount,
		TransactionTime:   resMidtrans.TransactionTime,
		TransactionStatus: resMidtrans.TransactionStatus,
		Actions:           returnGopay,
		Bill_key:          resMidtrans.BillKey,
		Biller_code:       resMidtrans.BillerCode,
	}
}

func ResponEventList(data []event.Response) []ResponseListEvent {

	var dataRespon []ResponseListEvent
	for _, v := range data {
		dataRespon = append(dataRespon, ResponseListEvent{
			ID:           v.ID,
			Logo:         v.Logo,
			Title:        v.Title,
			Descriptions: v.Descriptions,
			Date:         v.Date,
			Price:        v.Price,
			Members:      v.Members,
		})
	}

	return dataRespon
}

func ResponseEventListComu(dataComu event.CommunityEvent) ResponseCommunityEvent {
	return ResponseCommunityEvent{
		ID:          dataComu.ID,
		Role:        dataComu.Role,
		Logo:        dataComu.Logo,
		Title:       dataComu.Title,
		Description: dataComu.Description,
		Count:       dataComu.Count,
		Event:       ResponEventList(dataComu.Event),
	}
}

func ResponseEventDetails(data event.EventDetail) ResponseEventDetail {
	return ResponseEventDetail{
		ID:            data.ID,
		Title:         data.Title,
		Description:   data.Description,
		Status:        data.Status,
		Penyelenggara: data.Penyelenggara,
		Date:          data.Date,
		Price:         data.Price,
		Location:      data.Location,
	}
}
