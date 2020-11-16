package ui

import (
	"fmt"
	stc "strconv"

	"../app/todo"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {

	page, _ := stc.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := stc.Atoi(c.DefaultQuery("limit", "100"))
	order := c.DefaultQuery("order", "last_achieved")

	res, err := todo.ToGetAll(limit, page, order)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"TodoArray": res,
		"limit":     limit,
		"page":      page,
		"order":     order,
	})

	return

}

func GetOneUserTodo(c *gin.Context) {

	name := c.Param("name")
	order := c.DefaultQuery("order", "last_achieved")

	res, err := todo.ToGetOneUser(name, order)

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, gin.H{
		"Todo":  res,
		"order": order,
	})

}

func PostTodo(c *gin.Context) {

	// id, _ := stc.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	content := c.PostForm("content")

	err := todo.ToPost(name, content)

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
