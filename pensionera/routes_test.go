package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "testing"
)

func TestCreateBook(t *testing.T) {

  url := "http://localhost:8080/api/v1/book"
  method := "POST"

  payload := strings.NewReader("BookID=3&Title=Vulcan%20Philosophy&Authors=Huber%20Ugbert&ISBN=3&ISBN13=1000000000003&Ratings=25&Reviews=15&NumPages=999&LanguageCode=eng&AverageRating=2.5")

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}

func TestSearchByISBN(t *testing.T) {
  url := "http://localhost:8080/api/v1/book/isbn/3"
  method := "GET"
  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)
  if err != nil {
    fmt.Println(err)
    return
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
