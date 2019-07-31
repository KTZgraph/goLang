package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// var mySigningKey = os.GET("MY_JWT_TOKEN") - hasła z zmiennych środowiskocyh systemu
var mySigningKey = []byte("mysupersecrephrase")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil) //Nowe zapytanie GET
	req.Header.Set("Token", validToken)                             //Tworzenie JWT Tokena
	response, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// fmt.Fprintf(w, validToken)
	fmt.Fprintf(w, string(body))

}

func handleRequest() {
	http.HandleFunc("/", homePage) //taki url z Django
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Eliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("[generateJWT][Error] Nie mozna wygenerować")
		return "", err
	}

	return tokenString, nil
}

func main() {
	fmt.Println("My simple client")

	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Error generator token string")
	}

	fmt.Println("tokenString: ", tokenString)

	handleRequest()

}
