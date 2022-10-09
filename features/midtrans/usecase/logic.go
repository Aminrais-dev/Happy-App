package usecase

import (
	"capstone/happyApp/features/midtrans"
	"capstone/happyApp/utils/helper"
	"fmt"
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
	data, msg, err := service.do.WeebHookUpdateTransaction(orderid, transactionstatus)
	if err != nil {
		return msg, err
	}
	if transactionstatus == "settlement" || transactionstatus == "capture" {
		email := helper.Email{
			Name:   data.Name,
			Status: "Success",
		}
		subject := fmt.Sprintf("Pembayaran Atas Item dengan Id %s, Telah Selesai", orderid)
		helper.SendEmailTransNotif(data.Email, subject, email)
	}
	return "Success Update", nil
}

func (service *Service) WeebHookJoinEvent(orderid, transactionstatus string) (string, error) {

	data, err := service.do.WeebHookUpdateJoinEvent(orderid, transactionstatus)
	if err != nil {
		return "failed update data", err
	}

	if transactionstatus == "settlement" || transactionstatus == "capture" {

		dataEmail := helper.BodyEmail{
			Name:  data.Name,
			Event: data.TitleEvent,
			Date:  data.Date.Format("Monday, 02-Jan-06 15:04:05") + " WIB",
		}

		subject := fmt.Sprintf("Sukses bergabung dalam event %s", data.TitleEvent)
		helper.SendEmailNotif(data.Email, subject, dataEmail)

	}

	return "success update", nil

}

// nothing
