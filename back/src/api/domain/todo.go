package domain

import "time"

type AchieveInfo struct {
	Last  string `json:"LastAchieved"`
	Today bool   `json:"TodayAchieved"`
}

type TodoObjInfo struct {
	TodoID    int    `json:"TodoID"`
	IsDeleted bool   `json:"IsDeleted"`
	Content   string `json:"Content"`
	CreatedAt string `json:"CreatedAt"`
	AchieveInfo
}

type GoalObjInfo struct {
	TodoID        int       `json:"TodoID"`
	Content       string    `json:"content"`
	GoaledAt      time.Time `json:"GoaledAt"`
	AchievedCount int64     `json:"AchievedCount`
}
