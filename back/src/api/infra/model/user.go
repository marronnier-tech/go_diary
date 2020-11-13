package model

import "time"

type User struct {
	ID          int       `gorm:"column:id"`
	Name        string    `gorm:"column:name;unique"`
	MailAddress string    `gorm:"column:mail_address"`
	Password    []byte    `gorm:"column:password"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
