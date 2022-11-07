package todoSvc

import (
	"strconv"
	"study-todo-backend/src/models/todoModel"
)

func GetTodo(data string) string {
	todos := todoModel.GetTotods()
	value, _ := strconv.Atoi(data)

	return todos[value]
}
