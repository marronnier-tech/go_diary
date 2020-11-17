package ui

import (
	"../app/admin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionLogin(c *gin.Context) (id int, user string, err error) {

	session := sessions.Default(c)

	// var hashStr []byte
	name := session.Get("name")
	password := session.Get("password")

	if name == nil || password == nil {
		c.Redirect(302, "/login")
	}

	strname := name.(string)
	strpassword := password.(string)

	id, user, err = admin.Validation(strname, strpassword)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}
	return
}

func Login(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	session := sessions.Default(c)

	_, _, err := admin.Validation(name, password)

	if err != nil {
		c.JSON(500, gin.H{"error": "validation error"})
		return
	}

	session.Set("name", name)
	session.Set("password", password)
	session.Save()

	c.JSON(200, gin.H{"message": "success"})
	return

}

func Register(c *gin.Context) {
	session := sessions.Default(c)

	name := c.PostForm("name")
	password := c.PostForm("password")

	if err := admin.SignUp(name, password); err != nil {
		c.JSON(500, gin.H{"err": err})
	}

	session.Set("name", name)
	session.Set("password", password)
	session.Save()

	c.Redirect(302, "/success")

}
