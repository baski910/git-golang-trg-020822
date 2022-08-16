package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetEntries(t *testing.T) {
	req, err := http.NewRequest("GET", "/entries", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	m := Router()
	m.HandleFunc("/entries", GetEntries)
	m.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong code: got %v want %v", status, http.StatusOK)
	}
	expected := `[{"id":1,"first_name":"bob","last_name":"tom","email_address":"bob@example.com","phone_number":"9875023123"},{"id":2,"first_name":"alice","last_name":"tom","email_address":"alice@example.com","phone_number":"4567891232"}]`
	result := rr.Body.String()
	result = strings.TrimSpace(result)
	if result != expected {
		t.Errorf("handler returned unexpected results: got %v want %v", result, expected)
	}
}

func TestGenEntryById(t *testing.T) {
	req, err := http.NewRequest("GET", "/entry/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	/*
		q := req.URL.Query()
		q.Add("id", "1")
		req.URL.RawQuery = q.Encode()
	*/
	rr := httptest.NewRecorder()
	m := Router()
	m.HandleFunc("/entry/{id}", GetEntryById)
	m.ServeHTTP(rr, req)
	//handler := http.HandlerFunc(GetEntryById)
	//handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong code got %v want %v", status, http.StatusOK)
	}
	expected := `{"id":1,"first_name":"bob","last_name":"tom","email_address":"bob@example.com","phone_number":"9875023123"}`
	result := strings.TrimSpace(rr.Body.String())

	if result != expected {
		t.Errorf("handler returned unexpected body got %v want %v", result, expected)
	}
}

func TestCreatEntry(t *testing.T) {
	var jsonStr = []byte(`{"id":11,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`)
	req, err := http.NewRequest("POST", "/entry", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	m := Router()
	m.HandleFunc("/entry", CreateEntry)
	m.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong code got %v want %v", status, http.StatusOK)
	}

}
