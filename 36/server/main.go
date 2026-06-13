package main

import (
	"log"
	"net/http"

	"golang.org/x/net/webdav"
)

func main() {
	handler := &webdav.Handler{
		Prefix:     "/",
		FileSystem: webdav.Dir("./storage"),
		LockSystem: webdav.NewMemLS(),
	}

	log.Println("WebDAV server started on :8080")

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
