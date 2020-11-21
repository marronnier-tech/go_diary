package goal

import (
	"../../domain"
)

type allGoalArray struct {
	GoalObj domain.GoalObjInfo    `json:"GoalObj"`
	User    domain.UserSimpleInfo `json:"User"`
}

type userGoalArray struct {
	User    domain.UserSimpleInfo `json:"User"`
	GoalObj []domain.GoalObjInfo  `json:"GoalObj"`
}

/*
type todayTodo struct {
	TodoLog       table.TodoAchievedLog `json:"TodoLogInfo"`
	TodayAchieved bool                  `json:"TodayAchieved"`
}

*/
