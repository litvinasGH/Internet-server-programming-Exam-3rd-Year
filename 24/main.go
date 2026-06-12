package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
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

	log.Printf("User: %s, %d лет\n", user.Name, user.Age)

	w.Header().Set("Content-Type", "application/json")

	response := map[string]any{
		"status": "success",
		"user":   user,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	port := ":8080"
	http.HandleFunc("/", handler)

	log.Print("erver started on " + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error:", err)
	}
}
