package main

import (
    "fmt"
    "log"
    "net/url"
    "net/http"
    "encoding/json"
)

func main() {
    http.HandleFunc("/api/truite", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
			addTruite(r.PostForm)
		} else if r.Method == "GET" {
			if err := json.NewEncoder(w).Encode(getTruite()); err != nil {
				panic(err)
			}
		}
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}

type Truite struct {
	Id 		int		`json:"id"`
	Body 	string	`json:"body"`
}

func addTruite(data url.Values){
	fmt.Printf("Post!")
}

func getTruite() Truite {
	return Truite{Id: 1, Body: "Hello, world!"}
}

type Truites []Truite


