package ui

import (
	stc "strconv"

	"../app/admin"
	"../app/goal"
	"github.com/gin-gonic/gin"
)

func GetGoal(c *gin.Context) {
	page, _ := stc.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := stc.Atoi(c.DefaultQuery("limit", "100"))
	order := c.DefaultQuery("order", "last_achieved")

	res, err := goal.ToGetAllGoal(limit, page, order)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, gin.H{
		"GoalArray": res,
		"limit":     limit,
		"page":      page,
		"order":     order,
	})

}

func GetOneUserGoal(c *gin.Context) {

	_, user, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		c.Abort()
	}

	name := c.Param("name")
	order := c.DefaultQuery("order", "last_achieved")

	res, err := goal.ToGetOneGoal(name, order)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, gin.H{
		"Goal":  res,
		"order": order,
		"owner": admin.JudgeOwner(user, name),
	})

}
