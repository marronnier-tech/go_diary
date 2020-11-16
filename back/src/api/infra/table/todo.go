package table

import (
	"time"

	"github.com/lib/pq"
)

type TodoList struct {
	ID           int         `gorm:"column:id;autoIncrement"`
	UserID       int         `gorm:"column:user_id"`
	Content      string      `gorm:"column:content"`
	CreatedAt    time.Time   `gorm:"column:created_at"`
	LastAchieved pq.NullTime `gorm:"column:last_achieved"`
	IsDeleted    bool        `gorm:"column:is_deleted"`
	IsGoaled     bool        `gorm:"column:is_goaled"`
	GoaledAt     pq.NullTime `gorm:"column:goaled_at"`
}

// 達成ログ

type TodoAchievedLog struct {
	ID           int         `gorm:"column:id; autoIncrement"`
	TodoID       int         `gorm:"column:todo_id"`
	AchievedDate pq.NullTime `gorm:"column:achieved_date"`
}
