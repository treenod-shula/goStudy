package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study-todo-backend/src/controller/todoCtrl"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { //post, get, put 등등 , 라우터를 만든다, 컨트롤러
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/todos", todoCtrl.GetTodos)

	//r.GET("/todos/:id", func(c *gin.Context) { //post, get, put 등등 , 라우터를 만든다
	//	id := c.Param("id")
	//	c.JSON(http.StatusOK, gin.H{
	//		"id": id,
	//	})
	//})
	////0 -> 1 , 2-> 3

	r.GET("/todos/:id", todoCtrl.GetTodo)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
