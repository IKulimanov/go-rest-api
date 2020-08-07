package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var (
	jwtKey = []byte("my_secret_key")

	users = map[string]string{
		"user1": "password1",
	}

)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`

}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(writer http.ResponseWriter, request *http.Request){
	var creds Credentials

	err := json.NewDecoder(request.Body).Decode(&creds)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{

			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Printf("ERR")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	})
}
