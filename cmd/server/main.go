package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHello)
	mux.HandleFunc("/health", handleHealth)

	log.Println("Server running on: 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Hello World",
	}
	writeJSON(w, http.StatusOK, response)
}
func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "active",
	}
	writeJSON(w, http.StatusOK, response)
}
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
