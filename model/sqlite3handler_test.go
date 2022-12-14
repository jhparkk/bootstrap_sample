package model

import (
	"log"
	"testing"
)

func TestSqlite3Handler(t *testing.T) {
	dbhandler := newSqlite3Handler()
	todo := dbhandler.AddTodo("test")
	log.Println(todo)
}
