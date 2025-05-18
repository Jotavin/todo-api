package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct{
	ID int
	Title string
	Description string
	Done bool
}

	var taskList = []Task{}
	var idTask int = 0

func main(){
	

	http.HandleFunc("/createTask", createTaskHandler)
	http.HandleFunc("/getTasks", getTaskHandler)

	http.ListenAndServe(":8080", nil)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request){
	idTask += 1

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var details struct{
		Title string `json:"title"`
		Description string `json:"description"`
	}

	err := json.NewDecoder(r.Body).Decode(&details)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	
	var tempTask Task = Task{idTask, details.Title, details.Description, false}
	taskList = append(taskList, tempTask)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(tempTask)
	
	fmt.Println(taskList[idTask - 1])
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(taskList)
}