package app

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"jhpark.sinsiway.com/bootstrap/model"
	"net/http"
	"strconv"
)

var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	//var list []*model.Todo
	//for _, v := range todoMap {
	//	log.Println("v:", v)
	//	list = append(list, v)
	//}
	list := model.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func addTodoListHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := model.AddTodo(name)

	//id := len(todoMap) + 1
	//todo := &Todo{id, name, false, time.Now()}
	//todoMap[id] = todo

	rd.JSON(w, http.StatusOK, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id, _ := strconv.Atoi(vars["id"])

	id, _ := strconv.Atoi(r.FormValue("id"))
	ok := model.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
	//if _, ok := todoMap[id]; ok {
	//	delete(todoMap, id)
	//	rd.JSON(w, http.StatusOK, Success{true})
	//	log.Println("deleted id:", id)
	//} else {
	//	rd.JSON(w, http.StatusOK, Success{false})
	//	log.Println("not found id:", id)
	//}

}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"

	ok := model.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}

	//if todo, ok := todoMap[id]; ok {
	//	todo.Completed = complete
	//	rd.JSON(w, http.StatusOK, Success{true})
	//} else {
	//	rd.JSON(w, http.StatusOK, Success{false})
	//}
}

//func addTestTodo() {
//	todoMap[1] = &Todo{1, "Init Data1", false, time.Now()}
//	todoMap[2] = &Todo{2, "Init Data2", true, time.Now()}
//	todoMap[3] = &Todo{3, "Init Data3", false, time.Now()}
//}

func MakeHandler() http.Handler {
	//todoMap = make(map[int]*Todo)
	//addTestTodo()
	rd = render.New()

	r := mux.NewRouter()

	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", addTodoListHandler).Methods("POST")
	r.HandleFunc("/todos", removeTodoHandler).Methods("DELETE")
	//r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoHandler).Methods("GET")

	r.HandleFunc("/", indexHandler)
	return r
}
