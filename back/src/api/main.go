package main

import (
	"fmt"
	stc "strconv"
	"time"

	"./infra"
	"./infra/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	_ "gorm.io/gorm"
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

	// auth := r.Group("/admin", gin.BasicAuth(gin.Accounts{

	// 	"foo":    "bar",
	// 	"austin": "1234",
	// 	"lena":   "hello2",
	// 	"manu":   "4321",
	// }))

	r.POST("/secrets", func(c *gin.Context) {
		var sign model.User

		if err := c.Bind(&sign); err != nil {
			c.JSON(500, gin.H{"err": err})
			c.Abort()
		} else {
			id := c.PostForm("id")
			user := c.PostForm("user")
			mailaddress := c.PostForm("mailaddress")
			pass := c.PostForm("pass")

			if err := signUp(id, user, mailaddress, pass); err != nil {
				c.JSON(500, gin.H{"err": err})
			}

			c.Redirect(302, "/success")

		}
	})

	r.GET("/login", func(c *gin.Context) {
		var userauth model.User

		var hashStr []byte
		inputuser := c.PostForm("user")
		inputpass := c.PostForm("password")

		hashStr = gormdb.Select("pass").Where("name = ?", inputuser).Find(&userauth)

		if hashStr == nil {
			c.JSON(500, gin.H{"err": "you are not a user."})
		}

		err := bcrypt.CompareHashAndPassword([]byte(hashStr), []byte(inputpass))

		if err != nil {
			c.JSON(500, gin.H{"err": err})
		}

		c.JSON(200, gin.H{"user": inputuser, "status": "success"})

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

// ここからあとで分ける

func signUp(id string, user string, mailaddress string, pass string) (err error) {

	sqldb, gormdb := infra.DBConnect()
	defer sqldb.Close()

	newid, _ := stc.Atoi(id)

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	fmt.Printf("user:%s, pass:%s", user, pass)

	newuser := model.User{
		ID:          newid,
		Name:        user,
		MailAddress: mailaddress,
		Pass:        hash,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = gormdb.Create(&newuser).Error

	if err != nil {
		return err
	}

	return nil
}

func getUser(username string) model.User {
	sqldb, gormdb := infra.DBConnect()
	defer sqldb.Close()
	var user model.User
	gormdb.First(&user, "user = ", username)
	return user
}
