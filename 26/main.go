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

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form error", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	age := r.FormValue("age")

	log.Printf("ID: %s, Name: %s, Age: %s", id, name, age)

	fmt.Fprintf(w, "ID: %s, Name: %s, Age: %s", id, name, age)
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
