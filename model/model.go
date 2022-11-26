package model

import (
	"time"
)

type Todo struct {
	Id        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Completed bool      `json:"completed" gorm:"column:completed"`
	CreateAt  time.Time `json:"create_at" gorm:"column:create_at"`
}

var todoMap map[int]*Todo

func init() {
	todoMap = make(map[int]*Todo)
}

func GetTodos() []*Todo {
	var list []*Todo
	for _, v := range todoMap {
		//log.Println("v:", v)
		list = append(list, v)
	}
	return list
}

func AddTodo(name string) *Todo {
	id := len(todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	todoMap[id] = todo

	return todo
}

func RemoveTodo(id int) bool {
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		return true
	}

	return false
}

func CompleteTodo(id int, complete bool) bool {
	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		return true
	}

	return false
}
