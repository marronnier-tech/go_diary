package todo

import (
	"errors"
	"time"

	"../../domain"
	"../../infra"
	"../../infra/table"
	"github.com/lib/pq"
)

func ToPost(userid int, content string) (err error) {

	db, err := infra.DBConnect()

	if err != nil {
		return err
	}

	var u domain.UserSimpleInfo

	err = db.Table("users").
		Select("id").
		Where("id = ?", userid).
		Scan(&u).
		Error

	if err != nil {
		return
	}

	userID := u.UserID

	data := table.TodoList{
		UserID:       userID,
		Content:      content,
		CreatedAt:    time.Now(),
		LastAchieved: pq.NullTime{Time: time.Now(), Valid: false},
		IsDeleted:    false,
		IsGoaled:     false,
		GoaledAt:     pq.NullTime{Time: time.Now(), Valid: false},
	}

	db.Create(&data)

	return

}

func ToDelete(todoid int, userid int) (err error) {

	db, err := infra.DBConnect()

	if err != nil {
		return err
	}

	var todo table.TodoList

	db.Table("todo_lists").
		Where("id = ?", todoid).
		First(&todo)

	if todo.UserID != userid {
		err = errors.New("This is not your todo!")
		return
	}

	todo.IsDeleted = true

	db.Save(&todo)

	return

}

func ToPutAchieve(todoid int, userid int) (out todayTodo, err error) {
	db, err := infra.DBConnect()
	if err != nil {
		return
	}

	var todo table.TodoList

	err = db.Table("todo_lists").
		Where("id = ?", todoid).
		First(&todo).
		Error

	if err != nil {
		return
	}

	if userid != todo.UserID {
		err = errors.New("This user is invalid")
		return
	}

	todo.LastAchieved = pq.NullTime{Time: time.Now(), Valid: true}

	db.Save(&todo)

	data := table.TodoAchievedLog{
		TodoID:       todoid,
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

func ToClearAchieve(todoid int, userid int) (out todayTodo, err error) {
	db, err := infra.DBConnect()
	if err != nil {
		return
	}

	var dellog table.TodoAchievedLog

	db.Table("todo_achieved_logs").
		Where("todo_id = ?", todoid).
		Order("achieved_date desc").
		Limit(1).
		Delete(&dellog)

	var todo table.TodoList

	err = db.Table("todo_lists").
		Where("id = ?", todoid).
		Scan(&todo).
		Error

	if err != nil {
		return
	}

	if todo.UserID != userid {
		err = errors.New("This user is invalid")
		return
	}

	var count int64

	db.Table("todo_achieved_logs").
		Where("todo_id = ?", todoid).
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
			Where("todo_id = ?", todoid).
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

func ToPatchGoal(todoid int, userid int) (err error) {
	db, err := infra.DBConnect()
	if err != nil {
		return err
	}

	var todo table.TodoList

	err = db.Table("todo_lists").
		Where("id = ?", todoid).
		First(&todo).
		Error

	if err != nil {
		return
	}

	if todo.UserID != userid {
		err = errors.New("This user is invalid")
		return
	}

	todo.IsGoaled = true
	todo.GoaledAt = pq.NullTime{Time: time.Now(), Valid: true}

	db.Save(&todo)

	return

}
