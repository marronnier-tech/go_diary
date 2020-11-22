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

	tx, err := infra.DBConnect()

	if err != nil {
		return err
	}

	var u domain.UserSimpleInfo

	err = tx.Table("users").
		Select("id").
		Where("id = ?", userid).
		Scan(&u).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	userID := u.UserID

	data := table.TodoList{
		UserID:       userID,
		Content:      content,
		CreatedAt:    time.Now(),
		LastAchieved: pq.NullTime{Time: time.Now(), Valid: false},
	}

	if err = tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error

	return

}

func ToDelete(todoid int, userid int) (err error) {

	tx, err := infra.DBConnect()

	if err != nil {
		return err
	}

	var todo table.TodoList

	err = tx.Table("todo_lists").
		Where("id = ?", todoid).
		First(&todo).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	if todo.UserID != userid {
		err = errors.New("Error:user is wrong")
		tx.Rollback()
		return
	}

	todo.IsDeleted = true

	if err = tx.Save(&todo).Error; err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error

	return

}

func ToPutAchieve(todoid int, userid int) (out todayTodo, err error) {
	tx, err := infra.DBConnect()
	if err != nil {
		return
	}

	var todo table.TodoList

	err = tx.Table("todo_lists").
		Where("id = ?", todoid).
		First(&todo).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	todo.LastAchieved = pq.NullTime{Time: time.Now(), Valid: true}
	todo.Count--

	if userid != todo.UserID {
		err = errors.New("Error:This user is invalid")
		tx.Rollback()
		return
	}

	if err = tx.Save(&todo).Error; err != nil {
		tx.Rollback()
		return
	}

	data := table.TodoAchievedLog{
		TodoID:       todoid,
		AchievedDate: pq.NullTime{Time: time.Now(), Valid: true},
	}

	if err = tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return
	}

	out = todayTodo{
		TodoLog: table.TodoAchievedLog{
			ID:           data.ID,
			TodoID:       data.TodoID,
			AchievedDate: data.AchievedDate,
		},
		TodayAchieved: true,
	}

	err = tx.Commit().Error

	return

}

func ToClearAchieve(todoid int, userid int) (out todayTodo, err error) {
	tx, err := infra.DBConnect()
	if err != nil {
		return
	}

	var dellog table.TodoAchievedLog

	err = tx.Table("todo_achieved_logs").
		Where("todo_id = ?", todoid).
		Order("achieved_date desc").
		Limit(1).
		Delete(&dellog).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	var todo table.TodoList

	err = tx.Table("todo_lists").
		Where("id = ?", todoid).
		Scan(&todo).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	todo.Count--

	if todo.UserID != userid {
		err = errors.New("Error:This user is invalid")
		tx.Rollback()
		return
	}

	var counter int64

	err = tx.Table("todo_achieved_logs").
		Where("todo_id = ?", todoid).
		Count(&counter).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	if counter == 0 {
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

		err = tx.Table("todo_achieved_logs").
			Where("todo_id = ?", todoid).
			Order("achieved_date desc").
			First(&lastlog).
			Error

		if err != nil {
			tx.Rollback()
			return
		}

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

	if err = tx.Save(&todo).Error; err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error

	return
}

func ToPatchGoal(todoid int, userid int) (err error) {
	tx, err := infra.DBConnect()
	if err != nil {
		return err
	}

	var todo table.TodoList

	err = tx.Table("todo_lists").
		Where("id = ?", todoid).
		First(&todo).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	if todo.UserID != userid {
		err = errors.New("Error:This user is invalid")
		tx.Rollback()
		return
	}

	todo.IsGoaled = true

	var u domain.UserSimpleInfo

	err = tx.Table("users").
		Select("id, name, handle_name, img, goaled_count").
		Where("id = ?", userid).
		First(&u).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	u.GoaledCount++

	data := table.GoalList{TodoID: todoid, Count: todo.Count, GoaledAt: time.Now()}

	if err = tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Save(&u).Error; err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Save(&todo).Error; err != nil {
		tx.Rollback()
		return
	}

	return

}
