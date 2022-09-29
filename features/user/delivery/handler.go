package delivery

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/user"
	"capstone/happyApp/middlewares"
	"capstone/happyApp/utils/helper"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type userDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, data user.UsecaseInterface) {
	handler := &userDelivery{
		userUsecase: data,
	}

	e.POST("/register", handler.CreateUser)
	e.DELETE("/user/profile", handler.DeleteAccount, middlewares.JWTMiddleware())
	e.PUT("/user/profile", handler.UpdateAccount, middlewares.JWTMiddleware())
	e.GET("/user/profile", handler.GetProfile, middlewares.JWTMiddleware())

}

func (delivery *userDelivery) CreateUser(c echo.Context) error {

	var reqData Request
	err := c.Bind(&reqData)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error Bind"))
	}

	reqData.Photo = config.DEFAULT_PROFILE
	row := delivery.userUsecase.PostUser(reqData.reqToCore())
	if row == -2 {
		return c.JSON(400, helper.FailedResponseHelper("please input all request"))
	} else if row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("failed sign up"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success sign up"))

}

func (delivery *userDelivery) DeleteAccount(c echo.Context) error {

	idToken := middlewares.ExtractToken(c)

	row := delivery.userUsecase.DeleteUser(idToken)
	if row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("failed delete account"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success delete account"))

}

func (delivery *userDelivery) UpdateAccount(c echo.Context) error {

	var updateData Request
	err := c.Bind(&updateData)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	idToken := middlewares.ExtractToken(c)

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
			imageName := strconv.Itoa(idToken) + "_" + updateData.Username + waktu + "." + ext

			imageaddress, errupload := helper.UploadFileToS3(config.DirImage, imageName, config.FileImageType, fileData)
			if errupload != nil {
				return c.JSON(400, helper.FailedResponseHelper("failed to upload file"))
			}
			updateData.Photo = imageaddress
		}
	}

	var Update user.CoreUser
	if updateData.Name != "" {
		Update.Name = updateData.Name
	}
	if updateData.Username != "" {
		Update.Username = updateData.Username
	}
	if updateData.Gender != "" {
		Update.Gender = updateData.Gender
	}
	if updateData.Email != "" {
		Update.Email = updateData.Email
	}
	if updateData.Password != "" {
		Update.Password = updateData.Password
	}
	if updateData.Photo != "" {
		Update.Photo = updateData.Photo
	}

	Update.ID = uint(idToken)

	row := delivery.userUsecase.UpdateUser(Update)
	if row == -1 {
		return c.JSON(400, helper.FailedResponseHelper("failed update account"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success update account"))

}

func (delivery *userDelivery) GetProfile(c echo.Context) error {

	idToken := middlewares.ExtractToken(c)

	data, comu, err := delivery.userUsecase.GetUser(idToken)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed get profile"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get profile", toRespon(data, comu)))

}
