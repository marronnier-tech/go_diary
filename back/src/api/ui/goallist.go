package ui

import (
	stc "strconv"

	"../app/todo"
	"github.com/gin-gonic/gin"
)

func GetGoal(c *gin.Context) {
	page, _ := stc.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := stc.Atoi(c.DefaultQuery("limit", "100"))
	order := c.DefaultQuery("order", "last_achieved")

	res, err := todo.ToGetAllGoal(limit, page, order)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, gin.H{
		"TodoArray": res,
		"limit":     limit,
		"page":      page,
		"order":     order,
	})

}

func GetOneUserGoal(c *gin.Context) {
	name := c.Param("name")
	order := c.DefaultQuery("order", "last_achieved")

	res, err := todo.ToGetOneGoal(name, order)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, gin.H{
		"Todo":  res,
		"order": order,
	})

}
