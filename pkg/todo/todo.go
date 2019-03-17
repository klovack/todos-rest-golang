package todo

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// RestAPI Interface
type RestAPI interface {
	GetTodos()
	GetTodo()
	CreateTodo()
	UpdateTodo()
	DeleteTodo()
}

// List of todos
type List struct {
	Todos *[]Todo
}

// Todo Struct (Model)
type Todo struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	IsDone      bool    `json:"isDone"`
	Author      *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// NewTodoList create list
func NewTodoList(todoList *[]Todo) *List {
	list := List{}
	list.Todos = todoList
	return &list
}

// GetTodos Get All Todos
func (l *List) GetTodos(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// Create new Json Encoder and encode todos
	json.NewEncoder(res).Encode(*l.Todos)
}

// GetTodo Get Single Todos
func (l *List) GetTodo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// Get Params
	params := mux.Vars(req)

	// Loop through todos and find with id
	for _, item := range *l.Todos {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}

	json.NewEncoder(res).Encode(&Todo{})
}

// CreateTodo Create new Todo
func (l *List) CreateTodo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var todo Todo
	_ = json.NewDecoder(req.Body).Decode(&todo)

	// Mock ID - not safe
	todo.ID = strconv.Itoa(rand.Intn(1000000))
	*l.Todos = append(*l.Todos, todo)

	json.NewEncoder(res).Encode(todo)
}

// UpdateTodo Update Todo
func (l *List) UpdateTodo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	todos := *l.Todos
	var updatedTodo Todo

	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)

			_ = json.NewDecoder(req.Body).Decode(&updatedTodo)

			// Mock ID - not safe
			updatedTodo.ID = item.ID
			todos = append(todos, updatedTodo)
			l.Todos = &todos
			json.NewEncoder(res).Encode(updatedTodo)
			return
		}
	}

	json.NewEncoder(res).Encode(Todo{})
}

// DeleteTodo Delete Todo
func (l *List) DeleteTodo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	todos := *l.Todos

	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}

	l.Todos = &todos

	json.NewEncoder(res).Encode(*l.Todos)
}
