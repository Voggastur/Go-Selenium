package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)


func TestCreateBook(t *testing.T) {

  url := "http://localhost:8080/api/v1/book/ISBN/1"
  method := "POST"
  payload := strings.NewReader("BookID=2&Title=Johans%20Polar%20Expedition&Authors=Jules%20Verne&ISBN=2&ISBN13=1000000000002&Ratings=25&Reviews=15&NumPages=999&LanguageCode=eng&AverageRating=1.6")

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
	// do something, if the result is not as expected
	t.Errorf("test result was not successful")
}