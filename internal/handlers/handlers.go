package handlers

import (
	"cloud_assignment/internal/models"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	tasks   = make(map[string]models.Task)
	tasksMu sync.Mutex
)

// HelloWorld handlers GET / - returns {"message": "Hellow World"}
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello World"})
}

// CreateTask handles POST /tasks - creates a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if task.ID == "" {
		task.ID = uuid.New().String()
	}

	tasksMu.Lock()
	tasks[task.ID] = task
	tasksMu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// ListTasks handles GET /tasks - return all tasks
func ListTasks(w http.ResponseWriter, r *http.Request) {
	tasksMu.Lock()
	defer tasksMu.Unlock()
	var taskList []models.Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskList)
}

// GetTask handles GET /tasks/{id} - returns a specific task
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	tasksMu.Lock()
	task, exists := tasks[id]
	tasksMu.Unlock()

	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// HealthCheck handles GET /health - return {"status": "active"}
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "active"})
}

// UpdateTask handles PUT /tasks/{id} - updates a task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasksMu.Lock()
	defer tasksMu.Unlock()

	if _, exists := tasks[id]; !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	updatedTask.ID = id
	tasks[id] = updatedTask

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTask handles DELETE /tasks/{id} - deletes a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	tasksMu.Lock()
	defer tasksMu.Unlock()

	if _, exists := tasks[id]; !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	delete(tasks, id)
	w.WriteHeader(http.StatusNoContent)
}
