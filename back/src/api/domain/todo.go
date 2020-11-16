package domain

import (
	"time"

	"github.com/lib/pq"
)

type TodoObjInfo struct {
	TodoID       int         `json:"TodoID"`
	Content      string      `json:"Content"`
	CreatedAt    time.Time   `json:"CreatedAt"`
	LastAchieved pq.NullTime `json:"LastAchieved"`
}
