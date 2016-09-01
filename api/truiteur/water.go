package truiteur

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	
	"github.com/gorilla/context"
	"github.com/dgrijalva/jwt-go"
)

type TruiteurClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func HomePage(res http.ResponseWriter, req *http.Request) {}

func loginCookie(username string) http.Cookie {
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	expireCookie := time.Now().Add(time.Hour * 24)

	claims := TruiteurClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "truiteur",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("secret"))

	return http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}

}

func Validate(protectedPage http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		cookie, err := req.Cookie("Auth")
		if err != nil {
			http.NotFound(res, req)
			return
		}

		splitCookie := strings.Split(cookie.String(), "Auth=")

		token, err := jwt.ParseWithClaims(splitCookie[1], &TruiteurClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if claims, ok := token.Claims.(*TruiteurClaims); ok && token.Valid {
			context.Set(req, "Claims", claims)
		} else {
			http.NotFound(res, req)
			return
		}

		protectedPage(res, req)
	})
}

/*func profile(res http.ResponseWriter, req *http.Request) {
	claims := context.Get(req, "Claims").(*MyCustomClaims)
	res.Write([]byte(claims.Username))
	context.Clear(req)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Home Page"))
}*/
