package main

import (
	"github.com/gin-gonic/gin"
	"study-todo-backend/src/db"
	"study-todo-backend/src/sys"
)

func main() {
	db.SetupDatabase()
	r := gin.Default()
	sys.SetupRouteTable(r)
	//r.GET("/ping", func(c *gin.Context) { //post, get, put 등등 , 라우터를 만든다, 컨트롤러
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})

	//r.GET("/todos", todoCtrl.GetTodos)
	//todos := r.Group("/todos")
	//{
	//	//todos.GET("/:id", todoCtrl.GetTodo)
	//	//todos.GET("/insert/:id/:title", todoCtrl.InsertTodo)
	//	//todos.GET("/update/:id/:title", todoCtrl.UpdateTodo)
	//	//todos.GET("/delete/:id", todoCtrl.DeleteTodo)
	//}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
