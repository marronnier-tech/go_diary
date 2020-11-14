package todo

import (
	"time"

	"../../infra"
	"../../infra/model"
)

func ToGetAll() (todo []model.ToDoList, err error) {

	db, err := infra.DBConnect()

	if err != nil {
		return nil, err
	}

	db.Find(&todo)
	return

}

func ToPost(id int, user int, content string) (err error) {

	db, err := infra.DBConnect()

	if err != nil {
		return err
	}

	data := model.ToDoList{
		ID:        id,
		UserID:    user,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.Create(&data)

	return

}

func ToDelete(id int) (err error) {

	db, err := infra.DBConnect()

	if err != nil {
		return err
	}

	data := model.ToDoList{}
	db.Delete(&data, id)

	return

}
