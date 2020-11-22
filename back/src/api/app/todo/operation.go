package todo

import (
	"errors"
	"time"

	"../../domain"
	"../../infra"
	"../../infra/table"
	"../timecalc"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func ToPost(userid int, content string) (out OperationView, err error) {

	tx, err := infra.DBConnect()

	if err != nil {
		return
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

	var rows []sameCheck

	err = tx.Table("todo_lists").
		Select("is_deleted, content").
		Where("content = ?", content).
		Scan(&rows).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	same := false

	for _, r := range rows {
		if r.IsDeleted {
			continue
		}
		same = true
		break
	}

	if same {
		err = errors.New("同一のToDoが既に存在します")
		tx.Rollback()
		return
	}

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

	out = OperationView{
		TodoID:        data.ID,
		Content:       content,
		CreatedAt:     timecalc.PickDate(data.CreatedAt),
		LastAchieved:  "達成した日はありません",
		Count:         0,
		TodayAchieved: false,
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
		err = errors.New("user is wrong")
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

	if userid != todo.UserID {
		err = errors.New("This user is invalid")
		tx.Rollback()
		return
	}

	if todo.IsGoaled {
		err = errors.New("既にゴールしたToDoです")
		tx.Rollback()
		return
	}

	if todo.LastAchieved.Time.YearDay() == time.Now().YearDay() {
		err = errors.New("今日は既にToDoが完了しています")
		tx.Rollback()
		return
	}

	todo.LastAchieved = pq.NullTime{Time: time.Now(), Valid: true}
	todo.Count++

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

	var todo table.TodoList

	err = tx.Table("todo_lists").
		Where("id = ?", todoid).
		Scan(&todo).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	if todo.UserID != userid {
		err = errors.New("This user is invalid")
		tx.Rollback()
		return
	}

	if todo.LastAchieved.Time.YearDay() != time.Now().YearDay() {
		err = errors.New("今日のToDoは完了していないため、何も処理をしていません")
		tx.Rollback()
		return
	}

	todo.Count--

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
		err = errors.New("This user is invalid")
		tx.Rollback()
		return
	}

	todo.IsGoaled = true

	var u table.User

	data := table.GoalList{TodoID: todoid, Count: todo.Count, GoaledAt: time.Now()}

	if err = tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return
	}

	err = tx.Model(&u).
		Where("id = ?", userid).
		Update("goaled_count", gorm.Expr("goaled_count + 1")).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Save(&todo).Error; err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	return

}
