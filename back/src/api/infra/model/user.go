package model

import "time"

type User struct {
	ID          int       `gorm:"column:id"`
	Name        string    `gorm:"column:name;unique"`
	MailAddress string    `gorm:"column:mail_address"`
	Pass        []byte    `gorm:"column:pass"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
