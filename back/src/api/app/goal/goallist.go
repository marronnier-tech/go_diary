package goal

import (
	"../../domain"
	"../../infra"
)

func ToGetAllGoal(limit int, page int, order string) (out []allGoalArray, err error) {

	db, err := infra.DBConnect()

	if err != nil {
		return
	}

	var rows []inGoal

	base := db.Table("goal_lists").
		Select("goal_lists.ID, goal_lists.todo_id, goal_lists.count, goal_lists.goaled_at, todo_lists.Content, todo_lists.is_deleted, users.id, users.name, users.handle_name, users.img").
		Where("todo_lists.is_deleted = ?", false).
		Joins("join todo_lists on goal_lists.todo_id = todo_lists.id").
		Joins("join users on todo_lists.user_id = users.id").
		Limit(limit).
		Offset(limit * (page - 1))

	err = base.
		Order("todo_lists.last_achieved").
		Scan(&rows).
		Error

	if err != nil {
		return
	}

	var obj domain.GoalObjInfo
	var user domain.UserSimpleInfo

	for _, r := range rows {

		obj = domain.GoalObjInfo{
			TodoID:        r.ID,
			Content:       r.Content,
			GoaledAt:      r.GoaledAt,
			AchievedCount: r.Count,
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

		out = append(out, allGoalArray{
			GoalObj: obj,
			User:    user,
		})

	}

	return

}

func ToGetOneGoal(name string, order string) (out userGoalArray, err error) {
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

	var rows []domain.GoalObjInfo

	base := db.Table("todo_lists").
		Select("todo_lists.id, todo_lists.content, goal_lists.goaled_at, goal_lists.count").
		Where("todo_lists.user_id = ? and todo_lists.is_deleted = ? and todo_lists.is_goaled = ?", userID, false, true).
		Joins("todo_lists.id = goal_lists.todo_id")

	err = base.
		Order("last_achieved").
		Scan(&rows).
		Error

	if err != nil {
		return
	}

	var obj domain.GoalObjInfo
	var objArray []domain.GoalObjInfo

	for _, r := range rows {

		obj = domain.GoalObjInfo{
			TodoID:        r.TodoID,
			Content:       r.Content,
			GoaledAt:      r.GoaledAt,
			AchievedCount: r.AchievedCount,
		}

		objArray = append(objArray, obj)

	}

	out = userGoalArray{
		User:    user,
		GoalObj: objArray,
	}

	return

}
