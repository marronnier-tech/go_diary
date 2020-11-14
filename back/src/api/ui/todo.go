package ui

import (
	"fmt"
	stc "strconv"

	"../app/todo"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {

	res, err := todo.ToGetAll()

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, res)

	return

}

func PostTodo(c *gin.Context) {

	id, _ := stc.Atoi(c.PostForm("id"))
	user, _ := stc.Atoi(c.PostForm("user"))
	content := c.PostForm("content")

	err := todo.ToPost(id, user, content)

	if err != nil {
		errHundle(err, 500, c)
	}

	c.JSON(201, nil)

	return

}

func DeleteTodo(c *gin.Context) {

	id, _ := stc.Atoi(c.Param("id"))

	err := todo.ToDelete(id)

	if err != nil {
		errHundle(err, 500, c)
	}

	c.JSON(201, nil)

	return
}
