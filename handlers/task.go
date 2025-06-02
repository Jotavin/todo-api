package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-api/models"
)

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
	models.TaskList = append(models.TaskList, tempTask)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(tempTask)
	
	fmt.Println(idTask)
	fmt.Println(len(models.TaskList) - 1)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	id, err := strconv.Atoi(idParam)
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
	
	idTask -= 1
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	type data struct {
		ID int	`json:"id"`
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