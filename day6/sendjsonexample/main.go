package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	//handler := http.HandlerFunc(handleRequest)
	http.HandleFunc("/example", handleRequest)
	log.Print("starting server at port 8000")
	http.ListenAndServe(":8000", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
