package ui

import (
	stc "strconv"

	"../app/todo"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {

	lists, err := todo.ToGetAll()

	if err != nil {
		errHundle(err, 500, c)
	}

	c.JSON(200, gin.H{
		"lists": lists,
	})

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
