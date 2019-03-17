package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klovack/traversy-rest/pkg/todo"
)

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock Data - @todo -implement DB
	todos := []todo.Todo{}
	todos = append(
		todos,
		todo.Todo{
			ID:          "1",
			Title:       "Finish Tutorial",
			Description: "Hurry up and finish the tutorial",
			IsDone:      false,
			Author: &todo.Author{
				Firstname: "Rizki",
				Lastname:  "Fikriansyah",
			},
		})
	todos = append(
		todos,
		todo.Todo{
			ID:          "2",
			Title:       "Buy some food",
			Description: "I'm hungry and we don't have any food, we should go to supermarket",
			IsDone:      false,
			Author: &todo.Author{
				Firstname: "Vitri",
				Lastname:  "Indriyani",
			},
		})

	todoList := todo.NewTodoList(&todos)

	// Route Handlers / Endpoints
	router.HandleFunc("/api/todos", todoList.GetTodos).Methods("GET")
	router.HandleFunc("/api/todos/{id}", todoList.GetTodo).Methods("GET")
	router.HandleFunc("/api/todos", todoList.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", todoList.UpdateTodo).Methods("PUT")
	router.HandleFunc("/api/todos/{id}", todoList.DeleteTodo).Methods("DELETE")

	port := 8000
	fmt.Printf("Starting application on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}
