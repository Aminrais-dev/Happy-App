package delivery

import (
	"capstone/happyApp/features/product"
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
	Message string
	Status  string
}

func TestGetListProductComu(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseProduct)
	New(e, usecase)
	handlerTest := &productDelivery{
		productUsecase: usecase,
	}

	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}

	var returnData2 []product.ProductCore
	returnData2 = append(returnData2, product.ProductCore{ID: 1, Name: "baju baru", Description: "baju kain woll terbaik", Photo: "https://photo", Stock: 10, Price: 30000, CommunityID: 1})
	returnData := product.Comu{ID: 1, Title: "comunity keren", Role: "member", Logo: "https://logo", Description: "community untuk jadi keren", Count: 12}

	t.Run("Success get product", func(t *testing.T) {
		usecase.On("GetProductComu", mock.Anything, mock.Anything).Return(returnData, returnData2, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/community/:id/store", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/store")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := Response{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetListProductCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed get product", func(t *testing.T) {
		usecase.On("GetProductComu", mock.Anything, mock.Anything).Return(product.Comu{}, nil, errors.New("error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/community/:id/store", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/store")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetListProductCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Failed get product detail", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed param", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/community/:id/store", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/store")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("abs")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetListProductCommunity))(echoContext)
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

func TestGetProductDetail(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseProduct)
	New(e, usecase)
	handlerTest := &productDelivery{
		productUsecase: usecase,
	}

	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}

	returnData := product.Comu{ID: 1, Title: "comunity keren", Role: "member", Logo: "https://logo", Description: "community untuk jadi keren", Count: 12}
	returnData2 := product.ProductCore{ID: 1, Name: "baju baru", Description: "baju kain woll terbaik", Photo: "https://photo", Stock: 10, Price: 30000, CommunityID: 1}

	t.Run("Success get product", func(t *testing.T) {
		usecase.On("GetProduct", mock.Anything, mock.Anything).Return(returnData, returnData2, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/community/:id/store", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/store")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := Response{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetProductDetail))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed get product", func(t *testing.T) {
		usecase.On("GetProduct", mock.Anything, mock.Anything).Return(product.Comu{}, product.ProductCore{}, errors.New("error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/community/:id/store", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/store")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetProductDetail))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Failed get product detail", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed get product", func(t *testing.T) {
		usecase.On("GetProduct", mock.Anything, mock.Anything).Return(product.Comu{}, product.ProductCore{}, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/community/:id/store", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/store")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetProductDetail))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "product not found", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed param", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/community/:id/store", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/community/:id/store")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("abs")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.GetProductDetail))(echoContext)
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

func TestDeleteProduct(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UsecaseProduct)
	New(e, usecase)
	handlerTest := &productDelivery{
		productUsecase: usecase,
	}

	token, errToken := middlewares.CreateToken(1)
	if errToken != nil {
		assert.Error(t, errToken)
	}

	t.Run("Success delete product", func(t *testing.T) {
		usecase.On("DeleteProduct", mock.Anything, mock.Anything).Return(1).Once()

		req := httptest.NewRequest(http.MethodDelete, "/store/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/store/:id")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.DeleteProductCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success delete product", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed delete product", func(t *testing.T) {
		usecase.On("DeleteProduct", mock.Anything, mock.Anything).Return(-1).Once()

		req := httptest.NewRequest(http.MethodDelete, "/store/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/store/:id")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.DeleteProductCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "Failed delete product", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed delete product", func(t *testing.T) {
		usecase.On("DeleteProduct", mock.Anything, mock.Anything).Return(-2).Once()

		req := httptest.NewRequest(http.MethodDelete, "/store/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/store/:id")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseMassage{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.DeleteProductCommunity))(echoContext)
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

	t.Run("Failed param", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/store/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/store/:id")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("abs")

		var responseData ResponseMassage

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(handlerTest.DeleteProductCommunity))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			fmt.Println(responseData, responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "param must be number", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

}
