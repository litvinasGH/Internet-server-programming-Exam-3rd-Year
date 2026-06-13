package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	log.Println("Search:", query)

	fmt.Fprintf(w, "Search: %s", query)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/search", searchHandler).
		Methods("GET").
		Queries("q", "{q}")

	log.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
