package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-api/models"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	taskResponseRequest := models.Task{}
	err := json.NewDecoder(r.Body).Decode(&taskResponseRequest)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	
	var tempTask models.Task = models.Task{
		Title: taskResponseRequest.Title,
		Description: taskResponseRequest.Description,
		Done: false,
	}
	models.TaskList = append(models.TaskList, tempTask)
	db := models.ConnectDB()
	db.Create(&tempTask)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(tempTask)

}
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	task := models.Task{}
	db := models.ConnectDB()
	result := db.First(&task)
	fmt.Printf("Result get result %v", result)
	

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.TaskList)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	var found bool = false
	for i, task := range models.TaskList {
		if task.ID == id {
			models.TaskList = append(models.TaskList[:i], models.TaskList[i+1:]...)
			found = true
			break
		}
	}
	
	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
	}
	
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut {
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
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	
	if data_update.ID == 0 {
		http.Error(w, "Missing parameter id", http.StatusBadRequest)
		return
	}

	if data_update.Description == "" {
		http.Error(w, "Missing parameter description", http.StatusBadRequest)
		return
	}

	found := false
	for i, task := range models.TaskList {
		if task.ID == data_update.ID {
			models.TaskList[i].Description = data_update.Description
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}

func HandleMigrations(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	migration_result, err := models.MigrateDB()
	if err != nil {
		fmt.Printf("panic: %v\n", err)
	}

	json.NewEncoder(w).Encode(migration_result)

}