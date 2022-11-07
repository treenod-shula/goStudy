package todoCtrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study-todo-backend/src/models/todoModel"
	"study-todo-backend/src/service/todoSvc"
)

func GetTodos(c *gin.Context) { //post, get, put 등등 , 라우터를 만든다
	todos := todoModel.GetTotods()
	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	data := todoSvc.GetTodo(id)
	c.JSON(http.StatusOK, gin.H{
		"todos": data,
	})
}
