package delivery

import (
	"bytes"
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

type UnivResponse struct {
	Message string
	Data    []interface{}
}

type ResponseList struct {
	Messege string
	Data    []Respose
}
type Response struct {
	Messege string
	Data    Respose
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

func TestListMember(t *testing.T) {
	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}
	UsecaseMock := new(mocks.UsecaseInterface)
	e := echo.New()
	hand := &Delivery{
		From: UsecaseMock,
	}

	names := []string{"Muhammad", "Nawi", "Husen"}

	t.Run("Success Get All", func(t *testing.T) {
		UsecaseMock.On("GetMembers", mock.Anything).Return(names, "Sebuah Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/members/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/members/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("1")

		responseData := UnivResponse{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.ListMembersCommunity))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "Nawi", responseData.Data[1])
		}
		UsecaseMock.AssertExpectations(t)
	})

	t.Run("Failed 1", func(t *testing.T) {
		// UsecaseMock.On("GetMembers", mock.Anything).Return(names, "Sebuah Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/members/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/members/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("")

		responseData := ResponsePesan{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.ListMembersCommunity))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Parameter must be number", responseData.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
	t.Run("Failed 2", func(t *testing.T) {
		UsecaseMock.On("GetMembers", mock.Anything).Return(names, "Sebuah Pesan", errors.New("Error")).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/members/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/members/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("1")

		responseData := UnivResponse{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.ListMembersCommunity))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Sebuah Pesan", responseData.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
}

func TestOutFromCommunity(t *testing.T) {
	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}
	UsecaseMock := new(mocks.UsecaseInterface)
	e := echo.New()
	hand := &Delivery{
		From: UsecaseMock,
	}

	t.Run("Success", func(t *testing.T) {
		UsecaseMock.On("Leave", mock.Anything, mock.Anything).Return("Sebuah Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("1")

		responseData := ResponsePesan{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.OutFromCommunity))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "Sebuah Pesan", responseData.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
	t.Run("Failed 1", func(t *testing.T) {
		// UsecaseMock.On("Leave", mock.Anything, mock.Anything).Return("Sebuah Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("")

		responseData := ResponsePesan{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.OutFromCommunity))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Parameter must be number", responseData.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
	t.Run("Failed 2", func(t *testing.T) {
		UsecaseMock.On("Leave", mock.Anything, mock.Anything).Return("Sebuah Pesan", errors.New("Errors")).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("2")

		responseData := ResponsePesan{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.OutFromCommunity))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Sebuah Pesan", responseData.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
}

func TestJoinCommunity(t *testing.T) {
	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}
	UsecaseMock := new(mocks.UsecaseInterface)
	e := echo.New()
	hand := &Delivery{
		From: UsecaseMock,
	}

	t.Run("Success", func(t *testing.T) {
		UsecaseMock.On("JoinCommunity", mock.Anything, mock.Anything).Return("Sebuah Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/join/community/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/join/community/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("1")

		responseData := ResponsePesan{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.JoinCommunity))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "Sebuah Pesan", responseData.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
	t.Run("Failed 1", func(t *testing.T) {
		// UsecaseMock.On("Leave", mock.Anything, mock.Anything).Return("Sebuah Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/join/community/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/join/community/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("")

		responseData := ResponsePesan{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.JoinCommunity))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Parameter must be number", responseData.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
	t.Run("Failed 2", func(t *testing.T) {
		UsecaseMock.On("JoinCommunity", mock.Anything, mock.Anything).Return("Sebuah Pesan", errors.New("Errors")).Once()
		req := httptest.NewRequest(http.MethodGet, "/join/community/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/join/community/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("2")

		responseData := ResponsePesan{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.JoinCommunity))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Sebuah Pesan", responseData.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
}
func TestDeatailCommunity(t *testing.T) {
	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}
	UsecaseMock := new(mocks.UsecaseInterface)
	e := echo.New()
	hand := &Delivery{
		From: UsecaseMock,
	}
	Com := community.CoreCommunity{ID: 1, Title: "Genshin", Descriptions: "Pemain Genshin", Logo: config.DEFAULT_COMMUNITY, Members: 10}

	t.Run("Failed 1", func(t *testing.T) {
		UsecaseMock.On("GetCommunityFeed", mock.Anything, mock.Anything).Return(Com, "Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("2")
		response := Response{}
		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.GetDetailCommunity))(echoContext)

		// if assert.NoError(t, srv.AddUser(echoContext)) {
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &response)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "Genshin", response.Data.Title)
		}
		UsecaseMock.AssertExpectations(t)
	})

	t.Run("Failed 1", func(t *testing.T) {
		// UsecaseMock.On("GetCommunityFeed", mock.Anything, mock.Anything).Return(Com, "Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("")
		response := ResponsePesan{}
		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.GetDetailCommunity))(echoContext)

		// if assert.NoError(t, srv.AddUser(echoContext)) {
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &response)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Parameter must be number", response.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
	t.Run("Failed 2", func(t *testing.T) {
		UsecaseMock.On("GetCommunityFeed", mock.Anything, mock.Anything).Return(Com, "Pesan", errors.New("Errors")).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/:communityid", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:communityid")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("1")
		response := ResponsePesan{}
		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.GetDetailCommunity))(echoContext)

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

func TestAddFeed(t *testing.T) {
	reqBody, err := json.Marshal(mock_comment)
	if err != nil {
		t.Error(t, err, "error")
	}
	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}
	UsecaseMock := new(mocks.UsecaseInterface)
	e := echo.New()
	hand := &Delivery{
		From: UsecaseMock,
	}

	t.Run("Success", func(t *testing.T) {
		UsecaseMock.On("PostFeed", mock.Anything).Return("Pesan", nil).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/:communityid/feed", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:communityid/feed")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("1")
		response := ResponsePesan{}
		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.AddFeed))(echoContext)

		// if assert.NoError(t, srv.AddUser(echoContext)) {
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &response)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "Pesan", response.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})
	t.Run("Failed 1", func(t *testing.T) {
		UsecaseMock.On("PostFeed", mock.Anything).Return("Pesan", errors.New("error")).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/:communityid/feed", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:communityid/feed")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("1")
		response := ResponsePesan{}
		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.AddFeed))(echoContext)

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
	t.Run("Failed 2", func(t *testing.T) {
		// UsecaseMock.On("PostFeed", mock.Anything).Return("Pesan", errors.New("error")).Once()
		req := httptest.NewRequest(http.MethodGet, "/community/:communityid/feed", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:communityid/feed")
		echoContext.SetParamNames("communityid")
		echoContext.SetParamValues("")
		response := ResponsePesan{}
		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.AddFeed))(echoContext)

		// if assert.NoError(t, srv.AddUser(echoContext)) {
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &response)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Parameter must be number", response.Message)
		}
		UsecaseMock.AssertExpectations(t)
	})

	// t.Run("Failed 3", func(t *testing.T) {
	// 	// UsecaseMock.On("PostFeed", mock.Anything).Return("Pesan", errors.New("error")).Once()
	// 	req := httptest.NewRequest(http.MethodGet, "/community/:communityid/feed", bytes.NewBuffer(reqBody))
	// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// 	rec := httptest.NewRecorder()
	// 	echoContext := e.NewContext(req, rec)
	// 	echoContext.SetPath("/community/:communityid/feed")
	// 	echoContext.SetParamNames("communityid")
	// 	echoContext.SetParamValues("1")
	// 	response := ResponsePesan{}
	// 	callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(hand.AddFeed))(echoContext)

	// 	// if assert.NoError(t, srv.AddUser(echoContext)) {
	// 	if assert.NoError(t, callFunc) {
	// 		responseBody := rec.Body.String()
	// 		err := json.Unmarshal([]byte(responseBody), &response)
	// 		// fmt.Println("res", responseBody)
	// 		if err != nil {
	// 			assert.Error(t, err, "error")
	// 		}
	// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
	// 		// assert.Equal(t, "Parameter must be number", response.Message)
	// 	}
	// 	UsecaseMock.AssertExpectations(t)
	// })
}
