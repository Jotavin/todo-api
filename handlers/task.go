package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todo-api/metrics"
	"todo-api/models"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request){
	start := time.Now()
	defer func() {
		metrics.RequestDuration.WithLabelValues("/createTask", "POST").Observe(time.Since(start).Seconds())
	}()

	if r.Method != http.MethodPost {
		metrics.RequestsTotal.WithLabelValues("/createTask", "POST", "405").Inc()
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	taskResponseRequest := models.Task{}
	err := json.NewDecoder(r.Body).Decode(&taskResponseRequest)
	if err != nil {
		metrics.RequestsTotal.WithLabelValues("/createTask", "POST", "400").Inc()
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	
	var tempTask models.Task = models.Task{
		Title: taskResponseRequest.Title,
		Description: taskResponseRequest.Description,
		Done: false,
	}

	db, err := models.ConnectDB()
	if err != nil{
		metrics.RequestsTotal.WithLabelValues("/createTask", "POST", "500").Inc()
		metrics.DatabaseErrors.WithLabelValues("connect", "connection_failed").Inc()

		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	result := db.Create(&tempTask)
	if result.Error != nil {
		metrics.RequestsTotal.WithLabelValues("/createTask", "POST", "500").Inc()
		metrics.DatabaseErrors.WithLabelValues("create", "insert_failed").Inc()

		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	metrics.RequestsTotal.WithLabelValues("/createTask", "POST", "201").Inc()
	metrics.TasksCreated.Inc()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tempTask)

}

func GetTasksByTitleHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func(){
		metrics.RequestDuration.WithLabelValues("/getTasksByTitle", "GET").Observe(time.Since(start).Seconds())
	}()

	if r.Method != http.MethodGet {
		metrics.RequestsTotal.WithLabelValues("/getTasksByTitle", "GET", "405").Inc()
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks := []models.Task{}

	db, err := models.ConnectDB()
	if err != nil{
		metrics.RequestsTotal.WithLabelValues("/getTasksByTitle", "GET", "500").Inc()
		metrics.DatabaseErrors.WithLabelValues("connect", "connection_failed").Inc()

		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	title := r.URL.Query().Get("title")
	searchPattern := "%"
	if title != "" {
		searchPattern = "%" + title + "%"
	}

	result := db.Where("title LIKE ?", searchPattern).Find(&tasks)
	if result.Error != nil{
		metrics.RequestsTotal.WithLabelValues("/getTasksByTitle", "GET", "503").Inc()
		metrics.DatabaseErrors.WithLabelValues("query", "find_failed").Inc()

		http.Error(w, "Error quering database", http.StatusServiceUnavailable)
		return
	}

	metrics.RequestsTotal.WithLabelValues("/getTasksByTitle", "GET", "200").Inc()
	metrics.TasksQueried.Inc()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		metrics.RequestDuration.WithLabelValues("/deleteTask", "DELETE").Observe(time.Since(start).Seconds())
	}()


	if r.Method != http.MethodDelete{
		metrics.RequestsTotal.WithLabelValues("/deleteTask", "DELETE", "405").Inc()

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		metrics.RequestsTotal.WithLabelValues("/deleteTask", "DELETE", "400").Inc()

		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		metrics.RequestsTotal.WithLabelValues("/deleteTask", "DELETE", "400").Inc()
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	
	task := models.Task{}

	db, err := models.ConnectDB()
	if err != nil{
		metrics.RequestsTotal.WithLabelValues("/deleteTask", "DELETE", "500").Inc()
		metrics.DatabaseErrors.WithLabelValues("connect", "connection_failed").Inc()

		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	result := db.First(&task, id)
	if result.Error != nil {
		metrics.RequestsTotal.WithLabelValues("/deleteTask", "DELETE", "404").Inc()
		metrics.TasksNotFound.Inc()
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	deleteResult := db.Delete(&task, id)
	if deleteResult.Error != nil {
		metrics.RequestsTotal.WithLabelValues("/deleteTask", "DELETE", "500").Inc()
		metrics.DatabaseErrors.WithLabelValues("delete", "delete_failed").Inc()

		http.Error(w, "Error to delete", http.StatusInternalServerError)
		return
	}

	metrics.RequestsTotal.WithLabelValues("/deleteTask", "DELETE", "200").Inc()
	metrics.TasksDeleted.Inc()
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Task deleted")
	
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request){
	start := time.Now()
	defer func() {
		metrics.RequestDuration.WithLabelValues("/updateTask", "PUT").Observe(time.Since(start).Seconds())
	}()

	if r.Method != http.MethodPut {
		metrics.RequestsTotal.WithLabelValues("/updateTask", "PUT", "405").Inc()

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	type data struct {
		ID int64	`json:"id"`
		Description string `json:"description"`
	}
	data_update := data{}
	
	err := json.NewDecoder(r.Body).Decode(&data_update)
	if err != nil {
		metrics.RequestsTotal.WithLabelValues("/updateTask", "PUT", "400").Inc()

		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	
	if data_update.ID == 0 {
		metrics.RequestsTotal.WithLabelValues("/updateTask", "PUT", "400").Inc()

		http.Error(w, "Missing parameter id", http.StatusBadRequest)
		return
	}

	if data_update.Description == "" {
		metrics.RequestsTotal.WithLabelValues("/updateTask", "PUT", "400").Inc()

		http.Error(w, "Missing parameter description", http.StatusBadRequest)
		return
	}

	var task models.Task

	db, err := models.ConnectDB()
	if err != nil{
		metrics.RequestsTotal.WithLabelValues("/updateTask", "PUT", "500").Inc()
		metrics.DatabaseErrors.WithLabelValues("connect", "connection_failed").Inc()

		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	result := db.First(&task, data_update.ID)
	if result.Error != nil {
		metrics.RequestsTotal.WithLabelValues("/updateTask", "PUT", "404").Inc()
		metrics.TasksNotFound.Inc()

		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	
	// check later
	updates := map[string]interface{}{"description": data_update.Description}
	result = db.Model(&task).Where("id = ?", data_update.ID).Updates(updates)
	if result.Error != nil {
		metrics.RequestsTotal.WithLabelValues("/updateTask", "PUT", "500").Inc()
		metrics.DatabaseErrors.WithLabelValues("update", "update_failed").Inc()

		http.Error(w, "Operation update failed", http.StatusInternalServerError)
		return
	}

	metrics.RequestsTotal.WithLabelValues("/updateTask", "PUT", "200").Inc()
	metrics.TasksUpdated.Inc()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Column updated successfully")
}

func HandleMigrations(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		metrics.RequestDuration.WithLabelValues("/migrations", "GET").Observe(time.Since(start).Seconds())
	}()

	if r.Method != http.MethodGet {
		metrics.RequestsTotal.WithLabelValues("/migrations", "GET", "405").Inc()
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	migrationResult, err := models.MigrateDB()
	if err != nil {
		metrics.RequestsTotal.WithLabelValues("/migrations", "GET", "500").Inc()
		metrics.DatabaseErrors.WithLabelValues("migration", "migration_failed").Inc()
		http.Error(w, fmt.Sprintf("Migration failed: %v", err), http.StatusInternalServerError)
		return
	}

	metrics.RequestsTotal.WithLabelValues("/migrations", "GET", "200").Inc()
	metrics.MigrationsExecuted.Inc()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Migrations executed successfully",
		"result":  migrationResult,
	})

}