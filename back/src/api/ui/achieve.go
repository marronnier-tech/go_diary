package ui

import (
	stc "strconv"

	"../app/todo"

	"github.com/gin-gonic/gin"
)

func PutAchieveTodo(c *gin.Context) {
	id, _ := stc.Atoi(c.Param("id"))

	res, err := todo.ToPutAchieve(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
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
