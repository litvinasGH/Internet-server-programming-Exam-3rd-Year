package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(
		"Request: %s %s\n",
		r.Method,
		r.URL.Path,
	)

	fmt.Fprint(w, "RESPONSE!")
}

func main() {
	port := ":8080"
	http.HandleFunc("/", handler)

	fmt.Println("Server started on " + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
