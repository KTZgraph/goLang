package main

// go get github.com/dgrijalva/jwt-go
//cd client
// go mod init github.com/pawlaczyk/go-jwt/client

// [!] kropka przed nawiasami w token.Claims.(jwt.MapClaims)
//inaczej blad
//# command-line-arguments
// .\main.go:14:24: type jwt.MapClaims is not an expression
// .\main.go:14:24: cannot call non-function token.Claims (type jwt.Claims)

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go" //jwt jako alias do repo
)

var mySigningKey = []byte("mysupersecretphrase") //lepiej odczytywac to ze zmiennej srodowiskowej
// set My_JWT_TOKEN=mysupersecretphrase
//var mySigningKey = os.Get("My_JWT_TOKEN")  dobra praktyka dla jakiegogolwkiek hasła, chroni przed upubliczniniem gdzies hasla na repo

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{} //tworzy nowy Request
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
	req.Header.Set("Token", validToken) //dopisanie do hedera
	res, err := client.Do(req)
	if err != nil {
		//jak cos nie tak
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Fprintf(w, err.Error())

	}

	fmt.Fprintf(w, string(body))
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims) //[!] Kropka przed nawiasami !!!!!!!!!

	claims["authorized"] = true
	claims["client"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println("Somethig get wrong %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	fmt.Println("My Simple Client")
	handleRequests()
}
