package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	log.Println("User ID:", id)

	fmt.Fprintf(w, "User ID: %s", id)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user/{id}", userHandler)

	log.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
