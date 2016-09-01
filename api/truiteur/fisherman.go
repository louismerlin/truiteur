package truiteur

import (
	"fmt"
	"net/url"
	"net/http"
	"encoding/json"
	
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
      AddTruite(req.PostForm)
    } else if req.Method == "GET" {
      if err := json.NewEncoder(res).Encode(GetTruite(req.PostForm)); err != nil {
        panic(err)
      }
    }
}

//For test purpose
func AddFishermanTest() {
	fishermen[0] = Fisherman{"Othman", encryptPassword("Ben"), nil}
	fishermen[1] = Fisherman{"Louis", encryptPassword("Mer"), nil}
}

func AddFisherman(data url.Values) {
	fmt.Printf("New user!")
}

func encryptPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  fmt.Println(string(hashedPassword))
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func Login(res http.ResponseWriter, req *http.Request) {
  fmt.Println(fishermen[0].Username)
  fmt.Println("LOGIN!!")
	login := false;
	for i := 0; i < 2; i++ {
		if fishermen[i].Username ==  req.PostFormValue("username") {
			login = true;
			if bcrypt.CompareHashAndPassword([]byte(fishermen[i].Password), []byte(req.PostFormValue("password"))) == nil {
				cookie := loginCookie(fishermen[i].Username)
				http.SetCookie(res, &cookie)
				http.Redirect(res, req, "/", 301)
			} else {
				fmt.Println("You've entered the wrong password, please try again!")
			}
		}
	}

	if login == false {
		fmt.Printf("There is no such username, please try again")
	}
}

func Signup(res http.ResponseWriter, req *http.Request) {
  fmt.Printf("THIS IS A PRIVATE CLUB GET OUT")
}
