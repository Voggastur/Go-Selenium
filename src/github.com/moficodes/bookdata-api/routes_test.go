package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestCreateBook(t *testing.T) {

	var jsonStr = []byte(`{"book_id":2,"title":"Johans Polar Expedition","authors":"Jules Verne","isbn":2,"isbn_13":1000000000002,"language_code":"eng","num_pages":999,"ratings":25,"reviews":15}`)

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/book", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createBook)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"success": "created"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


func TestGetBookByISBN(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/book/isbn/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(searchByISBN)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"book_id":"2","title":"Johans Polar Expedition","authors":"Jules Verne","average_rating":1.6,"isbn":"2","isbn_13":"1000000000002","language_code":"eng","num_pages":999,"ratings":25,"reviews":15}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
