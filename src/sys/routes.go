package sys

import (
	"github.com/gin-gonic/gin"
	"study-todo-backend/src/controller/todoCtrl"
)

type RouteInfo struct {
	Path    string
	Handler func(ctx *gin.Context)
}

func SetupRouteTable(r *gin.Engine) {
	routePostInfos := []RouteInfo{
		{"/getTodos", todoCtrl.GetTodos},
		{"/addTodo", todoCtrl.AddTodo},
	}

	for _, routeInfo := range routePostInfos {
		r.POST("/pokopoko"+routeInfo.Path+".php", routeInfo.Handler)
	}
}
