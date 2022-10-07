package usecase

import (
	"capstone/happyApp/features/event"

	"github.com/midtrans/midtrans-go/coreapi"
)

type usecaseEvent struct {
	eventData event.DataInterface
}

func New(data event.DataInterface) event.UsecaseInterface {
	return &usecaseEvent{
		eventData: data,
	}
}

func (usecase *usecaseEvent) PostEvent(data event.EventCore, id int) int {

	if data.Description == "" || data.Location == "" || data.Title == "" {
		return -3
	}

	row := usecase.eventData.InsertEvent(data, id)
	return row

}

func (usecase *usecaseEvent) GetEvent(search string) ([]event.Response, error) {

	data, err := usecase.eventData.SelectEvent(search)
	if err != nil {
		return nil, err
	}

	dataRes := usecase.eventData.GetMembers(data)

	return dataRes, nil
}

func (usecase *usecaseEvent) GetEventComu(search string, idComu, userId int) (event.CommunityEvent, error) {

	data, _ := usecase.eventData.SelectEvent(search)

	dataRes := usecase.eventData.GetMembers(data)

	dataReturn, err := usecase.eventData.SelectEventComu(dataRes, idComu, userId)
	if err != nil {
		return dataReturn, err
	}

	return dataReturn, nil
}

func (usecase *usecaseEvent) GetEventDetail(idEvent, userId int) (event.EventDetail, error) {

	data, err := usecase.eventData.SelectEventDetail(idEvent, userId)
	if err != nil {
		return data, err
	}

	return data, nil

}

func (usecase *usecaseEvent) GetAmountEvent(idEvent int) uint64 {

	data := usecase.eventData.SelectAmountEvent(idEvent)
	return data

}

func (usecase *usecaseEvent) CreatePaymentMidtrans(reqMidtrans coreapi.ChargeReq, userId, idEvent int, method string) (*coreapi.ChargeResponse, error) {

	chargeResponse, err := usecase.eventData.CreatePayment(reqMidtrans, userId, idEvent, method)
	if err != nil {
		return nil, err
	}

	return chargeResponse, nil
}
