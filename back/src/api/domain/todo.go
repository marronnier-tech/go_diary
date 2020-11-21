package domain

import (
	"time"
)

type TodoObjInfo struct {
	TodoID       int       `json:"TodoID"`
	Content      string    `json:"Content"`
	CreatedAt    time.Time `json:"CreatedAt"`
	LastAchieved string    `json:"LastAchieved"`
}
