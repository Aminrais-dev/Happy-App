package data

import (
	cart "capstone/happyApp/features/cart/data"
	"fmt"
	"strconv"
	"strings"

	event "capstone/happyApp/features/event/data"
	"capstone/happyApp/features/midtrans"
	user "capstone/happyApp/features/user/data"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) midtrans.DataInterface {
	return &Storage{
		query: db,
	}
}

func (storage *Storage) WeebHookUpdateTransaction(orderid, status string) (string, error) {
	var trans cart.Transaction
	trans.Status_Payment = status
	trans.OrderID = orderid

	slice := strings.Split(orderid, "-")
	id, err := strconv.Atoi(slice[1])
	if err != nil {
		return "Id Hanya Bisa Angka", err
	}

	tx := storage.query.Model(&cart.Transaction{}).Where("id = ?", id).Updates(trans)
	if tx.Error != nil {
		return "Gagal Update Data Transaction", tx.Error
	}

	var back cart.Transaction
	tx2 := storage.query.Find(&back, "id = ?", id)
	if tx2.Error != nil {
		return "Gagal Update Data Transaction", tx2.Error
	}
	if back.Status_Payment == "deny" || back.Status_Payment == "cancel" || back.Status_Payment == "expire" {
		var zonk []cart.TransactionCart
		tx3 := storage.query.Find(&zonk, "transaction_id = ?", id)
		if tx3.Error != nil {
			return "Gagal Mengambil data yang akan di kembalikan", tx3.Error
		}

		for _, v := range zonk {
			var cartback cart.Cart
			tx4 := storage.query.Find(&cartback, "transaction_id = ?", v.CartID)
			if tx4.Error != nil {
				return "Gagal Mengambil id cart yang akan di kembalikan", tx4.Error
			}
			var product cart.Product
			tx5 := storage.query.Find(&product, "transaction_id = ?", cartback.ProductID)
			if tx5.Error != nil {
				return "Gagal Mengambil id product akan di kembalikan", tx5.Error
			}
			product.Stock += 1
			tx6 := storage.query.Model(&cart.Product{}).Where("id = ?", product.ID).Updates(product)
			if tx6.Error != nil {
				return "Gagal Mengembalikan Pembatalan Checkout", tx6.Error
			}
		}
	}

	return fmt.Sprintf("Pembelian atas id %s Sukses", orderid), nil
}

func (storage *Storage) WeebHookUpdateJoinEvent(orderID, status string) (midtrans.DropData, error) {

	var returnDefault midtrans.DropData
	if status == "settlement" || status == "capture" {
		status = "paid"
	}
	tx := storage.query.Model(&event.JoinEvent{}).Where("order_id = ? ", orderID).Update("status_payment", status)
	if tx.Error != nil {
		return returnDefault, tx.Error
	}

	var split = strings.Split(orderID, "-")
	idUser, err := strconv.Atoi(split[2])
	if err != nil {
		return returnDefault, err
	}
	idEvent, errId := strconv.Atoi(split[1])
	if err != nil {
		return returnDefault, errId
	}

	var dataEvent event.Event
	var dataUser user.User
	txDropEvent := storage.query.First(&dataEvent, "id = ? ", idEvent)
	if txDropEvent.Error != nil {
		return returnDefault, txDropEvent.Error
	}
	txDropUser := storage.query.First(&dataUser, "id = ? ", idUser)
	if txDropUser.Error != nil {
		return returnDefault, txDropUser.Error
	}

	return midtrans.DropData{
		Date:       dataEvent.Date,
		Name:       dataUser.Name,
		TitleEvent: dataEvent.Title,
		Email:      dataUser.Email,
	}, nil

}
