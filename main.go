package main

import (
	"net/http"
	"todo-api/handlers"
)

func main(){
	

	http.HandleFunc("/createTask", handlers.CreateTaskHandler)
	http.HandleFunc("/getTasks", handlers.GetTaskHandler)
	http.HandleFunc("/deleteTask", handlers.DeleteTaskHandler)

	http.ListenAndServe(":8080", nil)
}