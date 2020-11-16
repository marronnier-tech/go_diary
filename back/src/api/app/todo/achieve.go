package todo

import (
	"time"

	"../../infra"
	"../../infra/table"
	"github.com/lib/pq"
)

func ToPutAchieve(id int) (out todayTodo, err error) {
	db, err := infra.DBConnect()
	if err != nil {
		return
	}

	var todo table.TodoList

	err = db.Table("todo_lists").
		Where("id = ?", id).
		First(&todo).
		Error

	if err != nil {
		return
	}

	todo.LastAchieved = pq.NullTime{Time: time.Now(), Valid: true}

	db.Save(&todo)

	data := table.TodoAchievedLog{
		TodoID:       id,
		AchievedDate: pq.NullTime{Time: time.Now(), Valid: true},
	}

	db.Create(&data)

	out = todayTodo{
		TodoLog: table.TodoAchievedLog{
			ID:           data.ID,
			TodoID:       data.TodoID,
			AchievedDate: data.AchievedDate,
		},
		TodayAchieved: true,
	}

	return

}

func ToClearAchieve(id int) (out todayTodo, err error) {
	db, err := infra.DBConnect()
	if err != nil {
		return
	}

	var dellog table.TodoAchievedLog

	db.Table("todo_achieved_logs").
		Where("todo_id = ?", id).
		Order("achieved_date desc").
		Limit(1).
		Delete(&dellog)

	var todo table.TodoList

	err = db.Table("todo_lists").
		Where("id = ?", id).
		Scan(&todo).
		Error

	var count int64

	db.Table("todo_achieved_logs").
		Where("todo_id = ?", id).
		Count(&count)

	if count == 0 {
		todo.LastAchieved = pq.NullTime{
			Time:  time.Now(),
			Valid: false,
		}

		out = todayTodo{
			TodoLog: table.TodoAchievedLog{
				ID:           0,
				TodoID:       todo.ID,
				AchievedDate: todo.LastAchieved,
			},
			TodayAchieved: false,
		}

	} else {
		var lastlog table.TodoAchievedLog

		db.Table("todo_achieved_logs").
			Where("todo_id = ?", id).
			Order("achieved_date desc").
			First(&lastlog)

		todo.LastAchieved = pq.NullTime{
			Time:  lastlog.AchievedDate.Time,
			Valid: true,
		}

		out = todayTodo{
			TodoLog: table.TodoAchievedLog{
				ID:           lastlog.ID,
				TodoID:       lastlog.TodoID,
				AchievedDate: lastlog.AchievedDate,
			},
			TodayAchieved: false,
		}
	}

	db.Save(&todo)

	return
}

/* func ToPatchGoal(id int) (err error) {
	db, err := infra.DBConnect()
	if err != nil {
		return err
	}

	var todo table.TodoList

	err = db.Table("todo_lists").
		Where("id = ?", id).
		First(&todo).
		Error

	if err != nil {
		return err
	}

	todo.IsGoaled = true
	todo.GoaledAt = pq.NullTime{Time: time.Now(), Valid: true}

	db.Save(&todo)

	return

}

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

		obj = domain.TodoObjInfo{
			TodoID:       r.ID,
			Content:      r.Content,
			CreatedAt:    r.CreatedAt,
			LastAchieved: r.LastAchieved,
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

		obj = domain.TodoObjInfo{
			TodoID:       r.ID,
			Content:      r.Content,
			CreatedAt:    r.CreatedAt,
			LastAchieved: r.LastAchieved,
		}

		objArray = append(objArray, obj)

	}

	out = userTodoArray{
		User:    user,
		TodoObj: objArray,
	}

	return

}
*/
