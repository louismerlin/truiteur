package truiteur

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var fishermen := [2]Fisherman{new Fisherman{Othman, Ben, nil}, new Fisherman{Louis, Mer, nil}}

type Fisherman struct {
	Username   	string  `json:"username"`
	Password 	string 	`json:"password"`
	Truites		Truite[] `json:"truites"`
}

func AddFisherman(data url.Values) {
	fmt.Printf("New user!")
}

func encryptPassword(password string) string {
	pswd := []byte(password)
	
	hashedPassword, err := bcrypt.GenerateFromPassword(pswd, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	
	return hasehdPassword
}
