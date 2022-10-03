package data

import (
	"capstone/happyApp/features/cart"
	community "capstone/happyApp/features/community/data"

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

func (storage *Storage) DeleteFromCart(cartid int) (string, error) {
	tx := storage.query.Where("id = ? ", cartid).Delete(&Cart{})
	if tx.Error != nil {
		return "Gagal Menghapus Cart", tx.Error
	}
	return "Success Menghapus Cart", nil
}
