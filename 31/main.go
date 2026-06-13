package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) //10МБ(10*2^20)Побитовый сдвиг
	if err != nil {
		http.Error(w, "Form error", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	age := r.FormValue("age")

	log.Printf("Name: %s, Age: %s", name, age)

	fmt.Fprintf(w, "Name: %s, Age: %s", name, age)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/data", uploadHandler).Methods("POST")

	log.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
