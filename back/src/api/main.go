package main

import (
	"net/http"

	"./db"

	"github.com/gin-gonic/gin"
)

// var secrets = gin.H{
// 	"foo":     gin.H{"email": "foo@bar.com", "phone": "123"},
// 	"austion": gin.H{"email": "austin@example.com", "phone": "666"},
// 	"nyao":    gin.H{"email": "nyao@mails.com", "phone": "54232"},
// }

func main() {

	db.Init()

	r := gin.Default()
	r.LoadHTMLGlob("../../../front/templates/*")

	// todo一覧
	r.GET("", func(c *gin.Context) {
		animal := "neco"
		lists := db.GetAll()
		// c.HTML(http.StatusOK, "index.html", gin.H{
		// 	"lists":  lists,
		// 	"animal": animal,
		// })
		c.JSON(http.StatusOK, gin.H{
			"lists":  lists,
			"animal": animal,
		})
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

}
