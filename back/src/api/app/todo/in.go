package todo

import (
	"../../infra/table"
)

type inGetAll struct {
	table.TodoList
	UserName string `gorm:"column:name"`
	UserHN   string `gorm:"column:handle_name"`
	UserImg  string `gorm:"column:img"`
}
