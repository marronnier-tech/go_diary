package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	stc "strconv"
	"time"

	"./infra"
	"./infra/model"

	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

var secrets = gin.H{
	"foo":     gin.H{"email": "foo@bar.com", "phone": "123"},
	"austion": gin.H{"email": "austin@example.com", "phone": "666"},
	"nyao":    gin.H{"email": "nyao@mails.com", "phone": "54232"},
}

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

func Encrypt(key, pass []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, pass, nil)

	return ciphertext, nil

}

func readKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	file, err := os.Open("key.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	for {
		read, err := file.Read(key)
		if read == 0 {
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return key, nil
}

func signUp(id string, user string, mailaddress string, pass string) (err error) {

	sqldb, gormdb := infra.DBConnect()
	defer sqldb.Close()

	newid, _ := stc.Atoi(id)

	passbyte := []byte(pass)

	key, err := readKey()
	if err != nil {
		return err
	}

	cipherpass, err := Encrypt(key, passbyte)
	if err != nil {
		return err
	}

	fmt.Printf("ciphertext: %s\n", hex.EncodeToString(cipherpass))
	//passEncrypt, _ := crypto.PasswordEncrypt(pass)

	newuser := model.User{
		ID:          newid,
		Name:        user,
		MailAddress: mailaddress,
		Pass:        cipherpass,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = gormdb.Create(&newuser).Error

	if err != nil {
		return err
	}

	return nil
}
