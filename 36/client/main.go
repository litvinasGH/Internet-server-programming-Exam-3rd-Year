package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const serverURL = "http://localhost:8080"

func main() {

	fmt.Println("=== CREATE COLLECTION ===")
	createCollection("docs")
	time.Sleep(2 * time.Second)

	fmt.Println("=== CREATE FILE ===")
	createFile("docs/test.txt", "Hello WebDAV")
	time.Sleep(2 * time.Second)

	fmt.Println("=== READ FILE ===")
	readFile("docs/test.txt")
	time.Sleep(2 * time.Second)

	fmt.Println("=== UPDATE FILE ===")
	createFile("docs/test.txt", "Updated content")
	time.Sleep(2 * time.Second)

	fmt.Println("=== READ UPDATED FILE ===")
	readFile("docs/test.txt")
	time.Sleep(2 * time.Second)

	fmt.Println("=== PROPFIND ===")
	propfind("docs")
	time.Sleep(2 * time.Second)

	fmt.Println("=== DELETE FILE ===")
	deleteResource("docs/test.txt")

	time.Sleep(2 * time.Second)
	fmt.Println("=== DELETE COLLECTION ===")
	deleteResource("docs")
}

func createCollection(name string) {
	req, err := http.NewRequest(
		"MKCOL",
		serverURL+"/"+name,
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}

func createFile(path string, content string) {
	req, err := http.NewRequest(
		"PUT",
		serverURL+"/"+path,
		strings.NewReader(content),
	)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}

func readFile(path string) {
	resp, err := http.Get(serverURL + "/" + path)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	fmt.Println("Status:", resp.Status)
	fmt.Println("Content:", string(data))
}

func deleteResource(path string) {
	req, err := http.NewRequest(
		"DELETE",
		serverURL+"/"+path,
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}

func propfind(path string) {
	req, err := http.NewRequest(
		"PROPFIND",
		serverURL+"/"+path,
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	fmt.Println("Status:", resp.Status)
	fmt.Println(string(data))
}
