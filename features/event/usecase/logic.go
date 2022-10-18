package usecase

import (
	"capstone/happyApp/features/event"
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

func (usecase *usecaseEvent) GetEventComu(idComu, userId int) (event.CommunityEvent, error) {

	dataReturn, err := usecase.eventData.SelectEventComu(idComu, userId)
	if err != nil {
		return dataReturn, err
	}

	dataRes := usecase.eventData.GetMembers(dataReturn.Event)
	dataReturn.Event = dataRes

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

func (usecase *usecaseEvent) PostTransaction(data event.JoinEventCore) error {

	err := usecase.eventData.InsertTransaction(data)
	return err

}

func (usecase *usecaseEvent) CheckStatus(userId, idEvent int) error {

	err := usecase.eventData.CheckJoin(userId, idEvent)
	return err

}
