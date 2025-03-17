package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Hello, %s", task)
	fmt.Fprintln(w, response)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task = reqBody.Task

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Task updated successfully")
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", PostTaskHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}
