package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DBuser      = "tocchy"
	DBpass      = "remon109166"
	DBProtocol  = "tcp(127.0.0.1:3306)"
	DBname      = "daily_todo"
	DBchar      = "utf8mb4"
	DBparseTime = "True"
	DBloc       = "Local"
)

type ToDoList struct {
	ID        int       `gorm:"column:id"`
	UserID    int       `gorm:"column:user_id"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func DBConnect() *gorm.DB {
	connect := fmt.Sprintf(
		"%s:%s@%s/%s?charset=%s&parseTime=%s&loc=%s",
		DBuser, DBpass, DBProtocol, DBname, DBchar, DBparseTime, DBloc,
	)
	sqlDB, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return sqlDB
}

func Init() {
	db := DBConnect()

	db.AutoMigrate(&ToDoList{})

}

func GetAll() []ToDoList {
	db := DBConnect()

	var todo []ToDoList
	db.Find(&todo)

	return todo
}
