package main

import (
	"net/http"
	"todo-api/handlers"
)

func main(){

	http.HandleFunc("/createTask", handlers.CreateTaskHandler)
	http.HandleFunc("/getTasksByTitle", handlers.GetTaskByTitleHandler)
	http.HandleFunc("/deleteTask", handlers.DeleteTaskHandler)
	http.HandleFunc("/updateTask", handlers.UpdateTaskHandler)
	http.HandleFunc("/migrations", handlers.HandleMigrations)

	http.ListenAndServe(":8080", nil)
}