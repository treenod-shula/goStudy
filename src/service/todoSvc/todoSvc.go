package todoSvc

import (
	"study-todo-backend/src/models/todoModel"
)

// 파라미터 받아서 모델쪽에 접근하는것
func GetTodos(id int) ([]todoModel.Todo, error) {
	todos, err := todoModel.GetTodos(id)
	//value, _ := strconv.Atoi(data)

	return todos, err
}
