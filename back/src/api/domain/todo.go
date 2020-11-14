package domain

import "time"

type TodoList struct {
	ID           int       `gorm:"column:id"`
	UserID       int       `gorm:"column:user_id"`
	Content      string    `gorm:"column:content"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	LastAchieved time.Time `gorm:"column:last_achieved"`
	IsDeleted    bool      `gorm:"column:is_deleted"`
	IsGoaled     bool      `gorm:"column:is_goaled"`
	GoaledAt     time.Time `gorm:"column:goaled_at"`
}

// 達成ログ
/* type TodoAchievedLogs struct {
	ID           int       `gorm:"column:id"`
	TodoID       int       `gorm:"column:todo_id"`
	AchievedDate time.Time `gorm:"column:achieved_date"`
} */
