package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-api/models"
)

var taskList = []models.Task{}
var idTask int = 0

func CreateTaskHandler(w http.ResponseWriter, r *http.Request){
	idTask += 1

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
		ID: idTask, 
		Title: taskResponseRequest.Title,
		Description: taskResponseRequest.Description,
		Done: false}
	taskList = append(taskList, tempTask)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(tempTask)
	
	fmt.Println(taskList[idTask - 1])
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(taskList)
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

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	var found bool = false
	for i, task := range taskList {
		if task.ID == id {
			taskList = append(taskList[:i], taskList[i+1:]...)
			found = true
			break
		}
	}
	
	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
	}	
}