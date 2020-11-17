package ui

import (
	"fmt"
	"time"

	"./../domain"
	"./../infra"
	"./../infra/table"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SessionLogin(c *gin.Context) (member int, err error) {

	session := sessions.Default(c)

	// var hashStr []byte
	name := session.Get("name")
	password := session.Get("password")

	if name == nil || password == nil {
		c.Redirect(302, "/login")
	}

	strname := name.(string)
	strpassword := password.(string)

	member, err = validation(strname, strpassword)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}
	return
}

func Login(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	session := sessions.Default(c)

	_, err := validation(name, password)

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

func validation(name string, password string) (member int, err error) {
	db, err := infra.DBConnect()

	if err != nil {
		return
	}

	var userauth domain.LoginInfo

	db.Table("users").Select("id, name, password").Where("name = ?", name).Find(&userauth)
	selectpass := userauth.Password

	err = bcrypt.CompareHashAndPassword(selectpass, []byte(password))

	if err != nil {
		return
	}

	member = userauth.UserID
	return

}

func Register(c *gin.Context) {
	session := sessions.Default(c)

	var sign domain.LoginInfo

	if err := c.Bind(&sign); err != nil {
		c.JSON(500, gin.H{"err": err})
		c.Abort()
	} else {
		name := c.PostForm("name")
		password := c.PostForm("password")

		if err := signUp(name, password); err != nil {
			c.JSON(500, gin.H{"err": err})
		}

		session.Set("name", name)
		session.Set("password", password)
		session.Save()

		c.Redirect(302, "/success")

	}
}

func signUp(user string, password string) (err error) {

	gormdb, err := infra.DBConnect()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	fmt.Println(hash)

	if err != nil {
		return err
	}

	newuser := table.User{
		Name:        user,
		Password:    hash,
		MailAddress: nil,
		HN:          nil,
		Img:         nil,
		FinalGoal:   nil,
		Profile:     nil,
		Twitter:     nil,
		Instagram:   nil,
		Facebook:    nil,
		Github:      nil,
		URL:         nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}

	err = gormdb.Create(&newuser).Error

	if err != nil {
		return err
	}

	return nil
}
