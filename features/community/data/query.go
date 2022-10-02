package data

import (
	"capstone/happyApp/features/community"
	event "capstone/happyApp/features/event/data"
	user "capstone/happyApp/features/user/data"
	"errors"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) community.DataInterface {
	return &Storage{
		query: db,
	}
}

func (storage *Storage) Insert(userid int, core community.CoreCommunity) (string, error) {
	model := ToModel(core)
	tx := storage.query.Create(&model)
	if tx.Error != nil {
		return "Masalah Pada Insert Database", tx.Error
	}
	model2 := GetLeader(userid, int(model.ID))
	tx2 := storage.query.Create(&model2)
	if tx2.Error != nil {
		return "Masalah saat bergabung ke community", tx.Error
	}

	return "Suksus Bergabung ke Community", nil
}

func (storage *Storage) SelectList() ([]community.CoreCommunity, string, error) {
	var models []Community
	tx := storage.query.Find(&models)
	if tx.Error != nil {
		return nil, "Terjadi Kesalahan", tx.Error
	}

	listcore := ToCoreList(models)

	for k, v := range listcore {
		Count := storage.query.Model(&JoinCommunity{}).Where("community_id = ?", v.ID).Count(&v.Members)
		if Count.Error != nil {
			return nil, "Terjadi Kesalahan Saat Menghitung Member", Count.Error
		}
		listcore[k].Members = v.Members
	}

	return listcore, "Sukses Mendapatkan Semua Data", nil
}

func (storage *Storage) SelectMembers(communityid int) ([]string, string, error) {
	var models []JoinCommunity
	tx := storage.query.Find(&models, "community_id = ?", communityid)
	if tx.Error != nil {
		return nil, "Terjadi Kesalahan pada pengambilan member", tx.Error
	}

	var names []string
	for _, v := range models {
		var temp user.User
		txx := storage.query.Find(&temp, "id = ?", v.UserID)
		if txx.Error != nil {
			return nil, "Terjadi Kesalahan pada pengambilan nama", txx.Error
		}
		names = append(names, temp.Name)
	}

	return names, "Sekses Mendapatkan Semua Nama", nil
}

func (storage *Storage) Delete(userid, communityid int) (int64, string, error) {
	tx := storage.query.Where("user_id = ? and community_id = ?", userid, communityid).Delete(&JoinCommunity{})
	if tx.Error != nil {
		return 0, "Terjadi Kesalahan", tx.Error
	}
	var sum int64
	Count := storage.query.Model(&JoinCommunity{}).Where("community_id = ?", communityid).Count(&sum)
	if Count.Error != nil {
		return 0, "Gagal Menghitung Jumlah Member", Count.Error
	}
	return sum, "Sukses Keluar dari Community", nil
}

func (storage *Storage) GetUserRole(userid, communityid int) (string, error) {
	var model JoinCommunity
	tx := storage.query.Find(&model, "user_id = ? and community_id = ?", userid, communityid)
	if tx.Error != nil {
		return "Gagal mendapatkan role User", tx.Error
	}

	return model.Role, nil
}

func (storage *Storage) DeleteCommunity(communityid int) (string, error) {
	tx1 := storage.query.Where("id = ?", communityid).Delete(&Community{})
	if tx1.Error != nil || tx1.RowsAffected == 0 {
		return "Terjadi Kesalahan Pada Penghapusan Community", tx1.Error
	}
	tx2 := storage.query.Where("community_id = ?", communityid).Delete(&event.Event{})
	if tx2.Error != nil {
		return "Terjadi Kesalahan Pada Penghapusan Community", tx2.Error
	}
	tx3 := storage.query.Where("community_id = ?", communityid).Delete(&Product{})
	if tx3.Error != nil {
		return "Terjadi Kesalahan Pada Penghapusan Community", tx3.Error
	}
	tx4 := storage.query.Where("community_id = ?", communityid).Delete(&Feed{})
	if tx4.Error != nil {
		return "Terjadi Kesalahan Pada Penghapusan Community", tx4.Error
	}

	return "Community Terhapus", nil
}

func (storage *Storage) ChangeAdmin(communityid int) (string, string, error) {
	var member JoinCommunity
	tx1 := storage.query.Order("created_at").First(&member)
	if tx1.Error != nil {
		return "", "Gagal Mendapatkan Member Lain", tx1.Error
	}

	var name User
	tx := storage.query.Find(&name, "id = ?", member.UserID)
	if tx.Error != nil {
		return "", "Terjadi Kesalahan saat pengambilan nama", tx.Error
	}

	member.Role = "admin"
	tx2 := storage.query.Model(&member).Where("community_id = ? and user_id = ?", communityid, member.UserID).Updates(member)
	if tx2.Error != nil {
		return "", "Gagal Mengganti Admin", tx2.Error
	}

	return name.Name, "Sukses Mewariskan", nil

}

func (storage *Storage) UpdateCommunity(communityid int, core community.CoreCommunity) (string, error) {
	update := ToModel(core)
	tx := storage.query.Model(&update).Where("id = ?", communityid).Updates(update)
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Gagal Update", tx.Error
	}

	return "Sukses Update", nil

}

func (storage *Storage) CheckJoin(userid, communityid int) (string, error) {

	tx := storage.query.Find(&JoinCommunity{}, "user_id = ? and community_id = ?", userid, communityid)
	if tx.Error != nil {
		return "Gagal Chek Join", tx.Error
	} else if tx.RowsAffected == 1 {
		return "Anda Sudah ada di community", errors.New("Anda telah Bergabung")
	}

	return "Bisa Join", nil
}

func (storage *Storage) InsertToJoin(userid, communityid int) (string, error) {
	model := ToJoin(userid, communityid)
	tx := storage.query.Create(&model)
	if tx.Error != nil {
		return "Gagal Bergabung ke Community", tx.Error
	}

	return "Sukses Begabung Ke Community", nil
}

func (storage *Storage) InsertFeed(core community.CoreFeed) (string, error) {
	feed := ToModelFeed(core)
	tx := storage.query.Create(&feed)
	if tx.Error != nil {
		return "Gagal menambahkan Feed", tx.Error
	}

	return "Sukses Menambahkan Feed", nil
}

func (storage *Storage) SelectCommunity(communityid int) (community.CoreCommunity, string, error) {
	var get Community
	tx := storage.query.Preload("Feeds").Find(&get, "id = ?", communityid)
	if tx.Error != nil {
		return community.CoreCommunity{}, "Gagal Ke database", tx.Error
	}

	//harusnya bisa bikin interface sendiri
	var sum int64
	Count := storage.query.Model(&JoinCommunity{}).Where("community_id = ?", communityid).Count(&sum)
	if Count.Error != nil {
		return community.CoreCommunity{}, "Terjadi Kesalahan Saat Menghitung Member", Count.Error
	}

	var corefeed []community.CoreFeed
	for _, v := range get.Feeds {
		var name user.User
		tx1 := storage.query.Find(&name, "id = ?", v.UserID)
		if tx1.Error != nil {
			return community.CoreCommunity{}, "Gagal Pada Pengambilan Nama", tx1.Error
		}
		corefeed = append(corefeed, ToCoreFeed(v, name.Name))
	}

	last := ToCoreWithFeed(get, sum, corefeed)

	return last, "Sukses mendapat Data", nil
}

func (storage *Storage) SelectFeed(feedid int) (community.CoreFeed, string, error) {
	var feed Feed

	tx1 := storage.query.Preload("Comments").Find(&feed, "id = ?", feedid)
	if tx1.Error != nil {
		return community.CoreFeed{}, "Gagal Mendapatkan Feed Dari Database", tx1.Error
	}

	var feedname user.User
	tx2 := storage.query.Find(&feedname, "id = ?", feed.UserID)
	if tx2.Error != nil {
		return community.CoreFeed{}, "Gagal Mendapatkan Nama Feed", tx2.Error
	}

	var corecomment []community.CoreComment
	for _, v := range feed.Comments {
		var name user.User
		tx1 := storage.query.Find(&name, "id = ?", v.UserID)
		if tx1.Error != nil {
			return community.CoreFeed{}, "Gagal Pada Pengambilan Nama Comment", tx1.Error
		}
		corecomment = append(corecomment, ToCoreComment(v, name.Name))
	}
	corefeed := ToCoreWithComment(feed, feedname.Name, corecomment)

	return corefeed, "Berhasil Memproses Semua Data", nil
}

func (storage *Storage) InsertComment(core community.CoreComment) (string, error) {
	modelcomment := ToModelComment(core)
	tx := storage.query.Create(&modelcomment)
	if tx.Error != nil {
		return "Gagal Menambahkan Comment", tx.Error
	}

	return "Sukses Menambahka comment", nil
}

func (storage *Storage) SelectListCommunityWithParam(param string) ([]community.CoreCommunity, string, error) {
	var models []Community
	param = "%" + param + "%"
	tx := storage.query.Raw("select * from communities where title like ?", param).Scan(&models)
	if tx.Error != nil {
		return nil, "Terjadi Kesalahan Dalam Get Title", tx.Error
	}

	listcore := ToCoreList(models)

	for k, v := range listcore {
		Count := storage.query.Model(&JoinCommunity{}).Where("community_id = ?", v.ID).Count(&v.Members)
		if Count.Error != nil {
			return nil, "Terjadi Kesalahan Saat Menghitung Member", Count.Error
		}
		listcore[k].Members = v.Members
	}

	return listcore, "Sukses Mendapatkan Semua Data", nil
}
