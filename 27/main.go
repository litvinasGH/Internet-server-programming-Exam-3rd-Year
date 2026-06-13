package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST", http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("ID: %s, Name: %s, Age: %d",
		id, user.Name, user.Age)

	w.Header().Set("Content-Type", "application/json")

	response := map[string]any{
		"id":     id,
		"name":   user.Name,
		"age":    user.Age,
		"status": "success",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	port := ":8080"
	router := mux.NewRouter()

	router.HandleFunc("/user/{id}", userHandler).Methods("POST")

	log.Println("Server started on " + port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
