package usecase

import (
	"capstone/happyApp/config"
	"capstone/happyApp/features/cart"
	"capstone/happyApp/features/cart/delivery"
	"capstone/happyApp/utils/helper"
	"errors"
	"fmt"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Service struct {
	do cart.DataInterface
}

func New(data cart.DataInterface) cart.UsecaseInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) AddToCart(userid, productid int) (string, error) {
	msg, err := service.do.InsertIntoCart(userid, productid)
	return msg, err
}

func (service *Service) GetCartList(userid, communityid int) (cart.CoreCommunity, []cart.CoreCart, string, error) {
	corecommunity, msg1, ers := service.do.GetCommunity(communityid)
	if ers != nil {
		return cart.CoreCommunity{}, nil, msg1, ers
	}

	listcart, msg, err := service.do.SelectCartList(userid, communityid)
	return corecommunity, listcart, msg, err
}

func (service *Service) DeleteFromCart(userid, cartid int) (string, error) {
	msg, err := service.do.DeleteFromCart(userid, cartid)
	return msg, err
}

func (service *Service) InsertIntoTransaction(core cart.CoreHistory) (int, int, string, error) {
	// check stock
	listProduct, msgp, errp := service.do.CheckStock(core.Carts)
	if errp != nil {
		return 0, 0, msgp, errp
	}
	// update stock
	msgu, erru := service.do.UpdateStock(listProduct)
	if erru != nil {
		return 0, 0, msgu, erru
	}
	// insert to Transaction
	transid, msg, err := service.do.InsertIntoTransaction(core)
	if err != nil {
		return 0, 0, msg, err
	}

	// Get Gross
	gross, msg2, err2 := service.do.GetTotalTransaction(transid)
	if err2 != nil {
		return 0, 0, msg2, err2
	}

	// delete cart
	msg3, err3 := service.do.DeleteCart(core)

	return transid, gross, msg3, err3
}

func (service *Service) GetCharge(transid, gross int, payment, table string) (coreapi.ChargeReq, string, error) {
	var charge coreapi.ChargeReq
	time := time.Now().Unix()

	switch {
	case payment == "GOPAY":
		charge = coreapi.ChargeReq{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  helper.GenerateTransactionID(table, transid) + "-GOPAY-" + fmt.Sprintf("%v", time),
				GrossAmt: int64(gross),
			},
		}
	case payment == "BCA_VIRTUAL_ACCOUNT":
		charge = coreapi.ChargeReq{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  helper.GenerateTransactionID(table, transid) + "-BCA-" + fmt.Sprintf("%v", time),
				GrossAmt: int64(gross),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
		}
	case payment == "MANDIRI_VIRTUAL_ACCOUNT":
		charge = coreapi.ChargeReq{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  helper.GenerateTransactionID(table, transid) + "-MANDIRI-" + fmt.Sprintf("%v", time),
				GrossAmt: int64(gross),
			},
			EChannel: &coreapi.EChannelDetail{
				BillInfo1: "Buy at Alterra",
				BillInfo2: "Terima Kasih",
			},
		}
	default:
		return coreapi.ChargeReq{}, "case tidak terpenuhi", errors.New("Failed")
	}

	return charge, charge.TransactionDetails.OrderID, nil
}

func (service *Service) ChargeRequest(core coreapi.ChargeReq, payment string) (coreapi.ChargeReq, string, error) {
	var corecharge coreapi.ChargeReq
	switch {
	case payment == config.BCA_VIRTUAL_ACCOUNT:
		corecharge = delivery.ToCoreBCA(core)
	case payment == config.MANDIRI_VIRTUAL_ACCOUNT:
		corecharge = delivery.ToCoreMandiri(core)
	case payment == config.GOPAY:
		corecharge = delivery.ToCoreGopay(core)
	default:
		corecharge = delivery.ToCoreMidtransBank(core)
		return corecharge, "Gagal Format Charge ke Core", errors.New("Failed")
	}

	return corecharge, "Success To Core", nil
}

func (service *Service) InsertIntoPayment(payment cart.CorePayment) (string, error) {
	msg, err := service.do.InsertIntoPayment(payment)
	return msg, err
}

func (service *Service) GetCommunityHistory(userid, communityid int) (cart.CoreCommunity, []cart.CoreProductResponse, string, error) {
	role, errr := service.do.GetUserRole(userid, communityid)
	if errr != nil {
		return cart.CoreCommunity{}, nil, "Error Mendapatkan role", errr
	} else if role != "admin" {
		return cart.CoreCommunity{}, nil, "Hanya Admin Yang Bisa Melihat History", errors.New("Access Illegal")
	}

	cart, msg1, err1 := service.do.SelectCommunity(communityid)
	if err1 != nil {
		return cart, nil, msg1, err1
	}

	listhistory, msg2, err2 := service.do.ListHistoryProduct(communityid)
	return cart, listhistory, msg2, err2

}

//
