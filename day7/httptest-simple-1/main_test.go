package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)


func TestRequestHandler(t *testing.T) {
	expected := "Hello bob"
	req := httptest.NewRequest(http.MethodGet, "/greet?name=bob", nil)
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected Hello bob but got %v", string(data))
	}

}
