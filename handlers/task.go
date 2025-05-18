package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
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