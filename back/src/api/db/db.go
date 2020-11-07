package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	DB         = "mysql"
	DBuser     = "tocchy"
	DBpass     = "remon109166"
	DBProtocol = "tcp(127.0.0.1:3306)"
	DBname     = "daily_todo"
)

type ToDoList struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func gormConnect() *gorm.DB {
	connect := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", DBuser, DBpass, DBProtocol, DBname)
	db, err := gorm.Open(DB, connect)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func Init() {
	con := gormConnect()

	defer con.Close()

	con.AutoMigrate(&ToDoList{})

}

func GetAll() []ToDoList {
	con := gormConnect()

	var todo []ToDoList
	con.Find(&todo)
	return todo
}
