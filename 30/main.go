package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	file := "test.txt"

	w.Header().Set("Content-Disposition", "attachment; filename=test.txt")
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, file)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/download", downloadHandler).Methods("GET")

	log.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
