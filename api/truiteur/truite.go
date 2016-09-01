package truiteur

import (
	"fmt"
	"net/url"
)

type Truite struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

func ApiTruite(res http.ResponseWriter, req *http.Request) {
    if req.Method == "POST" {
      truiteur.AddTruite(r.PostForm)
    } else if req.Method == "GET" {
      if err := json.NewEncoder(w).Encode(truiteur.GetTruite(r.PostForm)); err != nil {
        panic(err)
      }
    }
  })

}


func AddTruite(data url.Values) {
	fmt.Printf("Post!")
}

func GetTruite(data url.Values) Truite {
	return Truite{Id: 1, Body: "Hello, world!"}
}
