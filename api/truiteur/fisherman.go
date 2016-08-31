package truiteur

import "golang.org/x/crypto/bcrypt"

type Fisherman struct {
	Username   	string  `json:"username"`
	Password 	string 	`json:"password"`
	Truites		Truite[] `json:"truites"`
}

func AddFisherman(data url.Values) {
	
}

func encryptPassword(password string) string {
	pswd := []byte(password)
	
	hashedPassword, err := bcrypt.GenerateFromPassword(pswd, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	
	return hasehdPassword
}
