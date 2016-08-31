package truiteur

import (
	"fmt"
	"net/url"
	//"net/http"
	"golang.org/x/crypto/bcrypt"
)

var fishermen [2]Fisherman

type Fisherman struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Truites  []Truite `json:"truites"`
}

//For test purpose
func addFisherman() {
	fishermen[0] = Fisherman{"Othman", "Ben", nil}
	fishermen[1] = Fisherman{"Louis", "Mer", nil}
}

func AddFisherman(data url.Values) {
	fmt.Printf("New user!")
}

func encryptPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

/*func login( toutcequilfaut httpmachin ) {
	// accéder au username : req.PostFormValue("username")

	// IF LOGIN SUCCESSFUL
	http.SetCookie(res, &cookie)
	http.Redirect(res, req, "/", 301)
}*/
