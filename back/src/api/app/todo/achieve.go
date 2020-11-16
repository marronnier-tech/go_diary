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
