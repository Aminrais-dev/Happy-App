package delivery

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/product"
	"capstone/happyApp/middlewares"
	"capstone/happyApp/utils/helper"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type productDelivery struct {
	productUsecase product.UsecaseInterface
}

func New(e *echo.Echo, data product.UsecaseInterface) {
	handler := &productDelivery{
		productUsecase: data,
	}

	e.POST("/community/:id/store", handler.PostNewProduct, middlewares.JWTMiddleware())
	e.DELETE("/store/:id", handler.DeleteProductCommunity, middlewares.JWTMiddleware())
	e.PUT("/store/:id", handler.UpdateProductCommunity, middlewares.JWTMiddleware())
	e.GET("/store/:id", handler.GetProductDetail, middlewares.JWTMiddleware())

}

func (delivery *productDelivery) PostNewProduct(c echo.Context) error {

	userId := middlewares.ExtractToken(c)
	id := c.Param("id")
	idComu, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(400, "param must be number")
	}

	var newData Request
	err := c.Bind(&newData)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("photo")
	if fileErr != http.ErrMissingFile || fileErr == nil {
		ext, errs := helper.CheckFileType(fileInfo.Filename)
		if errs != nil {
			return c.JSON(400, helper.FailedResponseHelper("gagal membaca file exetension"))
		}

		if ext == "jpg" || ext == "png" || ext == "jpeg" {
			err_size := helper.CheckFileSize(fileInfo.Size, config.FileImageType)
			if err_size != nil {
				return c.JSON(400, helper.FailedResponseHelper("image size error"))
			}

			waktu := fmt.Sprintf("%v", time.Now())
			imageName := strconv.Itoa(idComu) + "_" + newData.Name + waktu + "." + ext

			imageaddress, errupload := helper.UploadFileToS3(config.DirImage, imageName, config.FileImageType, fileData)
			if errupload != nil {
				return c.JSON(400, helper.FailedResponseHelper("failed to upload file"))
			}
			newData.Photo = imageaddress
		}
	} else {
		newData.Photo = config.DEFAULT_PRODUCT
	}

	row := delivery.productUsecase.PostProduct(newData.resToCore(idComu), userId)
	if row == -3 {
		return c.JSON(400, helper.FailedResponseHelper("please input all request"))
	} else if row == -2 {
		return c.JSON(400, helper.FailedResponseHelper("not have access in community"))
	} else if row == -1 {
		return c.JSON(500, helper.FailedResponseHelper("Failed create product"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success add product"))

}

func (delivery *productDelivery) DeleteProductCommunity(c echo.Context) error {

	userId := middlewares.ExtractToken(c)
	id := c.Param("id")
	idProduct, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(400, "param must be number")
	}

	row := delivery.productUsecase.DeleteProduct(idProduct, userId)
	if row == -2 {
		return c.JSON(400, helper.FailedResponseHelper("not have access in community"))
	} else if row == -1 {
		return c.JSON(500, helper.FailedResponseHelper("Failed delete product"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success delete product"))

}

func (delivery *productDelivery) UpdateProductCommunity(c echo.Context) error {

	userId := middlewares.ExtractToken(c)
	id := c.Param("id")
	idProduct, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(400, "param must be number")
	}

	var dataBind Request
	err := c.Bind(&dataBind)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("photo")
	if fileErr != http.ErrMissingFile || fileErr == nil {
		ext, errs := helper.CheckFileType(fileInfo.Filename)
		if errs != nil {
			return c.JSON(400, helper.FailedResponseHelper("gagal membaca file exetension"))
		}

		if ext == "jpg" || ext == "png" || ext == "jpeg" {
			err_size := helper.CheckFileSize(fileInfo.Size, config.FileImageType)
			if err_size != nil {
				return c.JSON(400, helper.FailedResponseHelper("image size error"))
			}

			waktu := fmt.Sprintf("%v", time.Now())
			imageName := strconv.Itoa(idProduct) + "_" + dataBind.Name + waktu + "." + ext

			imageaddress, errupload := helper.UploadFileToS3(config.DirImage, imageName, config.FileImageType, fileData)
			if errupload != nil {
				return c.JSON(400, helper.FailedResponseHelper("failed to upload file"))
			}
			dataBind.Photo = imageaddress
		}
	}

	var UpdateData product.ProductCore
	if dataBind.Name != "" {
		UpdateData.Name = dataBind.Name
	}
	if dataBind.Description != "" {
		UpdateData.Description = dataBind.Description
	}
	if dataBind.Photo != "" {
		UpdateData.Photo = dataBind.Photo
	}
	if dataBind.Stock != 0 {
		UpdateData.Stock = dataBind.Stock
	}
	if dataBind.Price != 0 {
		UpdateData.Price = dataBind.Price
	}

	UpdateData.ID = uint(idProduct)

	row := delivery.productUsecase.UpdateProduct(UpdateData, userId)
	if row == -2 {
		return c.JSON(400, helper.FailedResponseHelper("not have access in community"))
	} else if row == -1 {
		return c.JSON(500, helper.FailedResponseHelper("Failed update product"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success update product"))

}

func (delivery *productDelivery) GetProductDetail(c echo.Context) error {

	userId := middlewares.ExtractToken(c)
	id := c.Param("id")
	idProduct, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(400, "param must be number")
	}

	dataComu, dataProduct, err := delivery.productUsecase.GetProduct(idProduct, userId)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Failed get product detail"))
	} else if dataProduct.ID != uint(idProduct) {
		return c.JSON(404, helper.FailedResponseHelper("product not found"))
	}

	return c.JSON(200, ResponseDetail(dataComu, dataProduct))

}
