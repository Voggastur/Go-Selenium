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

  payload := strings.NewReader("BookID=2&Title=Johans%20Polar%20Expedition&Authors=Albert%20Einstein&ISBN=2&ISBN13=1000000000002&Ratings=25&Reviews=15&NumPages=999&LanguageCode=eng&AverageRating=1.6")

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

  url := "http://localhost:8080/api/v1/book/isbn/2"
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


func TestDeleteBook(t *testing.T) {

  url := "http://localhost:8080/api/v1/book/isbn/2"
  method := "DELETE"
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
