package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"cloud_assignment/internal/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", handlers.ListTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Println("Starting server on: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}