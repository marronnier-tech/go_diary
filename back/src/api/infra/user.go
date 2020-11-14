package infra

import (
	"../domain"

	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) []domain.User {

	var user []domain.User
	db.Find(&user)

	return user
}
