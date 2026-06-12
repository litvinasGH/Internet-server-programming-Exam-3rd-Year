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

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("User: %s, %d years", user.Name, user.Age)

	w.Header().Set("Content-Type", "application/json")

	response := map[string]any{
		"status": "success",
		"user":   user,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	port := ":8080"
	router := mux.NewRouter()

	router.HandleFunc("/user", createUser).Methods("POST")

	log.Println("Server started on " + port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
