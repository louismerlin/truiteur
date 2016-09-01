package main

import (
	"log"
	"net/http"
	"app/truiteur"
)

func main() {
	truiteur.AddFishermanTest()
	http.HandleFunc("/api/login", truiteur.Login)
	http.HandleFunc("/api/signup", truiteur.Signup)
	http.HandleFunc("/api/truite", truiteur.Validate(truiteur.ApiTruite))
	http.HandleFunc("/api/fisherman", truiteur.Validate(truiteur.ApiFisherman))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
