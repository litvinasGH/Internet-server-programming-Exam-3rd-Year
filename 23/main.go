package main

import (
	"fmt"
	"log"
	"net/http"
)

func Ahandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(
		"RequestA: %s %s",
		r.Method,
		r.URL.Path,
	)
	fmt.Fprint(w, "RESPONSE A!")
}

func Bhandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(
		"RequestB: %s %s",
		r.Method,
		r.URL.Path,
	)
	fmt.Fprint(w, "RESPONSE B!")
}

func main() {
	port := ":8080"
	http.HandleFunc("/A", Ahandler)
	http.HandleFunc("/B", Bhandler)

	log.Print("Server started on " + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error:", err)
	}
}
