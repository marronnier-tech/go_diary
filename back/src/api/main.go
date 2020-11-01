package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo":     gin.H{"email": "foo@bar.com", "phone": "123"},
	"austion": gin.H{"email": "austin@example.com", "phone": "666"},
	"nyao":    gin.H{"email": "nyao@mails.com", "phone": "54232"},
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../../../front/templates/*")

	r.GET("", func(c *gin.Context) {
		animal := "neco"
		c.HTML(http.StatusOK, "index.html", gin.H{
			"animal": animal,
		})
	})

	auth := r.Group("", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	auth.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NONE"})
		}
	})

	r.Run()

}
