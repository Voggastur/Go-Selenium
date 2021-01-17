package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)


func TestCreateBook() {

  url := "http://localhost:8080/api/v1/book"
  method := "POST"

  payload := strings.NewReader("BookID=2&Title=Johans%20Polar%20Expedition&Authors=Albert$

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


func TestSearchByISBN() {

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


func TestDeleteBook() {

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


func main() {
  TestCreateBook()
  TestSearchByISBN()
  TestDeleteBook()
}
