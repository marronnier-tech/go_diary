package domain

import "time"

type User struct {
	ID          int       `gorm:"column:id"`
	Name        string    `gorm:"column:name;unique"`
	Password    []byte    `gorm:"column:password"`
	MailAddress string    `gorm:"column:mail_address"`
	HandleName  string    `gorm:"column:handle_name"`
	Img         string    `gorm:"column:img"`
	FinalGoal   string    `gorm:"column:final_goal"`
	Profile     string    `gorm:"column:profile"`
	Twitter     string    `gorm:"column:twitter"`
	Instagram   string    `gorm:"column:instagram"`
	Facebook    string    `gorm:"column:facebook"`
	Github      string    `gorm:"column:github"`
	URL         string    `gorm:"column:url"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
