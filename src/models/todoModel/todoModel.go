package todoModel

import (
	"study-todo-backend/src/db"
)

//var dsn string = "treenod:xmflshem@tcp(ec2-13-125-229-201.ap-northeast-2.compute.amazonaws.com:3306)/study"
//var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// id(int), title(varchar), completed(BOOL),userId(smallint)
// 소문자는 전역변수가 안됨
type Todo struct {
	Id        int    `json:"id" gorm:"column:id"`
	Title     string `json:"title" gorm:"column:title"`
	Completed bool   `json:"completed" gorm:"column:completed"`
	Userid    int    `json:"userid" gorm:"column:userId"`
}

//func GetTodo(id string) []Todo {
//	if err != nil {
//		panic("failed to connect database")
//	}
//	todos := []Todo{}
//	value, _ := strconv.Atoi(id)
//	db.Where("userId =?", value).Table("todos").Find(&todos)
//	fmt.Println(todos)
//
//	return todos
//}
//
//func InsertTodo(id string, title string) []Todo {
//	if err != nil {
//		panic("failed to connect database")
//	}
//	var todos []Todo
//	UserId, _ := strconv.Atoi(id)
//	newTodo := Todo{USERID: UserId, TITLE: title, COMPLETED: false}
//	db.Table("todos").Create(&newTodo)
//	db.Where("userId =?", id).Table("todos").Find(&todos)
//	return todos
//}
//
//func UpdateTodo(id string, title string) ([]Todo, error) {
//	if err != nil {
//		panic("failed to connect database")
//	}
//	var todos []Todo
//	var comp bool
//	//구분값이 명확하지 않아서 자꾸 INSERT 된거임
//	Id, _ := strconv.Atoi(id)
//	db.Where("id =?", Id).Table("todos").Find(&todos)
//
//	if todos[0].COMPLETED == false {
//		comp = true
//	} else {
//		comp = false
//	}
//	compTodo := Todo{ID: Id, TITLE: title, COMPLETED: comp}
//	if saveErr := db.Table("todos").Save(&compTodo).Error; saveErr != nil {
//		return nil, saveErr
//	}
//
//	return todos, nil
//}
//
//func DeleteTodo(id string) ([]Todo, error) {
//	Id, _ := strconv.Atoi(id)
//
//	var todos []Todo
//	delTodo := Todo{ID: Id}
//
//	if delErr := db.Table("todos").Delete(&delTodo).Error; delErr != nil {
//		return nil, delErr
//	}
//
//	db.Table("todos").Find(&todos)
//	return todos, nil
//
//}

func GetTodos(userId int) ([]Todo, error) {
	var todos []Todo
	err := db.GetORM().Table("todos").Where("userId=?", userId).Find(&todos).Error
	return todos, err
}

func AddTodo(id int, title string) error {

	newTodo := Todo{Userid: id, Title: title, Completed: false}
	err := db.GetORM().Table("todos").Create(&newTodo).Error

	return err
}
