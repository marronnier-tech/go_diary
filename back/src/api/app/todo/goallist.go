package todo

import (
	"../../domain"
	"../../infra"
	"../../infra/table"
	"../timecalc"
)

func ToGetAllGoal(limit int, page int, order string) (out []allTodoArray, err error) {

	db, err := infra.DBConnect()

	if err != nil {
		return
	}

	var rows []inGetAll

	base := db.Table("todo_lists").
		Select("todo_lists.id, todo_lists.Content, todo_lists.user_id, todo_lists.created_at, todo_lists.last_achieved, todo_lists.is_deleted, todo_lists.is_goaled, users.name, users.handle_name, users.img").
		Where("todo_lists.is_deleted = ? and todo_lists.is_goaled = ?", false, true).
		Joins("left join users on users.ID = todo_lists.user_id").
		Limit(limit).
		Offset(limit * (page - 1))

	err = base.
		Order("todo_lists.last_achieved").
		Scan(&rows).
		Error

	if err != nil {
		return
	}

	var obj domain.TodoObjInfo
	var user domain.UserSimpleInfo

	for _, r := range rows {

		fromLastAchieved := timecalc.DifftoNow(r.LastAchieved)

		obj = domain.TodoObjInfo{
			TodoID:       r.ID,
			Content:      r.Content,
			CreatedAt:    r.CreatedAt,
			LastAchieved: fromLastAchieved,
		}

		if r.UserHN == "" {
			r.UserHN = r.UserName

		}

		user = domain.UserSimpleInfo{
			UserID:   r.UserID,
			UserName: r.UserName,
			UserHN:   r.UserHN,
			UserImg:  r.UserImg,
		}

		out = append(out, allTodoArray{
			TodoObj: obj,
			User:    user,
		})

	}

	return

}

func ToGetOneGoal(name string, order string) (out userTodoArray, err error) {
	db, err := infra.DBConnect()

	if err != nil {
		return
	}

	var u domain.UserSimpleInfo

	err = db.Table("users").
		Select("id, name, handle_name, img").
		Where("name = ?", name).
		Scan(&u).
		Error

	if err != nil {
		return
	}

	userID := u.UserID

	if u.UserHN == "" {
		u.UserHN = u.UserName

	}

	user := domain.UserSimpleInfo{
		UserID:   u.UserID,
		UserName: u.UserName,
		UserHN:   u.UserHN,
		UserImg:  u.UserImg,
	}

	var rows []table.TodoList

	base := db.Table("todo_lists").
		Select("id, user_id, content, created_at, last_achieved, is_deleted, is_goaled").
		Where("user_id = ? and is_deleted = ? and is_goaled = ?", userID, false, true)

	err = base.
		Order("last_achieved").
		Scan(&rows).
		Error

	if err != nil {
		return
	}

	var obj domain.TodoObjInfo
	var objArray []domain.TodoObjInfo

	for _, r := range rows {

		fromLastAchieved := timecalc.DifftoNow(r.LastAchieved)

		obj = domain.TodoObjInfo{
			TodoID:       r.ID,
			Content:      r.Content,
			CreatedAt:    r.CreatedAt,
			LastAchieved: fromLastAchieved,
		}

		objArray = append(objArray, obj)

	}

	out = userTodoArray{
		User:    user,
		TodoObj: objArray,
	}

	return

}
