package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecrephrase")

func homePage(w http.ResponseWriter, r *http.Request) {
	//dzieki isAuthorized nie trzeba zmieniac oryginalnej funkcji
	fmt.Fprintf(w, "Super Secret Information")
}

func isAuthorized(enpoint func(http.ResponseWriter, *http.Request)) http.Handler { //bierze request

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			fmt.Println("_______________TYP TOKENA ________________________")
			fmt.Println(reflect.TypeOf(r.Header["Token"]))

			fmt.Println("_______________r.Header[Token][0]________________________")
			fmt.Println(reflect.TypeOf(r.Header["Token"][0]))

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) { //sprawdzanie tokena
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error ")
				}
				//jesli token jest ok
				return mySigningKey, nil //[]byte("mysupersecrephrase")
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				//jesli token prawidlowy
				enpoint(w, r) //zwraca oryginalny endpoint
			}

		} else { //jesli brak token
			fmt.Fprintf(w, "Not authorized")
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("My Simple Server")
	handleRequests()
}
