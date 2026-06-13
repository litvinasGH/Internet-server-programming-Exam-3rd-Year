package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Params struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Request struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
	ID      *int   `json:"id,omitempty"`
}

type Response struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  int    `json:"result"`
	ID      int    `json:"id"`
}

func processRequest(req Request) *Response {
	var result int

	switch req.Method {
	case "sum":
		result = req.Params.A + req.Params.B
	case "mul":
		result = req.Params.A * req.Params.B
	case "div":
		result = req.Params.A / req.Params.B
	default:
		return nil
	}

	log.Printf(
		"Method: %s, A: %d, B: %d, Result: %d",
		req.Method,
		req.Params.A,
		req.Params.B,
		result,
	)

	if req.ID == nil {
		return nil
	}

	return &Response{
		Jsonrpc: "2.0",
		Result:  result,
		ID:      *req.ID,
	}
}

func rpcHandler(w http.ResponseWriter, r *http.Request) {

	var requests []Request

	err := json.NewDecoder(r.Body).Decode(&requests)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var responses []Response

	for _, req := range requests {
		resp := processRequest(req)

		if resp != nil {
			responses = append(responses, *resp)
		}
	}

	if len(responses) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/rpc", rpcHandler).Methods("POST")

	log.Println("JSON-RPC server started on :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
