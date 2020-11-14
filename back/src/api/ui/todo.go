package ui

import (
	"fmt"
	stc "strconv"
	"time"

	"../infra"
	"../infra/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTodo(c *gin.Context) {
	gormdb := getDB()

	animal := "neco"
	lists := infra.GetAll(gormdb)
	c.JSON(200, gin.H{
		"lists":  lists,
		"animal": animal,
	})

	return

}

func PutMyTodo(c *gin.Context) {
	gormdb := getDB()

	id, _ := stc.Atoi(c.PostForm("id"))
	user, _ := stc.Atoi(c.PostForm("user"))
	content := c.PostForm("content")

	data := model.ToDoList{
		ID:        id,
		UserID:    user,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	gormdb.Create(&data)
	c.JSON(201, nil)

	return

}

func DeleteMyTodo(c *gin.Context) {
	gormdb := getDB()

	id, _ := stc.Atoi(c.Param("id"))
	data := model.ToDoList{}

	gormdb.Delete(&data, id)

	c.JSON(201, nil)

	return
}

// gormdb取得
func getDB() (gormdb *gorm.DB) {
	gormdb, err := infra.DBConnect()

	if err != nil {
		fmt.Println("error")
	}

	return gormdb
}
