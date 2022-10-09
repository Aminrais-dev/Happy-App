package delivery

import (
	"bytes"
	"capstone/happyApp/features/user"
	"capstone/happyApp/middlewares"
	"capstone/happyApp/mocks"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ResponseMassage struct {
	Status  string
	Message string
}

func TestPostData(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseUser)
	New(e, usecase)
	handlerTest := &userDelivery{
		userUsecase: usecase,
	}

	requestRegister := Request{
		Name:     "amin",
		Email:    "amin@gmail.com",
		Username: "amin rais",
		Gender:   "Male",
		Password: "amin",
	}

	requestBody, err := json.Marshal(requestRegister)
	if err != nil {
		t.Error(t, err, "error")
	}

	t.Run("Success Post", func(t *testing.T) {
		usecase.On("PostUser", mock.Anything, mock.Anything).Return(1, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/register")

		responseData := ResponseMassage{}

		if assert.NoError(t, handlerTest.CreateUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success sign up, please confirm email in gmail", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("failed Post", func(t *testing.T) {
		usecase.On("PostUser", mock.Anything, mock.Anything).Return(-2, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/register")

		responseData := ResponseMassage{}

		if assert.NoError(t, handlerTest.CreateUser(echoContext)) {
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

	t.Run("failed Post", func(t *testing.T) {
		usecase.On("PostUser", mock.Anything, mock.Anything).Return(-3, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/register")

		responseData := ResponseMassage{}

		if assert.NoError(t, handlerTest.CreateUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "email sudah terdaftar", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("failed Post", func(t *testing.T) {
		usecase.On("PostUser", mock.Anything, mock.Anything).Return(-4, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/register")

		responseData := ResponseMassage{}

		if assert.NoError(t, handlerTest.CreateUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "username sudah ada", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("failed Post", func(t *testing.T) {
		usecase.On("PostUser", mock.Anything, mock.Anything).Return(-1, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/register")

		responseData := ResponseMassage{}

		if assert.NoError(t, handlerTest.CreateUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed sign up", responseData.Message)
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

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestFail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/register")

		responseData := ResponseMassage{}

		if assert.NoError(t, handlerTest.CreateUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "error Bind", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

}

func TestDeleteData(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseUser)
	New(e, usecase)
	handlerTest := &userDelivery{
		userUsecase: usecase,
	}

	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}

	t.Run("Success delete data", func(t *testing.T) {
		usecase.On("DeleteUser", mock.AnythingOfType("int")).Return(1, nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/user/profile", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/user/profile")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.DeleteAccount))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success delete account", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed delete data", func(t *testing.T) {
		usecase.On("DeleteUser", mock.AnythingOfType("int")).Return(-1, nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/user/profile", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/user/profile")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.DeleteAccount))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed delete account", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

}

func TestUpdateData(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseUser)
	New(e, usecase)
	handlerTest := &userDelivery{
		userUsecase: usecase,
	}

	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}

	input := user.CoreUser{
		ID:       uint(1),
		Name:     "amin",
		Email:    "amin@gmail.com",
		Username: "amin rais",
		Gender:   "Male",
		Password: "amin",
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "amin")
	writer.WriteField("email", "amin@gmail.com")
	writer.WriteField("username", "amin rais")
	writer.WriteField("gender", "male")
	writer.WriteField("password", "amin")
	writer.Close()

	t.Run("Success update data", func(t *testing.T) {
		usecase.On("UpdateUser", input).Return(1).Once()

		req := httptest.NewRequest(http.MethodPut, "/user/profile", body)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/user/profile")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.UpdateAccount))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success update account", responseData.Message)
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

		req := httptest.NewRequest(http.MethodPut, "/user/profile", bytes.NewBuffer(requestFail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/user/profile")

		responseData := ResponseMassage{}

		if assert.NoError(t, handlerTest.UpdateAccount(echoContext)) {
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

	t.Run("failed update data", func(t *testing.T) {
		usecase.On("UpdateUser", input).Return(-4).Once()

		req := httptest.NewRequest(http.MethodPut, "/user/profile", body)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/user/profile")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.UpdateAccount))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "username sudah ada", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("failed update data", func(t *testing.T) {
		usecase.On("UpdateUser", input).Return(-1).Once()

		req := httptest.NewRequest(http.MethodPut, "/user/profile", body)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/user/profile")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.UpdateAccount))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed update data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

}
