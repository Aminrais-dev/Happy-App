package delivery

import (
	"bytes"
	"capstone/happyApp/features/event"
	"capstone/happyApp/middlewares"
	"capstone/happyApp/mocks"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ResponseMassage struct {
	Status  string
	Message string
}

type ResponseList struct {
	Data    []ResponseListEvent
	Message string
}

func TestPostEvent(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseEvent)
	New(e, usecase)
	handlerTest := &eventDelivery{
		eventUsecase: usecase,
	}

	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}

	requestRegister := Request{Title: "event keren", Description: "baju kain woll terbaik", Location: "JAPAN", Date: "2022-11-09T15:30", Price: 30000}

	requestBody, err := json.Marshal(requestRegister)
	if err != nil {
		t.Error(t, err, "error")
	}

	t.Run("Success post event", func(t *testing.T) {
		usecase.On("PostEvent", mock.Anything, mock.Anything).Return(1).Once()

		req := httptest.NewRequest(http.MethodPost, "/community/:id/event", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/event")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.PostEventCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success create event", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Post", func(t *testing.T) {
		usecase.On("PostEvent", mock.Anything, mock.Anything).Return(-2).Once()

		req := httptest.NewRequest(http.MethodPost, "/community/:id/event", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/event")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.PostEventCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "not have access in community", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Post", func(t *testing.T) {
		usecase.On("PostEvent", mock.Anything, mock.Anything).Return(-1).Once()

		req := httptest.NewRequest(http.MethodPost, "/community/:id/event", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/event")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.PostEventCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed to create event", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Post", func(t *testing.T) {
		usecase.On("PostEvent", mock.Anything, mock.Anything).Return(-3).Once()

		req := httptest.NewRequest(http.MethodPost, "/community/:id/event", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/event")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.PostEventCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "please input all request", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed param", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/community/:id/event", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/event")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("abs")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.PostEventCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "param must be number", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("failed error bind", func(t *testing.T) {

		dataFail := map[string]int{
			"Title": 7777,
			"Date":  7777,
		}

		requestFail, err := json.Marshal(dataFail)
		if err != nil {
			t.Error(t, err, "error")
		}

		req := httptest.NewRequest(http.MethodPost, "/community/:id/event", bytes.NewBuffer(requestFail))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/event")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.PostEventCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "error bind", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

}

func TestGetEventList(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseEvent)
	New(e, usecase)
	handlerTest := &eventDelivery{
		eventUsecase: usecase,
	}

	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}

	var returnData []event.Response
	returnData = append(returnData, event.Response{ID: 1, Title: "event keren", Members: 0, Logo: "https://logo", Descriptions: "event untuk jadi keren", Date: time.Now(), Price: 200000})

	t.Run("Success get event", func(t *testing.T) {
		usecase.On("GetEvent", mock.Anything, mock.Anything).Return(returnData, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/event", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/event")

		responseData := ResponseList{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetEventList))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success get list event", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed get event", func(t *testing.T) {
		usecase.On("GetEvent", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/event", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/event")

		responseData := ResponseList{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetEventList))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed to get list event", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

}

func TestGetEventListCommunity(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseEvent)
	New(e, usecase)
	handlerTest := &eventDelivery{
		eventUsecase: usecase,
	}

	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}

	var temp []event.Response
	temp = append(temp, event.Response{ID: 1, Title: "event keren", Members: 0, Logo: "https://logo", Descriptions: "event untuk jadi keren", Date: time.Now(), Price: 200000})
	var returnData = event.CommunityEvent{ID: 1, Title: "comunity keren", Role: "member", Logo: "https://logo", Description: "community untuk jadi keren", Count: 12, Event: temp}

	t.Run("Success get event", func(t *testing.T) {
		usecase.On("GetEventComu", mock.Anything, mock.Anything).Return(returnData, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/community/:id/event", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/event")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseCommunityEvent{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetEventListCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, ResponseEventListComu(returnData).ID, responseData.ID)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Faild community not found", func(t *testing.T) {
		usecase.On("GetEventComu", mock.Anything, mock.Anything).Return(event.CommunityEvent{}, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/community/:id/event", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/event")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetEventListCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "community not found", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Faild get event", func(t *testing.T) {
		usecase.On("GetEventComu", mock.Anything, mock.Anything).Return(event.CommunityEvent{}, errors.New("error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/community/:id/event", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/event")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetEventListCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed to get list event in community", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed param", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/community/:id/event", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/event")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("abs")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetEventListCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "param must be number", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}

func TestGetEventDetail(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseEvent)
	New(e, usecase)
	handlerTest := &eventDelivery{
		eventUsecase: usecase,
	}

	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}

	returnData := event.EventDetail{ID: 1, Title: "event keren", Status: "join", Description: "event jadi keren", Penyelenggara: "community keren", Date: time.Now(), Partisipasi: 23, Price: 20000, Location: "JAPAN"}

	t.Run("Success get event", func(t *testing.T) {
		usecase.On("GetEventDetail", mock.Anything, mock.Anything).Return(returnData, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/event/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/event/:id")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseCommunityEvent{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetEventDetailbyId))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, ResponseEventDetails(returnData).ID, responseData.ID)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Faild event not found", func(t *testing.T) {
		usecase.On("GetEventDetail", mock.Anything, mock.Anything).Return(event.EventDetail{}, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/event/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/event/:id")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetEventDetailbyId))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "event not found", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Faild get event", func(t *testing.T) {
		usecase.On("GetEventDetail", mock.Anything, mock.Anything).Return(event.EventDetail{}, errors.New("error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/event/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/event/:id")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetEventDetailbyId))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed to get event detail", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed param", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/event/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/event/:id")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("abs")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetEventDetailbyId))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "param must be number", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}

// func TestCreatePayment(t *testing.T) {
// 	e := echo.New()
// 	usecase := new(mocks.UsecaseEvent)
// 	New(e, usecase)
// 	handlerTest := &eventDelivery{
// 		eventUsecase: usecase,
// 	}

// 	midtrans.ServerKey = config.MidtransServerKey()
// 	paymentEvent.New(midtrans.ServerKey, midtrans.Sandbox)

// 	token, errToken := middlewares.CreateToken(1)
// 	if errToken != nil {
// 		assert.Error(t, errToken)
// 	}

// 	request := RequestPayment{Payment_type: "BCA Virtual Account"}
// 	requestFail := RequestPayment{Payment_type: "BCA VIRTUAL ACCOUNT"}
// 	requestBody, err := json.Marshal(request)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}
// 	reqFail, errFail := json.Marshal(requestFail)
// 	if errFail != nil {
// 		t.Error(t, errFail, "error")
// 	}

// 	returnData := &coreapi.ChargeResponse{TransactionID: "12J1231", OrderID: "E-1231", GrossAmount: "20000", VaNumbers: []coreapi.VANumber{{VANumber: "1232131", Bank: "bca"}}}
// 	t.Run("Success create payment", func(t *testing.T) {
// 		usecase.On("GetAmountEvent", mock.Anything).Return(uint64(20000)).Once()
// 		usecase.On("CreatePaymentMidtrans", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(returnData, nil).Once()

// 		req := httptest.NewRequest(http.MethodPost, "/join/event/:id", bytes.NewBuffer(requestBody))
// 		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		echoContext := e.NewContext(req, rec)
// 		echoContext.SetPath("/join/event/:id")
// 		echoContext.SetParamNames("id")
// 		echoContext.SetParamValues("1")

// 		responseData := ResponseMassage{}

// 		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.CreatePaymentJoinEvent))(echoContext)
// 		if assert.NoError(t, callFunc) {
// 			responseBody := rec.Body.String()
// 			err := json.Unmarshal([]byte(responseBody), &responseData)
// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 			assert.Equal(t, http.StatusOK, rec.Code)
// 			assert.Equal(t, "success create transactions", responseData.Message)
// 		}
// 		usecase.AssertExpectations(t)
// 	})

// 	t.Run("failed create payment", func(t *testing.T) {
// 		usecase.On("GetAmountEvent", mock.Anything).Return(uint64(20000)).Once()
// 		usecase.On("CreatePaymentMidtrans", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&coreapi.ChargeResponse{}, errors.New("error")).Once()

// 		req := httptest.NewRequest(http.MethodPost, "/join/event/:id", bytes.NewBuffer(requestBody))
// 		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		echoContext := e.NewContext(req, rec)
// 		echoContext.SetPath("/join/event/:id")
// 		echoContext.SetParamNames("id")
// 		echoContext.SetParamValues("1")

// 		responseData := ResponseMassage{}

// 		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.CreatePaymentJoinEvent))(echoContext)
// 		if assert.NoError(t, callFunc) {
// 			responseBody := rec.Body.String()
// 			err := json.Unmarshal([]byte(responseBody), &responseData)
// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 			assert.Equal(t, "failed to create transaction", responseData.Message)
// 		}
// 		usecase.AssertExpectations(t)
// 	})

// 	t.Run("failed get amount payment", func(t *testing.T) {
// 		usecase.On("GetAmountEvent", mock.Anything).Return(uint64(00)).Once()

// 		req := httptest.NewRequest(http.MethodPost, "/join/event/:id", bytes.NewBuffer(requestBody))
// 		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		echoContext := e.NewContext(req, rec)
// 		echoContext.SetPath("/join/event/:id")
// 		echoContext.SetParamNames("id")
// 		echoContext.SetParamValues("1")

// 		responseData := ResponseMassage{}

// 		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.CreatePaymentJoinEvent))(echoContext)
// 		if assert.NoError(t, callFunc) {
// 			responseBody := rec.Body.String()
// 			err := json.Unmarshal([]byte(responseBody), &responseData)
// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 			assert.Equal(t, "failed to get gross amount", responseData.Message)
// 		}
// 		usecase.AssertExpectations(t)
// 	})

// 	t.Run("failed type payment", func(t *testing.T) {
// 		usecase.On("GetAmountEvent", mock.Anything).Return(uint64(20000)).Once()

// 		req := httptest.NewRequest(http.MethodPost, "/join/event/:id", bytes.NewBuffer(reqFail))
// 		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		echoContext := e.NewContext(req, rec)
// 		echoContext.SetPath("/join/event/:id")
// 		echoContext.SetParamNames("id")
// 		echoContext.SetParamValues("1")

// 		responseData := ResponseMassage{}

// 		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.CreatePaymentJoinEvent))(echoContext)
// 		if assert.NoError(t, callFunc) {
// 			responseBody := rec.Body.String()
// 			err := json.Unmarshal([]byte(responseBody), &responseData)
// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 			assert.Equal(t, "payment type not allowed", responseData.Message)
// 		}
// 		usecase.AssertExpectations(t)
// 	})

// 	t.Run("Failed param", func(t *testing.T) {

// 		req := httptest.NewRequest(http.MethodGet, "/join/event/:id", bytes.NewBuffer(requestBody))
// 		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		echoContext := e.NewContext(req, rec)
// 		echoContext.SetPath("/join/event/:id")
// 		echoContext.SetParamNames("id")
// 		echoContext.SetParamValues("abs")

// 		responseData := ResponseMassage{}

// 		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.CreatePaymentJoinEvent))(echoContext)
// 		if assert.NoError(t, callFunc) {
// 			responseBody := rec.Body.String()
// 			err := json.Unmarshal([]byte(responseBody), &responseData)
// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 			assert.Equal(t, "param must be number", responseData.Message)
// 		}
// 		usecase.AssertExpectations(t)
// 	})

// 	t.Run("failed error bind", func(t *testing.T) {

// 		dataFail := map[string]int{
// 			"Payment_type": 7777,
// 		}

// 		requestFail, err := json.Marshal(dataFail)
// 		if err != nil {
// 			t.Error(t, err, "error")
// 		}

// 		req := httptest.NewRequest(http.MethodPost, "/join/event/:id", bytes.NewBuffer(requestFail))
// 		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		echoContext := e.NewContext(req, rec)
// 		echoContext.SetPath("/join/event/:id")
// 		echoContext.SetParamNames("id")
// 		echoContext.SetParamValues("1")

// 		responseData := ResponseMassage{}

// 		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.CreatePaymentJoinEvent))(echoContext)
// 		if assert.NoError(t, callFunc) {
// 			responseBody := rec.Body.String()
// 			err := json.Unmarshal([]byte(responseBody), &responseData)
// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 			assert.Equal(t, "error bind", responseData.Message)
// 		}
// 		usecase.AssertExpectations(t)
// 	})

// }
