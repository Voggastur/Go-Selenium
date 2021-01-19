package main

import (
    "testing"
    "net/http"
    "io/ioutil"
    "fmt"
    "strings"
    "log"
)


func TestSearchBookByISBN(t *testing.T) {


    url := "http://localhost:8080/api/v1/book/isbn"
    isbn := "/074754624X"
    method := "GET"
    client := &http.Client {}

    req, err := http.NewRequest(method, url+isbn, nil)
    if err != nil {
        fmt.Println("HTTP call failed", err)
        return
    }
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("HTTP call failed:", err)
	return
    }
    if res.StatusCode >= 200 && res.StatusCode <= 299 {
        fmt.Println("Success, HTTP Status is in the 2xx range: Got ", res.StatusCode)
    } else {
        fmt.Println("HTTP Response not in the 2xx range: Got ", res.StatusCode)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
	fmt.Println(err)
	return
    }
    expected := `{"book_id":"43509","title":"Harry Potter and the Goblet of Fire (Harry Potter  #4)","authors":"J.K. Rowling","average_rating":4.55,"isbn":"074754624X","isbn_13":"9780747546245","language_code":"eng","num_pages":636,"ratings":18097,"reviews":860}`
    if string(body) != expected {
        fmt.Println("Fail: The found book does not match expected book")
    } else {
        fmt.Println("Success: The found book matches the expected Harry Potter book")
    }
    fmt.Println(string(body))
}


func TestSearchByAuthor(t *testing.T) {

    req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/books/authors/galileo%20galilei", nil)
    if err != nil {
        log.Fatal("HTTP Request Failed: Got ", err)
    }

    client := &http.Client {}
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("HTTP call failed: Got ", err)
        return
    }
    if res.StatusCode >= 200 && res.StatusCode <= 299 {
        fmt.Println("Success: HTTP Status is in the 2xx range: Got ", res.StatusCode)
    } else {
        fmt.Println("Fail: HTTP Response not what expected, Got: ", res.StatusCode)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    expected := `[{"book_id":"32610","title":"On the Shoulders of Giants: The Great Works of Physics and Astronomy","authors":"Stephen Hawking-Isaac Newton-Nicolaus Copernicus-Albert Einstein-Johannes Kepler-Galileo Galilei","average_rating":4.2,"isbn":"0762427329","isbn_13":"9780762427321","language_code":"eng","num_pages":256,"ratings":4014,"reviews":34}]`
    if string(body) != expected {
        fmt.Println("Fail: Does not match expected book")
    } else {
        fmt.Println("Success: Found expected book by Galileo Galilei")
    }
    fmt.Println(string(body))
}


func TestCreateBook(t *testing.T) {

  url := "http://localhost:8080/api/v1/book"
  method := "POST"

  payload := strings.NewReader("BookID=2&Title=Johans%20Marmelade%20Tips&Authors=Jules%20Verne&ISBN=2&ISBN13=1000000000002&Ratings=25&Reviews=15&NumPages=999&LanguageCode=eng&AverageRating=1.6")

  client := &http.Client {}
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println("Fail: HTTP Call failed: ", err)
    return
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }

  if res.StatusCode != 201 {
    fmt.Println("HTTP Status not what expected: Got ", res.StatusCode)
  } else {
    fmt.Println("Got expected 201 HTTP Status: Got ", res.StatusCode)
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  if string(body) != `{"success": "created"}` {
    fmt.Println("Success created not found, instead got: ", string(body))
  } else {
    fmt.Println("Test passed - Got expected success message ", string(body))
  }
}


func TestReadRecentCreatedBook(t *testing.T) {
  
    url := "http://localhost:8080/api/v1/book/isbn"
    isbn := "/2"
    method := "GET"
    client := &http.Client {}

    req, err := http.NewRequest(method, url+isbn, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("HTTP call failed: ", err)
        return
    }
    if res.StatusCode >= 200 && res.StatusCode <= 299 {
        fmt.Println("Success: HTTP Status is in the 2xx range ", res.StatusCode)
    } else {
        fmt.Println("Fail: HTTP Response not what expected: ", res.StatusCode)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(body))
}


func TestDeleteCreatedBook(t *testing.T) {


    url := "http://localhost:8080/api/v1/book/isbn"
    isbn := "/2"
    method := "DELETE"
    client := &http.Client {}

    req, err := http.NewRequest(method, url+isbn, nil)
    if err != nil {
        fmt.Println("HTTP call failed", err)
        return
    }
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("HTTP call failed:", err)
        return
    }
    if res.StatusCode >= 200 && res.StatusCode <= 299 {
        fmt.Println("Success: HTTP Status is in the 2xx range: Got ", res.StatusCode)
    } else {
        fmt.Println("Fail: HTTP Response not in the 2xx range: Got ", res.StatusCode)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    expected := `{"success": "deleted"}`
    if string(body) != expected {
        fmt.Println("Fail: Expected delete but instead got:", string(body))
    } else {
        fmt.Println("Success: Got correct delete message: ", string(body))
    }
}


func TestSearchDeletedBookByISBN(t *testing.T) {

    url := "http://localhost:8080/api/v1/book/isbn/"
    isbn := "2"
    method := "GET"
    client := &http.Client {}

    req, err := http.NewRequest(method, url+isbn, nil)
    if err != nil {
        fmt.Println("HTTP call failed", err)
        return
    }
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("HTTP call failed:", err)
        return
    }
    if res.StatusCode >= 400 && res.StatusCode <= 499 {
        fmt.Println("PASS: HTTP Status is in the 4xx range ", res.StatusCode)
    } else {
        fmt.Println("FAIL: HTTP Status outside the 4xx range: ", res.StatusCode)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println("Error reading response body")
        return
    }

    if string(body) != `{"error": "not found"}` {
	fmt.Println("Fail: Expected book to be deleted, instead got:", string(body))
    } else {
        fmt.Println("Success: Book is correctly missing after previous deletion, got this:", string(body))
    }
}


func TestCreateBookWithInvalidParameters(t *testing.T) {

  url := "http://localhost:8080/api/v1/book"
  method := "POST"

  // This payload has value hello for the BookID field, so should return error message
  payload := strings.NewReader("BookID=hello&Title=Aliens%20from%20Saturn&Authors=Donald%20Duck&ISBN=1&NumPages=12&LanguageCode=eng")

  client := &http.Client {}
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println("Fail: HTTP Call failed: ", err)
    return
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println("HTTP Response failed, got:", err)
    return
  }

  if res.StatusCode >= 400 && res.StatusCode <= 499 {
        fmt.Println("Success: HTTP Status in the 4xx range ", res.StatusCode)
    } else {
        fmt.Println("Fail: HTTP Status outside the 4xx range: ", res.StatusCode)
    }

  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }

  if string(body) != `{"success": "created"}` {
    fmt.Println("Success: Book can not be created, got this:", string(body))
  } else {
    fmt.Println("Fail: success created regardless of invalid field(s), message received:", string(body))

    // Request a delete of the irregular book in the else block

    url := "http://localhost:8080/api/v1/book/isbn/"
    isbn := "1"
    method := "DELETE"
    client := &http.Client {}

    req, err := http.NewRequest(method, url+isbn, nil)
    if err != nil {
        fmt.Println("HTTP call failed", err)
        return
    }
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("HTTP call failed:", err)
        return
    }
    res.Body.Close()
    fmt.Println("The irregularly created book has been deleted by its ISBN")
  }
}

