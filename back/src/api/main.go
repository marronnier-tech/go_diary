package main

import (
	stc "strconv"
	"time"

	"./infra"
	"./infra/model"

	"github.com/gin-gonic/gin"
)

// var secrets = gin.H{
// 	"foo":     gin.H{"email": "foo@bar.com", "phone": "123"},
// 	"austion": gin.H{"email": "austin@example.com", "phone": "666"},
// 	"nyao":    gin.H{"email": "nyao@mails.com", "phone": "54232"},
// }

func main() {

	sqldb, gormdb := infra.DBConnect()

	r := gin.Default()
	r.LoadHTMLGlob("../../../front/templates/*")

	// todo一覧
	r.GET("", func(c *gin.Context) {
		animal := "neco"
		lists := infra.GetAll(gormdb)
		// c.HTML(http.StatusOK, "index.html", gin.H{
		// 	"lists":  lists,
		// 	"animal": animal,
		// })
		c.JSON(200, gin.H{
			"lists":  lists,
			"animal": animal,
		})
	})

	r.POST("/list", func(c *gin.Context) {

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

	})

	r.DELETE("/list", func(c *gin.Context) {
		id, _ := stc.Atoi(c.PostForm("id"))
		data := model.ToDoList{}
		gormdb.Delete(&data, id)

		c.JSON(201, nil)
	})

	r.GET("/user", func(c *gin.Context) {
		userlist := infra.GetAllUsers(gormdb)

		c.JSON(200, userlist)

	})

	// todo登録

	// r.GET("/:user", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)
	// 	if secret, ok := secrets[user]; ok {
	// 		c.HTML(http.StatusOK, "user_top.html", gin.H{
	// 			"user":  user,
	// 			"email": secret,
	// 		})
	// 	} else {
	// 		c.HTML(http.StatusOK, "user_top.html", gin.H{
	// 			"user":  user,
	// 			"email": "NONE",
	// 		})
	// 	}
	// })

	r.Run()

	sqldb.Close()

}
