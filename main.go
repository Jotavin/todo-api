package main

import (
	"net/http"
	"todo-api/handlers"
	"todo-api/models"
)

func main(){
	db := models.ConnectDB()
	models.MigrateDB(db)

	http.HandleFunc("/createTask", handlers.CreateTaskHandler)
	http.HandleFunc("/getTasks", handlers.GetTaskHandler)
	http.HandleFunc("/deleteTask", handlers.DeleteTaskHandler)
	http.HandleFunc("/updateTask", handlers.UpdateTaskHandler)

	http.ListenAndServe(":8080", nil)
}