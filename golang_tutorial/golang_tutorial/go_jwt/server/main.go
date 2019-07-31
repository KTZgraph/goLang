package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go" //jwt jako alias do repo =~ import numpy as np ; jwt to takie np
)

var mySigningKey = []byte("mysupersecretphrase")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Information from Server")
	fmt.Println("Endpoint Hit: homePage")

}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	//modyfikacja funkcji homePage =~ jak dektorator w pythonie
	//sprawdza czy jest prawidlowy JWTToken
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { //[!] handleR nie samo handle

		if r.Header["Token"] != nil {
			//walidacja
			//parsowanie naglowka z tokenem
			// token *jwt.Token token to zmienna - [!] bez przecinka w func(token *jwt.Token) (interface{}, error)
			// inaczej bład:
			//# command-line-arguments
			// .\main.go:26:55: undefined: token
			// .\main.go:27:17: undefined: token
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil //nil zamias errora
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r) //jak wszystko ok, to zwraca oryginalnego endpointa, przekazanego jako argument
			}
		} else {
			//brak tokena
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	handleRequests()
}
