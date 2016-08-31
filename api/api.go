package main

import (
	"encoding/json"
	"log"
	"net/http"
	"app/truiteur"
)

func main() {
	http.HandleFunc("/api/truite", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			truiteur.AddTruite(r.PostForm)
		} else if r.Method == "GET" {
			if err := json.NewEncoder(w).Encode(truiteur.GetTruite()); err != nil {
				panic(err)
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
