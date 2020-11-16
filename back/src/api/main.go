package main

import (
	"fmt"
	stc "strconv"
	"time"

	"./ui"

	"./infra"
	"./infra/table"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	_ "gorm.io/gorm"
)

func main() {

	gormdb, err := infra.DBConnect()

	if err != nil {
		fmt.Println("error")
	}

	r := gin.Default()
	r.LoadHTMLGlob("../../../front/templates/*")

	todo := r.Group("/todo")
	{
		todo.GET("", ui.GetTodo)
		todo.GET("/:name", ui.GetOneUserTodo)
		todo.POST("", ui.PostTodo)
		todo.DELETE("/:id", ui.DeleteTodo)

		todo.POST("/:id/today", ui.PutAchieveTodo)
		todo.DELETE("/:id/today", ui.ClearAchieveTodo)

	}

	goal := r.Group("goal")
	{
		goal.PATCH("/:id", ui.PatchGoal)
		goal.GET("", ui.GetGoal)
		goal.GET("/:name", ui.GetOneUserGoal)

	}

	// auth := r.Group("/admin", gin.BasicAuth(gin.Accounts{

	// 	"foo":    "bar",
	// 	"austin": "1234",
	// 	"lena":   "hello2",
	// 	"manu":   "4321",
	// }))

	// ログイン、あとで移動

	r.POST("/admin/secrets", func(c *gin.Context) {
		var sign table.User

		if err := c.Bind(&sign); err != nil {
			c.JSON(500, gin.H{"err": err})
			c.Abort()
		} else {
			id := c.PostForm("id")
			user := c.PostForm("user")
			mailaddress := c.PostForm("mailaddress")
			password := c.PostForm("password")

			if err := signUp(id, user, mailaddress, password); err != nil {
				c.JSON(500, gin.H{"err": err})
			}

			c.Redirect(302, "/success")

		}
	})

	r.GET("/admin/login", func(c *gin.Context) {
		var userauth table.User

		// var hashStr []byte
		inputuser := c.PostForm("user")
		inputpass := c.PostForm("password")

		gormdb.Select("password").Where("name = ?", inputuser).Find(&userauth)
		selectpass := userauth.Password

		fmt.Println([]byte(inputpass))

		err := bcrypt.CompareHashAndPassword(selectpass, []byte(inputpass))

		fmt.Println(err)

		if err != nil {
			c.JSON(500, gin.H{"user": inputuser, "status": "password is wrong"})
		} else {
			c.JSON(200, gin.H{"user": inputuser, "status": "success"})
		}

	})

	r.GET("/success", func(c *gin.Context) {
		c.JSON(201, gin.H{"message": "success!"})
	})

	// userlist := infra.GetAllUsers(gormdb)
	// user := c.MustGet(gin.AuthUserKey).(string)
	// if userAdmin, ok := secrets[user]; ok {
	// 	c.JSON(200, gin.H{
	// 		"user":   user,
	// 		"secret": userAdmin,
	// 	})
	// } else {
	// 	c.JSON(200, gin.H{
	// 		"user":   user,
	// 		"secret": "NO SECRET",
	// 	})
	// }

	// })

	// ここでDBとじるのは問題なのであとでなんとかする

	/* sqldb, err := gormdb.DB()

	if err != nil {
		fmt.Println("cannot use sqldb.")
	}

	sqldb.Close() */

	r.Run()

}

// ここからあとで分ける

func signUp(id string, user string, mailaddress string, password string) (err error) {

	gormdb, err := infra.DBConnect()

	newid, _ := stc.Atoi(id)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	fmt.Println(hash)

	if err != nil {
		return err
	}

	newuser := table.User{
		ID:          newid,
		Name:        user,
		MailAddress: mailaddress,
		Password:    hash,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = gormdb.Create(&newuser).Error

	if err != nil {
		return err
	}

	return nil
}

/* func getUser(username string) table.User {
	gormdb, err := infra.DBConnect()
	defer gormdb.Close()
	var user table.User
	gormdb.First(&user, "user = ", username)
	return user
} */
