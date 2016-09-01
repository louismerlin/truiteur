package main

import (
	//"encoding/json"
	"log"
	"net/http"
	"app/truiteur"
)

func main() {
	http.HandleFunc("/", truiteur.Validate(truiteur.HomePage))
	http.HandleFunc("/login", truiteur.Login)
	http.HandleFunc("/signup", truiteur.Signup)
	http.HandleFunc("/api/truite", truiteur.Validate(truiteur.ApiTruite))
	http.HandleFunc("/api/fisherman", truiteur.Validate(truiteur.ApiFisherman))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
