package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)


func TestSearchByAuthor() {

  url := "localhost:8080/api/v1/book/ISBN/3"
  method := "curl"

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
