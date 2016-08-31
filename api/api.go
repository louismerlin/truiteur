package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"app/truiteur"
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
