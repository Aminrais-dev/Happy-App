package delivery

import (
	"bytes"
	"capstone/happyApp/mocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Response struct {
	Access_Token string
	Message      string
	Status       string
}

func TestGetAll(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseLogin)
	New(e, usecase)
	handlerTest := &loginDelivery{
		loginUsecase: usecase,
	}

	requestLogin := Request{
		Email:    "amin@gmail.com",
		Password: "amin",
	}

	requestBody, err := json.Marshal(requestLogin)
	if err != nil {
		t.Error(t, err, "error")
	}

	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjUxMzg0MDksInVzZXJJZCI6MjJ9.sfr5TapBFdTUmqU3oUddS2PSH9m6fkfNU13"
	t.Run("Success Login", func(t *testing.T) {
		usecase.On("LoginAuthorized", mock.Anything, mock.Anything).Return(str, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/login")

		responseData := Response{}

		if assert.NoError(t, handlerTest.LoginUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, str, responseData.Access_Token)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("failed error bind", func(t *testing.T) {

		dataFail := map[string]int{
			"Name":     7777,
			"Password": 7777,
		}

		requestFail, err := json.Marshal(dataFail)
		if err != nil {
			t.Error(t, err, "error")
		}

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestFail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/login")

		responseData := Response{}

		if assert.NoError(t, handlerTest.LoginUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "wrong request", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed created token", func(t *testing.T) {
		usecase.On("LoginAuthorized", mock.Anything, mock.Anything).Return("", errors.New("error")).Once()

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/login")

		responseData := Response{}

		if assert.NoError(t, handlerTest.LoginUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed email not found", func(t *testing.T) {
		usecase.On("LoginAuthorized", mock.Anything, mock.Anything).Return("email not found", nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/login")

		responseData := Response{}

		if assert.NoError(t, handlerTest.LoginUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "email not found", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed wrong password", func(t *testing.T) {
		usecase.On("LoginAuthorized", mock.Anything, mock.Anything).Return("wrong password", nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/login")

		responseData := Response{}

		if assert.NoError(t, handlerTest.LoginUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "wrong password", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed request not all filled", func(t *testing.T) {
		usecase.On("LoginAuthorized", mock.Anything, mock.Anything).Return("please input email and password", nil).Once()

		requestLoginNull := Request{
			Email:    "amin@gmail.com",
			Password: "",
		}

		reqBodyNotPassword, err := json.Marshal(requestLoginNull)
		if err != nil {
			t.Error(t, err, "error")
		}

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBodyNotPassword))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/login")

		responseData := Response{}

		if assert.NoError(t, handlerTest.LoginUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "please input email and password", responseData.Message)
		}
		usecase.AssertExpectations(t)

	})

	t.Run("Failed status not confirm", func(t *testing.T) {
		usecase.On("LoginAuthorized", mock.Anything, mock.Anything).Return("please confirm your account in gmail", nil).Once()

		requestLoginNull := Request{
			Email:    "amin@gmail.com",
			Password: "",
		}

		reqBodyNotPassword, err := json.Marshal(requestLoginNull)
		if err != nil {
			t.Error(t, err, "error")
		}

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBodyNotPassword))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/login")

		responseData := Response{}

		if assert.NoError(t, handlerTest.LoginUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "please confirm your account in gmail", responseData.Message)
		}
		usecase.AssertExpectations(t)

	})

}
