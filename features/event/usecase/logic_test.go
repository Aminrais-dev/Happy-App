package usecase

import (
	"capstone/happyApp/features/event"
	"capstone/happyApp/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetEventComu(t *testing.T) {

	eventMock := new(mocks.DataEvent)
	var temp []event.Response
	temp = append(temp, event.Response{ID: 1, Title: "event keren", Members: 20, Logo: "https://logo", Descriptions: "event untuk jadi keren", Date: time.Now(), Price: 200000})
	var returnData = event.CommunityEvent{ID: 1, Title: "comunity keren", Role: "member", Logo: "https://logo", Description: "community untuk jadi keren", Count: 12, Event: temp}

	comuId := uint(1)
	userId := 1

	t.Run("Get event community success", func(t *testing.T) {
		eventMock.On("SelectEventComu", mock.Anything, mock.Anything, mock.Anything).Return(returnData, nil).Once()

		useCase := New(eventMock)
		res, err := useCase.GetEventComu("", int(comuId), userId)
		assert.Equal(t, res.ID, comuId)
		assert.Nil(t, err)
		eventMock.AssertExpectations(t)

	})

	t.Run("Get event community failed", func(t *testing.T) {

		eventMock.On("SelectEventComu", mock.Anything, mock.Anything, mock.Anything).Return(event.CommunityEvent{}, errors.New("error")).Once()

		useCase := New(eventMock)
		res, err := useCase.GetEventComu("", int(comuId), userId)
		assert.Error(t, err)
		assert.NotEqual(t, res.ID, comuId)
		assert.NotNil(t, err)
		eventMock.AssertExpectations(t)

	})

}

func TestGetEventId(t *testing.T) {

	eventMock := new(mocks.DataEvent)
	returnData := event.EventDetail{ID: 1, Title: "event keren", Status: "join", Description: "event jadi keren", Penyelenggara: "community keren", Date: time.Now(), Partisipasi: 23, Price: 20000, Location: "JAPAN"}
	eventId := uint(1)
	userId := 1

	t.Run("Get event success", func(t *testing.T) {
		eventMock.On("SelectEventDetail", int(eventId), userId).Return(returnData, nil).Once()

		useCase := New(eventMock)
		res, err := useCase.GetEventDetail(int(eventId), userId)
		assert.Equal(t, res.ID, eventId)
		assert.Nil(t, err)
		eventMock.AssertExpectations(t)

	})

	t.Run("Get event failed", func(t *testing.T) {

		eventMock.On("SelectEventDetail", int(eventId), userId).Return(event.EventDetail{}, errors.New("error")).Once()

		useCase := New(eventMock)
		res, err := useCase.GetEventDetail(int(eventId), userId)
		assert.Error(t, err)
		assert.NotEqual(t, res.ID, eventId)
		assert.NotNil(t, err)
		eventMock.AssertExpectations(t)

	})

}

func TestGetEvent(t *testing.T) {

	eventMock := new(mocks.DataEvent)
	var returnData []event.Response
	returnData = append(returnData, event.Response{ID: 1, Title: "event keren", Members: 20, Logo: "https://logo", Descriptions: "event untuk jadi keren", Date: time.Now(), Price: 200000})

	t.Run("Get event list success", func(t *testing.T) {
		eventMock.On("SelectEvent", mock.Anything).Return(returnData, nil).Once()

		useCase := New(eventMock)
		res, err := useCase.GetEvent("")
		assert.Equal(t, res[0].ID, returnData[0].ID)
		assert.Nil(t, err)
		eventMock.AssertExpectations(t)

	})

	t.Run("Get event list failed", func(t *testing.T) {

		eventMock.On("SelectEvent", mock.Anything).Return(nil, errors.New("error")).Once()

		useCase := New(eventMock)
		res, err := useCase.GetEvent("")
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		eventMock.AssertExpectations(t)

	})

}

func TestPostEvent(t *testing.T) {

	eventMock := new(mocks.DataEvent)
	input := event.EventCore{Title: "event keren", Description: "baju kain woll terbaik", Location: "JAPAN", Date: time.Now(), Price: 30000}
	userId := 1

	t.Run("create success", func(t *testing.T) {

		eventMock.On("InsertEvent", mock.Anything, mock.Anything).Return(1).Once()

		useCase := New(eventMock)
		res := useCase.PostEvent(input, userId)
		assert.Equal(t, 1, res)
		eventMock.AssertExpectations(t)

	})

	t.Run("create failed", func(t *testing.T) {

		eventMock.On("InsertEvent", mock.Anything, mock.Anything).Return(-1).Once()

		useCase := New(eventMock)
		res := useCase.PostEvent(input, userId)
		assert.Equal(t, -1, res)
		eventMock.AssertExpectations(t)

	})

	t.Run("failed not have access", func(t *testing.T) {

		eventMock.On("InsertEvent", mock.Anything, mock.Anything).Return(-2).Once()

		useCase := New(eventMock)
		res := useCase.PostEvent(input, userId)
		assert.Equal(t, -2, res)
		eventMock.AssertExpectations(t)

	})

	t.Run("create failed because input title not filled", func(t *testing.T) {

		input.Title = ""
		useCase := New(eventMock)
		res := useCase.PostEvent(input, userId)
		assert.Equal(t, -3, res)
		eventMock.AssertExpectations(t)

	})

}

func TestGetAmount(t *testing.T) {

	eventMock := new(mocks.DataEvent)
	eventId := 1

	t.Run("Get amount succes", func(t *testing.T) {

		eventMock.On("SelectAmountEvent", mock.Anything).Return(uint64(200000)).Once()

		useCase := New(eventMock)
		res := useCase.GetAmountEvent(eventId)
		assert.NotEqual(t, uint64(00), res)
		eventMock.AssertExpectations(t)

	})

	t.Run("Get amount failed", func(t *testing.T) {

		eventMock.On("SelectAmountEvent", mock.Anything).Return(uint64(00)).Once()

		useCase := New(eventMock)
		res := useCase.GetAmountEvent(eventId)
		assert.Equal(t, uint64(00), res)
		eventMock.AssertExpectations(t)

	})

}

// func TestCreatePayment(t *testing.T) {

// 	eventMock := new(mocks.DataEvent)
// 	userId := 1
// 	eventId := 1
// 	method := "GOPAY"
// 	orderId := "event-131231"
// 	req := coreapi.ChargeReq{
// 		PaymentType: "gopay",
// 		TransactionDetails: midtrans.TransactionDetails{
// 			OrderID:  orderId,
// 			GrossAmt: 800000,
// 		},
// 	}

// 	dataGopayAction := coreapi.Action{Name: "gr-code", URL: "https//"}
// 	var dataGopay []coreapi.Action
// 	dataGopay = append(dataGopay, dataGopayAction)

// 	returnData := coreapi.ChargeResponse{
// 		TransactionStatus: "pending",
// 		OrderID:           orderId,
// 		Actions:           dataGopay,
// 	}

// 	t.Run("Create payment succes", func(t *testing.T) {

// 		eventMock.On("CreatePayment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(returnData, nil).Once()

// 		useCase := New(eventMock)
// 		res, err := useCase.CreatePaymentMidtrans(req, userId, eventId, method)
// 		assert.Equal(t, returnData, res)
// 		assert.Nil(t, err)
// 		eventMock.AssertExpectations(t)

// 	})

// 	t.Run("Create payment failed", func(t *testing.T) {

// 		eventMock.On("CreatePayment", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(00).Once()

// 		useCase := New(eventMock)
// 		res, err := useCase.CreatePaymentMidtrans(req, userId, eventId, method)
// 		assert.Error(t, err)
// 		assert.Equal(t, res[0].ID, returnData[0].ID)
// 		assert.Nil(t, err)
// 		eventMock.AssertExpectations(t)

// 	})

// }
