package todo

import (
	"time"

	"../../domain"
	"../../infra"
)

func ToGetAll() (out getAll, err error) {

	db, err := infra.DBConnect()

	if err != nil {
		return nil, err
	}

	var rows []domain.TodoList

	/* err = db.Table("todo_lists").
	Select("todo_lists.ID, todo_lists.Content, todo_lists.last_achieved, users.ID, users.name, users.handlename, users.img").
	Joins("left join users on users.ID = todo_lists.ID").
	Scan(&rows).
	Error */

	err = db.Table("todo_lists").
		Select("todo_lists.Content, todo_lists.user_id, todo_lists.content, users.name").
		Joins("left join users on users.ID = todo_lists.ID").
		Scan(&rows).
		Error

	if err != nil {
		return nil, err
	}

	var info []todoArray

	limit := 100

	for i, r := range rows {
		if i > limit {
			break
		}
		info[i] = todoArray{
			Content:  r.Content,
			UserID:   r.UserID,
			UserName: r.Users.Name,
		}

	}

	out = getAll{info, limit}

	return

}

func ToPost(id int, user int, content string) (err error) {

	db, err := infra.DBConnect()

	if err != nil {
		return err
	}

	data := domain.TodoList{
		ID:        id,
		UserID:    user,
		Content:   content,
		CreatedAt: time.Now(),
		IsDeleted: false,
		IsGoaled:  false,
	}

	db.Create(&data)

	return

}

func ToDelete(id int) (err error) {

	db, err := infra.DBConnect()

	if err != nil {
		return err
	}

	data := domain.TodoList{}
	db.Delete(&data, id)

	return

}
