package infra

import (
	"./model"

	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) []model.User {

	var user []model.User
	db.Find(&user)

	return user
}
