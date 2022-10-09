package delivery

import (
	"bytes"
	"capstone/happyApp/features/user"
	"capstone/happyApp/middlewares"
	"capstone/happyApp/mocks"
	"encoding/json"
	"errors"
	"fmt"
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

type ResponseData struct {
	Data    Response
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

// func TestUpdateData(t *testing.T) {
// 	e := echo.New()
// 	usecase := new(mocks.UsecaseUser)
// 	New(e, usecase)
// 	handlerTest := &userDelivery{
// 		userUsecase: usecase,
// 	}

// 	token, errToken := middlewares.CreateToken(1)
// 	if errToken != nil {
// 		assert.Error(t, errToken)
// 	}

// 	input := user.CoreUser{
// 		ID:       uint(1),
// 		Name:     "amin",
// 		Email:    "amin@gmail.com",
// 		Username: "amin rais",
// 		Gender:   "Male",
// 		Password: "amin",
// 	}

// 	body := new(bytes.Buffer)
// 	writer := multipart.NewWriter(body)
// 	writer.WriteField("name", "amin")
// 	writer.WriteField("email", "amin@gmail.com")
// 	writer.WriteField("username", "amin rais")
// 	writer.WriteField("gender", "male")
// 	writer.WriteField("password", "amin")
// 	writer.Close()

// 	t.Run("Success update data", func(t *testing.T) {
// 		usecase.On("UpdateUser", input).Return(1).Once()

// 		req := httptest.NewRequest(http.MethodPut, "/user/profile", body)
// 		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
// 		rec := httptest.NewRecorder()
// 		echoContext := e.NewContext(req, rec)
// 		echoContext.SetPath("/user/profile")

// 		responseData := ResponseMassage{}

// 		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.UpdateAccount))(echoContext)
// 		if assert.NoError(t, callFunc) {
// 			responseBody := rec.Body.String()
// 			err := json.Unmarshal([]byte(responseBody), &responseData)
// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 			assert.Equal(t, http.StatusOK, rec.Code)
// 			assert.Equal(t, "success update account", responseData.Message)
// 		}
// 		usecase.AssertExpectations(t)
// 	})

// 	t.Run("failed error bind", func(t *testing.T) {

// 		dataFail := map[string]int{
// 			"Name":     7777,
// 			"Password": 7777,
// 		}

// 		requestFail, err := json.Marshal(dataFail)
// 		if err != nil {
// 			t.Error(t, err, "error")
// 		}

// 		req := httptest.NewRequest(http.MethodPut, "/user/profile", bytes.NewBuffer(requestFail))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
// 		rec := httptest.NewRecorder()
// 		echoContext := e.NewContext(req, rec)
// 		echoContext.SetPath("/user/profile")

// 		responseData := ResponseMassage{}

// 		if assert.NoError(t, handlerTest.UpdateAccount(echoContext)) {
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

// 	t.Run("failed update data", func(t *testing.T) {
// 		usecase.On("UpdateUser", input).Return(-4).Once()

// 		req := httptest.NewRequest(http.MethodPut, "/user/profile", body)
// 		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
// 		rec := httptest.NewRecorder()
// 		echoContext := e.NewContext(req, rec)
// 		echoContext.SetPath("/user/profile")

// 		responseData := ResponseMassage{}

// 		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.UpdateAccount))(echoContext)
// 		if assert.NoError(t, callFunc) {
// 			responseBody := rec.Body.String()
// 			err := json.Unmarshal([]byte(responseBody), &responseData)
// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 			assert.Equal(t, "username sudah ada", responseData.Message)
// 		}
// 		usecase.AssertExpectations(t)
// 	})

// 	t.Run("failed update data", func(t *testing.T) {
// 		usecase.On("UpdateUser", input).Return(-1).Once()

// 		req := httptest.NewRequest(http.MethodPut, "/user/profile", body)
// 		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
// 		rec := httptest.NewRecorder()
// 		echoContext := e.NewContext(req, rec)
// 		echoContext.SetPath("/user/profile")

// 		responseData := ResponseMassage{}

// 		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.UpdateAccount))(echoContext)
// 		if assert.NoError(t, callFunc) {
// 			responseBody := rec.Body.String()
// 			err := json.Unmarshal([]byte(responseBody), &responseData)
// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 			assert.Equal(t, "failed update data", responseData.Message)
// 		}
// 		usecase.AssertExpectations(t)
// 	})

// }

func TestGetData(t *testing.T) {
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

	var returnData2 []user.CommunityProfile
	returnData2 = append(returnData2, user.CommunityProfile{ID: 1, Title: "community saya", Logo: "https://logo", Role: "member"})
	returnData := user.CoreUser{ID: 1, Name: "amin", Username: "aminrais89", Gender: "Male", Email: "amin@a.com", Password: "$2a$10$3qSIi7BiTknraN3A9tRX/eoI4N9yuln/oWI8Ft9KcrZNF3ec6jIHK", Photo: "https://photo_profile"}

	t.Run("Success get data", func(t *testing.T) {
		usecase.On("GetUser", mock.Anything).Return(returnData, returnData2, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/user/profile", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/user/profile")

		responseData := ResponseData{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetProfile))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success get profile", responseData.Message)
			assert.Equal(t, returnData.Name, responseData.Data.Name)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed get data", func(t *testing.T) {
		usecase.On("GetUser", mock.Anything).Return(user.CoreUser{}, nil, errors.New("error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/user/profile", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/user/profile")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetProfile))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed get profile", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

}

func TestGetGmailVerify(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseUser)
	New(e, usecase)
	handlerTest := &userDelivery{
		userUsecase: usecase,
	}

	t.Run("Success confirm email", func(t *testing.T) {
		usecase.On("UpdateStatus", mock.Anything).Return(1).Once()

		req := httptest.NewRequest(http.MethodGet, "/user/verify", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/user/verify")

		responseData := ResponseData{}

		if assert.NoError(t, handlerTest.GmailVerify(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success confirm your account", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed confirm email", func(t *testing.T) {
		usecase.On("UpdateStatus", mock.Anything).Return(-1).Once()

		req := httptest.NewRequest(http.MethodGet, "/user/verify", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/user/verify")

		responseData := ResponseMassage{}

		if assert.NoError(t, handlerTest.GmailVerify(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed to verify account", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

}
