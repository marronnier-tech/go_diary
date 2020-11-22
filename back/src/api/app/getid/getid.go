package getid

import (
	"../../domain"
	"gorm.io/gorm"
)

func Fromname(tx *gorm.DB, name string) (user domain.UserSimpleInfo, userID int, err error) {

	var u domain.UserSimpleInfo

	err = tx.Table("users").
		Select("id, name, handle_name, img, goaled_count").
		Where("name = ?", name).
		Scan(&u).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	userID = u.UserID

	if u.UserHN == "" {
		u.UserHN = u.UserName
	}

	user = domain.UserSimpleInfo{
		UserID:      u.UserID,
		UserName:    u.UserName,
		UserHN:      u.UserHN,
		UserImg:     u.UserImg,
		GoaledCount: u.GoaledCount,
	}

	return
}
