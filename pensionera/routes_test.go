package main

import (
    "testing"
    "net/http"
    "io/ioutil"
    "strings"
)


func TestSearchBookByISBN(t *testing.T) {


    url := "http://localhost:8080/api/v1/book/isbn"
    isbn := "/074754624X"
    method := "GET"
    client := &http.Client {}

    req, err := http.NewRequest(method, url+isbn, nil)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    res, err := client.Do(req)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    if res.StatusCode <= 200 && res.StatusCode >= 299 {
	    t.Errorf("FAIL, HTTP Status is outside of the 2xx range: Got %v", res.StatusCode)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    expected := `{"book_id":"43509","title":"Harry Potter and the Goblet of Fire (Harry Potter  #4)","authors":"J.K. Rowling","average_rating":4.55,"isbn":"074754624X","isbn_13":"9780747546245","language_code":"eng","num_pages":636,"ratings":18097,"reviews":860}`
    if string(body) != expected {
	    t.Errorf("Fail: The found book does not match expected book")
    }
}


func TestSearchByAuthor(t *testing.T) {

    req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/books/authors/galileo%20galilei", nil)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    client := &http.Client {}
    res, err := client.Do(req)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    if res.StatusCode <= 200 && res.StatusCode >= 299 {
	    t.Errorf("Fail: HTTP Response outside 2xx range, Got: %v", res.StatusCode)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    expected := `[{"book_id":"32610","title":"On the Shoulders of Giants: The Great Works of Physics and Astronomy","authors":"Stephen Hawking-Isaac Newton-Nicolaus Copernicus-Albert Einstein-Johannes Kepler-Galileo Galilei","average_rating":4.2,"isbn":"0762427329","isbn_13":"9780762427321","language_code":"eng","num_pages":256,"ratings":4014,"reviews":34}]`
    if string(body) != expected {
	    t.Errorf("Fail: Does not match expected book")
    }
}


func TestCreateBook(t *testing.T) {

  url := "http://localhost:8080/api/v1/book"
  method := "POST"
  payload := strings.NewReader("BookID=2&Title=My%20Marmelade%20Tips&Authors=Jules%20Verne&ISBN=2&ISBN13=1000000000002&Ratings=25&Reviews=15&NumPages=999&LanguageCode=eng&AverageRating=1.6")
  client := &http.Client {}

  req, err := http.NewRequest(method, url, payload)
  if err != nil {
	  t.Fatalf(err.Error())
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  res, err := client.Do(req)
  if err != nil {
	  t.Fatalf(err.Error())
  }
  if res.StatusCode != 201 {
	  t.Errorf("HTTP Status not what expected: got %v", res.StatusCode)
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
	  t.Fatalf(err.Error())
  }
  if string(body) != `{"success": "created"}` {
	  t.Errorf("FAIL: Success created not found, got: %v", string(body))
  }
}


func TestReadRecentCreatedBook(t *testing.T) {
  
    url := "http://localhost:8080/api/v1/book/isbn"
    isbn := "/2"
    method := "GET"
    client := &http.Client {}

    req, err := http.NewRequest(method, url+isbn, nil)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    res, err := client.Do(req)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    if res.StatusCode <= 200 && res.StatusCode >= 299 {
	    t.Errorf("Fail: HTTP Status is not in the 2xx range, got this: %v ", res.StatusCode)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    expected := `{"book_id":"2","title":"My Marmelade Tips","authors":"Jules Verne","average_rating":1.6,"isbn":"2","isbn_13":"1000000000002","language_code":"eng","num_pages":999,"ratings":25,"reviews":15}`
    if string(body) != expected {
	    t.Errorf("Fail: Does not match expected book")
    }
}


func TestDeleteCreatedBook(t *testing.T) {


    url := "http://localhost:8080/api/v1/book/isbn"
    isbn := "/2"
    method := "DELETE"
    client := &http.Client {}

    req, err := http.NewRequest(method, url+isbn, nil)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    res, err := client.Do(req)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    if res.StatusCode <= 200 && res.StatusCode >= 299 {
	    t.Errorf("Fail: HTTP Response not in the 2xx range: Got %v", res.StatusCode)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    expected := `{"success": "deleted"}`
    if string(body) != expected {
	    t.Errorf("FAIL: Expected delete message but instead got: %v", string(body))
    }
}


func TestSearchDeletedBookByISBN(t *testing.T) {

    url := "http://localhost:8080/api/v1/book/isbn/"
    isbn := "2"
    method := "GET"
    client := &http.Client {}

    req, err := http.NewRequest(method, url+isbn, nil)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    res, err := client.Do(req)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    if res.StatusCode <= 400 && res.StatusCode >= 499 {
	    t.Errorf("FAIL: HTTP Status is outside of 4xx range, got: %v", res.StatusCode)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
	    t.Fatalf(err.Error())
    }
    if string(body) != `{"error": "not found"}` {
	    t.Errorf("Fail: Expected book to be missing, got: %v", string(body))
    }
}


func TestCreateBookWithInvalidParameters(t *testing.T) {

  url := "http://localhost:8080/api/v1/book"
  method := "POST"
  // This payload has value hello for the BookID integer field, so should return error message
  payload := strings.NewReader("BookID=hello&Title=Aliens%20from%20Saturn&Authors=Donald%20Duck&ISBN=1&NumPages=12&LanguageCode=eng")
  client := &http.Client {}

  req, err := http.NewRequest(method, url, payload)
  if err != nil {
	  t.Fatalf(err.Error())
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  res, err := client.Do(req)
  if err != nil {
	  t.Fatalf(err.Error())
  }
  if res.StatusCode <= 400 && res.StatusCode >= 499 {
	  t.Errorf("FAIL: HTTP Status is outside of 4xx range, got this %v", res.StatusCode)
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
	  t.Fatalf(err.Error())
  }
  if string(body) == `{"success": "created"}` {
	  
	  // Request a delete of the irregular book
	  url := "http://localhost:8080/api/v1/book/isbn/"
	  isbn := "1"
	  method := "DELETE"
	  client := &http.Client {}
	  req, err := http.NewRequest(method, url+isbn, nil)
	  if err != nil {
		  t.Fatalf(err.Error())
	  }
	  res, err := client.Do(req)
	  if err != nil {
		  t.Fatalf(err.Error())
	  }
	  res.Body.Close()
	  t.Errorf("Book was created regardless of invalid parameter input, the book has been automatically deleted by its ISBN")
  }
}

