package profile

import (
	"time"

	"../../domain"
	"../../infra"
	"../../infra/table"
)

func ToPatch(userid int, HN string, Img string, FinalGoal string,
	Profile string, Twitter string, Instagram string,
	Facebook string, Github string, URL string) (err error) {

	tx, err := infra.DBConnect()

	if err != nil {
		return
	}

	var user table.User

	err = tx.Table("users").
		Select("id").
		Where("id = ?", userid).
		Scan(&user).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Model(&user).Updates(table.User{
		HN:        &HN,
		Img:       &Img,
		FinalGoal: &FinalGoal,
		Profile:   &Profile,
		Twitter:   &Twitter,
		Instagram: &Instagram,
		Facebook:  &Facebook,
		Github:    &Github,
		URL:       &URL,
		UpdatedAt: time.Now(),
	}).Error

	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error

	return

}

func ToGetOneProfile(name string) (out domain.UserDetailInfo, err error) {

	tx, err := infra.DBConnect()

	if err != nil {
		return
	}

	var p table.User

	err = tx.Table("users").
		Where("name = ?", name).
		Scan(&p).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	out = domain.UserDetailInfo{
		ID:        p.ID,
		Name:      p.Name,
		HN:        p.HN,
		Img:       p.Img,
		FinalGoal: p.FinalGoal,
		Profile:   p.Profile,
		Twitter:   p.Twitter,
		Instagram: p.Instagram,
		Facebook:  p.Facebook,
		Github:    p.Github,
		URL:       p.URL,
	}

	err = tx.Commit().Error

	return

}
