package delivery

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/community"
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

type ResponsePesan struct {
	Message string
}

type ResponseList struct {
	Messege string
	Data    []Respose
}

var (
	mock_commun = Request{
		Title:        "Genshin",
		Descriptions: "Pemain Genshin",
		Logo:         config.DEFAULT_COMMUNITY,
	}
	mock_feed = FeedRequst{
		Text: "Ini Adalah Feed",
	}
	mock_comment = CommentRequst{
		Text: "Ini Adalah Comment",
	}
)

// func TestAddCommunity(t *testing.T) {
// 	reqBody, err := json.Marshal(mock_commun)
// 	if err != nil {
// 		t.Error(t, err, "error")
// 	}
// 	UsecaseMock := new(mocks.UsecaseInterface)
// 	e := echo.New()
// 	hand := &Delivery{
// 		From: UsecaseMock,
// 	}
// 	token, errToken := middlewares.CreateToken(1)
// 	if errToken != nil {
// 		assert.Error(t, errToken)
// 	}
// 	fmt.Println(token)
// 	t.Run("Success", func(t *testing.T) {

// 		UsecaseMock.On("AddNewCommunity", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return("Sebuah Pesan", nil).Once()

// 		req := httptest.NewRequest(http.MethodPost, "/community", bytes.NewBuffer(reqBody))
// 		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		echoContext := e.NewContext(req, rec)
// 		echoContext.SetPath("/community")
// 		response := ResponsePesan{}
// 		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.AddCommunity))(echoContext)
// 		fmt.Println("2")
// 		// if assert.NoError(t, srv.AddUser(echoContext)) {
// 		if assert.NoError(t, callFunc) {
// 			responseBody := rec.Body.String()
// 			err := json.Unmarshal([]byte(responseBody), &response)
// 			// fmt.Println("res", responseBody)
// 			if err != nil {
// 				assert.Error(t, err, "error")
// 			}
// 			assert.Equal(t, http.StatusOK, rec.Code)
// 			// assert.Equal(t, returnData[0].Name, responseData.Data[0].Name)
// 		}
// 		UsecaseMock.AssertExpectations(t)
// 	})

// }

func TestListCommunity(t *testing.T) {
	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}
	UsecaseMock := new(mocks.UsecaseInterface)
	e := echo.New()
	hand := &Delivery{
		From: UsecaseMock,
	}
	Res := []Respose{{ID: 1, Title: "Genshin", Descriptions: "Pemain Genshin", Logo: config.DEFAULT_COMMUNITY, Members: 10}}
	Com := []community.CoreCommunity{{ID: 1, Title: "Genshin", Descriptions: "Pemain Genshin", Logo: config.DEFAULT_COMMUNITY, Members: 10}}

	t.Run("Success", func(t *testing.T) {
		UsecaseMock.On("GetListCommunity").Return(Com, "Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/community", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community")
		response := ResponseList{}
		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.ListCommunity))(echoContext)

		// if assert.NoError(t, srv.AddUser(echoContext)) {
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &response)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, Res[0].Title, response.Data[0].Title)
		}
		UsecaseMock.AssertExpectations(t)
	})
	t.Run("Success 2", func(t *testing.T) {
		UsecaseMock.On("GetListCommunityWithParam", mock.Anything).Return(Com, "Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/community?title=shin", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community?title=shin")
		response := ResponseList{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.ListCommunity))(echoContext)

		// if assert.NoError(t, srv.AddUser(echoContext)) {
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &response)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, Res[0].Title, response.Data[0].Title)
		}
		UsecaseMock.AssertExpectations(t)
	})
	t.Run("Failed 1", func(t *testing.T) {
		UsecaseMock.On("GetListCommunityWithParam", mock.Anything).Return(Com, "Pesan", errors.New("Error")).Once()
		req := httptest.NewRequest(http.MethodGet, "/community?title=shin", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community?title=shin")
		response := ResponsePesan{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.ListCommunity))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Pesan", response.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
	t.Run("Failed 2", func(t *testing.T) {
		UsecaseMock.On("GetListCommunityWithParam", mock.Anything).Return([]community.CoreCommunity{}, "Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/community?title=shin", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community?title=shin")
		fmt.Println(echoContext)
		response := ResponsePesan{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.ListCommunity))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "Title tidak ditemukan", response.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
	t.Run("Failed 3", func(t *testing.T) {
		UsecaseMock.On("GetListCommunity").Return(Com, "Pesan", errors.New("Error")).Once()
		req := httptest.NewRequest(http.MethodGet, "/community", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community")
		response := ResponsePesan{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.ListCommunity))(echoContext)

		// if assert.NoError(t, srv.AddUser(echoContext)) {
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &response)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Pesan", response.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
}

// func TestListMember(t *testing.T) {

// }
