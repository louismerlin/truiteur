package truiteur

import (
	"fmt"
	"net/url"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

var fishermen [2]Fisherman

type Fisherman struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Truites  []Truite `json:"truites"`
}

func ApiFisherman(res http.ResponseWriter, req *http.Request) {
    if req.Method == "POST" {
      truiteur.AddTruite(r.PostForm)
    } else if req.Method == "GET" {
      if err := json.NewEncoder(w).Encode(truiteur.GetTruite(r.PostForm)); err != nil {
        panic(err)
      }
    }
  })

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

func login(res http.ResponseWriter, req *http.Request) {
	login := false;
	
	for i := 0; i < 2; i++ {
		if fishermen[i].Username ==  req.PostFormValue("username") {
			login = true;
			if fishermen[i].Password == encryptPassword(req.PostFormValue("password")) {
				cookie := loginCookie(fishermen[i].Username)
				http.SetCookie(res, &cookie)
				http.Redirect(res, req, "/", 301)
			} else {
				fmt.Printf("You've entered the wrong password, please try again!")
			}
		}
	}
	
	if login == false {
		fmt.Printf("There is no such username, please try again")
	}
}
