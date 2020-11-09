package infra

import (
	"./model"

	"gorm.io/gorm"
)

func GetAll(db *gorm.DB) []model.ToDoList {

	var todo []model.ToDoList
	db.Find(&todo)

	return todo
}
