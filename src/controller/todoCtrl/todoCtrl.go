package todoCtrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study-todo-backend/src/models/todoModel"
	"study-todo-backend/src/service/todoSvc"
)

const RET_OK = 1

func GetTodos(c *gin.Context) {
	userId := 1
	todos, err := todoSvc.GetTodos(userId)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"ret":   RET_OK,
		"todos": todos,
	})
}

// 포스트 받을때 쉼표로 받음
func AddTodo(c *gin.Context) {
	userId := 1
	title := "todoTest"
	err := todoModel.AddTodo(userId, title)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"ret": RET_OK,
	})
}

//func GetTodo(c *gin.Context) {
//	id := c.Param("id")
//	todos := todoModel.GetTodo(id)
//	//data := todoSvc.GetTodo(id)
//	c.JSON(http.StatusOK, gin.H{
//		"todos": todos,
//	})
//}
//func InsertTodo(c *gin.Context) {
//	id := c.Param("id")
//	title := c.Param("title")
//	todos := todoModel.InsertTodo(id, title)
//	//data := todoSvc.GetTodo(id)
//	c.JSON(http.StatusOK, gin.H{
//		"todos": todos,
//	})
//}
//
//func UpdateTodo(c *gin.Context) {
//	id := c.Param("id")
//	title := c.Param("title")
//	todos, err := todoModel.UpdateTodo(id, title)
//	if err != nil {
//		_ = c.AbortWithError(http.StatusBadRequest, err)
//		return
//	}
//	//data := todoSvc.GetTodo(id)
//	c.JSON(http.StatusOK, gin.H{
//		"todos": todos,
//	})
//}
//
//func DeleteTodo(c *gin.Context) {
//	id := c.Param("id")
//	todos, err := todoModel.DeleteTodo(id)
//	if err != nil {
//		_ = c.AbortWithError(http.StatusBadRequest, err)
//		return
//	}
//	//data := todoSvc.GetTodo(id)
//	c.JSON(http.StatusOK, gin.H{
//		"todos": todos,
//	})
//}
