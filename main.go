package main

import (
	"net/http"
	"todo-api/handlers"
	"todo-api/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main(){

	http.HandleFunc("/createTask", handlers.CreateTaskHandler)
	http.HandleFunc("/getTasksByTitle", handlers.GetTaskByTitleHandler)
	http.HandleFunc("/deleteTask", handlers.DeleteTaskHandler)
	http.HandleFunc("/updateTask", handlers.UpdateTaskHandler)
	http.HandleFunc("/migrations", handlers.HandleMigrations)

	http.Handle("/metrics", promhttp.Handler())

	metrics.RecordMetrics()

	http.ListenAndServe(":8080", nil)
}