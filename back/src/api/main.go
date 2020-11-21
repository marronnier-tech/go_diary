package main

import (
	"./ui"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("../../../front/templates/*")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("useradmin", store))

	todolist := r.Group("/todo")
	{
		todolist.GET("", ui.GetTodo)
		todolist.GET("/:name", ui.GetOneUserTodo)
	}

	goallist := r.Group("/goal")
	{
		goallist.GET("", ui.GetGoal)
		goallist.GET("/:name", ui.GetOneUserGoal)
	}

	my := r.Group("/mypage")
	{
		my.GET("", ui.MyTodo)
		my.POST("", ui.PostTodo)
		my.DELETE("/:id", ui.DeleteTodo)

		my.POST("/:id/today", ui.PutAchieveTodo)
		my.DELETE("/:id/today", ui.ClearAchieveTodo)
		my.PATCH("/:id/goal", ui.PatchGoal)

	}

	profile := r.Group("/profile")
	{
		profile.GET("", ui.GetMyProfile)
		profile.PATCH("", ui.PatchProfile)

		profile.GET("/:name", ui.GetOnesProfile)

	}

	// ログイン、あとでsuccess変更

	r.POST("/register", ui.Register)

	r.POST("/login", ui.Login)
	r.DELETE("/logout", ui.Logout)

	r.DELETE("/delete", ui.DeleteMembership)

	r.GET("/success", func(c *gin.Context) {
		c.JSON(201, gin.H{"message": "success!"})
	})

	// ここでDBとじるのは問題なのであとでなんとかする

	/* sqldb, err := gormdb.DB()

	if err != nil {
		fmt.Println("cannot use sqldb.")
	}

	sqldb.Close() */

	r.Run()

}
