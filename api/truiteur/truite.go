package truiteur

import (
	"fmt"
	"net/url"
	"net/http"
	"encoding/json"
)

type Truite struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

func ApiTruite(res http.ResponseWriter, req *http.Request) {
    if req.Method == "POST" {
      AddTruite(req.PostForm)
    } else if req.Method == "GET" {
      if err := json.NewEncoder(res).Encode(GetTruite(req.PostForm)); err != nil {
        panic(err)
      }
	}
}

func AddTruite(data url.Values) {
	fmt.Printf("Post!")

func addTruite(data url.Values) {
	fmt.Printf("Post truite !")
}

func getTruite(data url.Values) Truite {
	return Truite{Id: 1, Body: "Hello, world!"}
}
