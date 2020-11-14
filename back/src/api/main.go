package main

import (
	"fmt"
	stc "strconv"
	"time"

	"./ui"

	"./infra"
	"./infra/model"

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

	r.GET("", ui.GetTodo)

	mytodo := r.Group("/todo")
	{
		mytodo.POST("", ui.PostTodo)
		mytodo.DELETE("/:id", ui.DeleteTodo)
	}

	// auth := r.Group("/admin", gin.BasicAuth(gin.Accounts{

	// 	"foo":    "bar",
	// 	"austin": "1234",
	// 	"lena":   "hello2",
	// 	"manu":   "4321",
	// }))

	// ログイン

	r.POST("/secrets", func(c *gin.Context) {
		var sign model.User

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

	r.GET("/login", func(c *gin.Context) {
		var userauth model.User

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

	sqldb, err := gormdb.DB()

	if err != nil {
		fmt.Println("cannot use sqldb.")
	}

	sqldb.Close()

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

	newuser := model.User{
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

/* func getUser(username string) model.User {
	gormdb, err := infra.DBConnect()
	defer gormdb.Close()
	var user model.User
	gormdb.First(&user, "user = ", username)
	return user
} */
