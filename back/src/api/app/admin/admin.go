package admin

import (
	"fmt"
	"time"

	"../../domain"
	"../../infra"
	"../../infra/table"
	"golang.org/x/crypto/bcrypt"
)

func Validation(name string, password string) (id int, user string, err error) {
	db, err := infra.DBConnect()

	if err != nil {
		return
	}

	var userauth domain.LoginInfo

	db.Table("users").Select("id, name, password").Where("name = ?", name).Find(&userauth)
	selectpass := userauth.Password

	err = bcrypt.CompareHashAndPassword(selectpass, []byte(password))

	if err != nil {
		return
	}

	id = userauth.UserID
	user = userauth.UserName
	return
}

func SignUp(user string, password string) (err error) {

	gormdb, err := infra.DBConnect()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	fmt.Println(hash)

	if err != nil {
		return err
	}

	newuser := table.User{
		Name:        user,
		Password:    hash,
		MailAddress: nil,
		HN:          nil,
		Img:         nil,
		FinalGoal:   nil,
		Profile:     nil,
		Twitter:     nil,
		Instagram:   nil,
		Facebook:    nil,
		Github:      nil,
		URL:         nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}

	err = gormdb.Create(&newuser).Error

	if err != nil {
		return err
	}

	return nil
}

func ToDeleteMember(userid int) (err error) {
	db, err := infra.DBConnect()

	if err != nil {
		return
	}

	var user table.User

	err = db.Table("users").
		Where("id = ?", userid).
		First(&user).
		Error

	if err != nil {
		return
	}

	user.IsDeleted = true

	db.Save(&user)

	return

}
