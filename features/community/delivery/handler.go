package delivery

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/community"
	"capstone/happyApp/middlewares"
	"capstone/happyApp/utils/helper"

	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	From community.UsecaseInterface
}

func New(e *echo.Echo, data community.UsecaseInterface) {
	handler := &Delivery{
		From: data,
	}
	e.GET("/community", handler.ListCommunity)                                                          // selesai
	e.POST("/community", handler.AddCommunity, middlewares.JWTMiddleware())                             // selesai
	e.GET("/community/members/:communityid", handler.ListMembersCommunity, middlewares.JWTMiddleware()) // selesai
	e.PUT("/community/:communityid", handler.UpdateCommunity, middlewares.JWTMiddleware())              // selesai
	e.DELETE("/community/:communityid", handler.OutFromCommunity, middlewares.JWTMiddleware())          // selesai
	e.POST("/join/community/:communityid", handler.JoinCommunity, middlewares.JWTMiddleware())          // selesai
	e.POST("/community/:communityid/feed", handler.AddFeed, middlewares.JWTMiddleware())                // selesai
	e.GET("/community/:communityid", handler.GetDetailCommunity, middlewares.JWTMiddleware())           // selesai
	e.GET("/feed/:feedid", handler.GetDetailFeed, middlewares.JWTMiddleware())                          // selesai
	e.POST("/feed/:feedid/comment", handler.AddComment, middlewares.JWTMiddleware())                    // selesai

}

func (user *Delivery) AddCommunity(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	var reqcom Request
	errb := c.Bind(&reqcom)
	if errb != nil {
		return c.JSON(400, helper.FailedResponseHelper("Gagal Bind Data"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("logo")
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
			imageName := strconv.Itoa(userid) + "_" + reqcom.Title + waktu + "." + ext

			imageaddress, errupload := helper.UploadFileToS3(config.DirImage, imageName, config.FileImageType, fileData)
			if errupload != nil {
				return c.JSON(400, helper.FailedResponseHelper("failed to upload file"))
			}
			reqcom.Logo = imageaddress
		}
	} else {
		reqcom.Logo = config.DEFAULT_COMMUNITY
	}
	msg, err := user.From.AddNewCommunity(userid, reqcom.ToCore())
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}

func (user *Delivery) ListCommunity(c echo.Context) error {
	param := c.QueryParam("title")
	if param != "" {
		listcore, msg, err := user.From.GetListCommunityWithParam(param)
		if err != nil {
			return c.JSON(400, helper.FailedResponseHelper(msg))
		} else if len(listcore) == 0 {
			return c.JSON(200, helper.SuccessResponseHelper("Title tidak ditemukan"))
		}
		return c.JSON(200, helper.SuccessDataResponseHelper(msg, ToResponseList(listcore)))
	}

	listcore, msg, err := user.From.GetListCommunity()
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(msg, ToResponseList(listcore)))
}

func (user *Delivery) ListMembersCommunity(c echo.Context) error {
	communityid, err := strconv.Atoi(c.Param("communityid"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}
	members, msg, errs := user.From.GetMembers(communityid)
	if errs != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(msg, members))
}

func (user *Delivery) OutFromCommunity(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	communityid, err := strconv.Atoi(c.Param("communityid"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	msg, errs := user.From.Leave(userid, communityid)
	if errs != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}

func (user *Delivery) UpdateCommunity(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	communityid, err := strconv.Atoi(c.Param("communityid"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	var reqcom Request
	errb := c.Bind(&reqcom)
	if errb != nil {
		return c.JSON(400, helper.FailedResponseHelper("Gagal Bind Data"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("logo")
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
			imageName := strconv.Itoa(userid) + "_" + reqcom.Title + waktu + "." + ext

			imageaddress, errupload := helper.UploadFileToS3(config.DirImage, imageName, config.FileImageType, fileData)
			if errupload != nil {
				return c.JSON(400, helper.FailedResponseHelper("failed to upload file"))
			}
			reqcom.Logo = imageaddress
		}
	}

	msg, err := user.From.UpdateCommunity(userid, reqcom.ToCoreWithId(communityid))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}

func (user *Delivery) JoinCommunity(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	communityid, err := strconv.Atoi(c.Param("communityid"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}
	msg, er := user.From.JoinCommunity(userid, communityid)
	if er != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}

func (user *Delivery) GetDetailCommunity(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	communityid, err := strconv.Atoi(c.Param("communityid"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}
	data, msg, errs := user.From.GetCommunityFeed(userid, communityid)
	if errs != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))

	}

	return c.JSON(200, helper.SuccessDataResponseHelper(msg, ResponseWithFeed(data)))
}

func (user *Delivery) AddFeed(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	communityid, err := strconv.Atoi(c.Param("communityid"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	var feed FeedRequst
	errb := c.Bind(&feed)
	if errb != nil {
		return c.JSON(400, helper.FailedResponseHelper("Gagal Bind Data"))
	}

	msg, errs := user.From.PostFeed(feed.ToCore(userid, communityid))
	if errs != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}

func (user *Delivery) GetDetailFeed(c echo.Context) error {
	feedid, err := strconv.Atoi(c.Param("feedid"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}
	feed, msg, ers := user.From.GetDetailFeed(feedid)
	if ers != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessFeedResponseHelper(msg, ResponseFeedWithComment(feed)))
}

func (user *Delivery) AddComment(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	feedid, err := strconv.Atoi(c.Param("feedid"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}
	var req CommentRequst
	errb := c.Bind(&req)
	if errb != nil {
		return c.JSON(400, helper.FailedResponseHelper("Gagal Bind Data"))
	}

	msg, errs := user.From.AddComment(req.ToCore(userid, feedid))
	if errs != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}
