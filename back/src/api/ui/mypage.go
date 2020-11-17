package ui

import (
	stc "strconv"

	"../app/todo"

	"github.com/gin-gonic/gin"
)

func PostTodo(c *gin.Context) {

	user, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	content := c.PostForm("content")

	err = todo.ToPost(user, content)

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(201, nil)

	return

}

func DeleteTodo(c *gin.Context) {

	id, _ := stc.Atoi(c.Param("id"))

	err := todo.ToDelete(id)

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(201, nil)

	return
}

func PutAchieveTodo(c *gin.Context) {

	user, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	id, _ := stc.Atoi(c.Param("id"))

	res, err := todo.ToPutAchieve(id, user)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, res)

}

func ClearAchieveTodo(c *gin.Context) {
	id, _ := stc.Atoi(c.Param("id"))

	res, err := todo.ToClearAchieve(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, res)
}

func PatchGoal(c *gin.Context) {
	id, _ := stc.Atoi(c.Param("id"))

	err := todo.ToPatchGoal(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(201, nil)

}
