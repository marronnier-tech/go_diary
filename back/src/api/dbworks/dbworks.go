package dbworks

import (
	"fmt"
	"time"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB          = "mysql"
	DBuser      = "tocchy"
	DBpass      = "******"
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

func DBConnect() (*sql.DB, *gorm.DB) {
	connect := fmt.Sprintf(
		"%s:%s@%s/%s?charset=%s&parseTime=%s&loc=%s",
		DBuser, DBpass, DBProtocol, DBname, DBchar, DBparseTime, DBloc,
	)

	sqlDB, err := sql.Open(DB, connect)

	if err != nil {
		panic(err.Error())
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return sqlDB, gormDB
}

func GetAll(db *gorm.DB) []ToDoList {

	var todo []ToDoList
	db.Find(&todo)

	return todo
}
