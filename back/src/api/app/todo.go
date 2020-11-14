package app

import (
	stc "strconv"
	"time"

	"../infra"
	"../infra/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTodo() (c *gin.Context) {
	gormdb := getDB()

	animal := "neco"
	lists := infra.GetAll(gormdb)
	c.JSON(200, gin.H{
		"lists":  lists,
		"animal": animal,
	})

}

func PutMyTodo() (c *gin.Context) {
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

}

// gormdb取得
func getDB(c *gin.Context) (gormdb *gorm.DB) {
	gormdb, err := infra.DBConnect()

	if err != nil {
		c.JSON(500, gin.H{"message": "server error."})
	}

	return gormdb
}
