package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	log.Printf("Name: %s, Age: %s", name, age)

	fmt.Fprintf(w, "Name: %s, Age: %s", name, age)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user", userHandler).Methods("GET")

	log.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
