package usecase

import (
	"capstone/happyApp/features/midtrans"
	"capstone/happyApp/utils/helper"
	"fmt"
	"time"
)

type Service struct {
	do midtrans.DataInterface
}

func New(data midtrans.DataInterface) midtrans.UsecaseInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) WeebHookTransaction(orderid, transactionstatus string) (string, error) {
	msg, err := service.do.WeebHookUpdateTransaction(orderid, transactionstatus)
	return msg, err
}

func (service *Service) WeebHookJoinEvent(orderid, transactionstatus string) (string, error) {

	data, err := service.do.WeebHookUpdateJoinEvent(orderid, transactionstatus)
	if err != nil {
		return "failed update  data", err
	}

	if transactionstatus == "settlement" || transactionstatus == "capture" {

		bodyEmail := helper.BodyEmail{
			Name:  data.Name,
			Title: data.TitleEvent,
			Date:  data.Date.Format(time.RFC822),
		}
		subject := fmt.Sprintf("Sukses bergabung dalam event %s", data.TitleEvent)
		errNotif := helper.SendEmail(data.Email, subject, bodyEmail)
		if errNotif != nil {
			fmt.Println("error notif")
		}

	}

	return "success update", nil

}
