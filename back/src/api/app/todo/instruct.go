package todo

import (
	"../../domain"
)

type inGetAll struct {
	domain.TodoList
	UserName string `gorm:"column:name"`
	UserHN   string `gorm:"column:handle_name"`
	UserImg  string `gorm:"column:img"`
}

type inUserInfo struct {
	UserID   int    `gorm:"column:id"`
	UserName string `gorm:"column:name"`
	UserHN   string `json:"column:handle_name"`
	UserImg  string `json:"column:img"`
}

type inGetOneUser struct {
	domain.TodoList
}