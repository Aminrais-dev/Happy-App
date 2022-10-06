package data

import (
	"capstone/happyApp/features/cart"
	community "capstone/happyApp/features/community/data"
	user "capstone/happyApp/features/user/data"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) cart.DataInterface {
	return &Storage{
		query: db,
	}
}

func (storage *Storage) InsertIntoCart(userid, productid int) (string, error) {
	txc := storage.query.Find(&Cart{}, "user_id = ? and product_id = ?", userid, productid)
	if txc.Error != nil {
		return "Gagal Mendapatkan data cart", txc.Error
	} else if txc.RowsAffected == 1 {
		return "Barang sudah ada di cart", txc.Error
	}

	cart := ToModelCart(uint(userid), uint(productid))
	tx := storage.query.Create(&cart)
	if tx.Error != nil {
		return "Gagal Menambahkan Ke Cart", tx.Error
	}

	return "Sukses Menambahkan Ke Cart", nil
}

func (storage *Storage) GetCommunity(communityId int) (cart.CoreCommunity, string, error) {
	var cartcommun community.Community
	tx := storage.query.Find(&cartcommun, "id = ?", communityId)
	if tx.Error != nil {
		return cart.CoreCommunity{}, "Gagal Mendapatkan data Community", tx.Error
	}

	var sum int64
	Count := storage.query.Model(&JoinCommunity{}).Where("community_id = ?", communityId).Count(&sum)
	if Count.Error != nil {
		return cart.CoreCommunity{}, "Terjadi Kesalahan Saat Menghitung Member", Count.Error
	}

	return ToCoreCommunity(cartcommun, sum), "", nil
}

func (storage *Storage) SelectCartList(userid, communityid int) ([]cart.CoreCart, string, error) {
	var cartall []Cart
	tx := storage.query.Find(&cartall, "user_id = ?", userid)
	if tx.Error != nil {
		return nil, "Gagal Mendapatkan List Cart", tx.Error
	}

	var communitycart []Cart
	for _, v := range cartall {
		var product Product
		txif := storage.query.Find(&product, "id = ?", v.ProductID)
		if txif.Error != nil {
			return nil, "Kesalahan pada Pengambilan product", txif.Error
		}
		if product.CommunityID == uint(communityid) {
			communitycart = append(communitycart, v)
		}
	}

	var list []cart.CoreCart
	for _, v := range communitycart {
		var product Product
		txif := storage.query.Find(&product, "id = ?", v.ProductID)
		if txif.Error != nil {
			return nil, "Kesalahan pada Pengambilan product", txif.Error
		}

		list = append(list, ProductToCart(product, v.ID))
	}

	return list, "Sukses Mengambil Semua Data", nil
}

func (storage *Storage) DeleteFromCart(userid, cartid int) (string, error) {
	tx := storage.query.Where("id = ? and user_id = ?", cartid, userid).Delete(&Cart{})
	if tx.Error != nil {
		return "Gagal Menghapus Cart", tx.Error
	}
	return "Success Menghapus Cart", nil
}
func (storage *Storage) CheckStock(id []int) ([]int, string, error) {
	listcart := id
	var listprodid []int
	for _, v1 := range listcart {
		var cart Cart
		ty := storage.query.Find(&cart, "id = ?", v1)
		if ty.Error != nil {
			return nil, "Gagal MenCheck Cart", ty.Error
		}
		var product Product
		ty2 := storage.query.Find(&product, "id = ?", cart.ProductID)
		if ty2.Error != nil {
			return nil, "Gagal MenCheck Product Stock", ty2.Error
		}

		if product.Stock < 1 {
			return nil, fmt.Sprintf("Stock dari Product %s Telah Habis", product.Name), errors.New("Stock Habis")
		}
		listprodid = append(listprodid, int(product.ID))
	}

	return listprodid, "Sukses mendapat prod id", nil
}

func (storage *Storage) UpdateStock(listprodid []int) (string, error) {
	for _, v := range listprodid {
		var product Product
		tx := storage.query.Find(&product, "id = ?", v)
		if tx.Error != nil {
			return "Gagal Mendapatkan Product Untuk dikurangi", tx.Error
		}
		product.Stock -= 1
		tx2 := storage.query.Model(&product).Where("id = ?", v).Updates(product)
		if tx2.Error != nil {
			return "Gagal Mengurangi Product", tx2.Error
		}
	}
	return "Sukses Mengupdate stock", nil
}

func (storage *Storage) InsertIntoTransaction(core cart.CoreHistory) (int, string, error) {
	listcart := core.Carts
	transmodel := ToModelTransaction(core)
	tx1 := storage.query.Create(&transmodel)
	if tx1.Error != nil {
		return 0, "Gagal Insert ke Transaction", tx1.Error
	}

	for _, v := range listcart {
		mod := ToModelTransactionCart(int(transmodel.ID), v)
		tx2 := storage.query.Create(&mod)
		if tx2.Error != nil {
			return 0, "Gagal Ke junk table", tx2.Error
		}
	}

	return int(transmodel.ID), "Success Insert Ke Semua Table", nil
}

func (storage *Storage) GetTotalTransaction(transid int) (int, string, error) {
	var cartid []TransactionCart
	tx := storage.query.Find(&cartid, "transaction_id = ?", transid)
	if tx.Error != nil {
		return 0, "Gagal Mendapatkan List Id Cart", tx.Error
	}

	var gross int
	for _, v := range cartid {
		var datacart Cart
		tx2 := storage.query.Find(&datacart, "id = ?", v.CartID)
		if tx2.Error != nil {
			return 0, "Gagal Mendapatkan Product id", tx2.Error
		}
		var harga Product
		tx3 := storage.query.Find(&harga, "id = ?", datacart.ProductID)
		if tx3.Error != nil {
			return 0, "Gagal Mendapatkan Harga Product", tx3.Error
		}

		gross += int(harga.Price)
	}

	return gross, "Sukses Mendapatkan Gross", nil
}

func (storage *Storage) DeleteCart(core cart.CoreHistory) (string, error) {
	for _, v := range core.Carts {
		tx := storage.query.Where("id = ?", v).Delete(&Cart{})
		if tx.Error != nil {
			return "Gagal Menghapus Cart", tx.Error
		}
	}

	return "Sucess MenJalankan Semua Metod", nil
}

func (storage *Storage) InsertIntoPayment(payment cart.CorePayment) (string, error) {
	var pay Payment
	tx := storage.query.Create(&pay)
	if tx.Error != nil {
		return "Gagal Insert Payment", tx.Error
	}

	return "Success Insert Payment", nil
}

func (storage *Storage) GetUserRole(userid, communityid int) (string, error) {
	var model JoinCommunity
	tx := storage.query.Find(&model, "user_id = ? and community_id = ?", userid, communityid)
	if tx.Error != nil {
		return "Gagal mendapatkan role User", tx.Error
	}

	return model.Role, nil
}

func (storage *Storage) SelectCommunity(communityid int) (cart.CoreCommunity, string, error) {
	var commu community.Community
	tx := storage.query.Find(&commu, "id = ?", communityid)
	if tx.Error != nil {
		return cart.CoreCommunity{}, "Error Mengambil Data Community", tx.Error
	}

	return ToCoreCommunity(commu, 0), "Sukcess Mendaptkan Data Community", nil
}

func (storage *Storage) ListHistoryProduct(communityid int) ([]cart.CoreProductResponse, string, error) {
	var hist []History
	tx := storage.query.Raw("select p.id as product_id, c.user_id as user_id, p.name as name, p.photo as photo, p.price as price from products p inner join carts c on p.id = c.product_id inner join transaction_carts tc on c.id = tc.cart_id inner join transactions t on tc.transaction_id = t.id where p.community_id = ? and (t.status_payment = 'settlement' or t.status_payment = 'capture')", communityid).Scan(&hist)
	if tx.Error != nil {
		return nil, "Gagal Mendapatkan List History", tx.Error
	}

	var listhist []cart.CoreProductResponse

	for _, v := range hist {
		var user user.User
		tx2 := storage.query.Find(&user, "id = ?", v.UserID)
		if tx2.Error != nil {
			return nil, "Gagal Mengambil Nama Pembeli", tx2.Error
		}
		listhist = append(listhist, ToCoreProductResponse(v, user.Username))
	}

	return listhist, "Success Menjalankan semua Fungsi", nil
}
