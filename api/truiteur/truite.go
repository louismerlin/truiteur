package truiteur

import (
  "fmt"
  "net/url"
)

type Truite struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

func AddTruite(data url.Values) {
	fmt.Printf("Post!")
}

func GetTruite(data url.Values) Truite {
	return Truite{Id: 1, Body: "Hello, world!"}
}
