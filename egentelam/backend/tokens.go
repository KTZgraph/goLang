package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("DoTokenowDlaUszytkownikow")

func GenerateJWT(_login string) (string, error) {
	//generowanie tokena jwt
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = _login
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("[generateJWT][Error] Nie mozna wygenerować")
		return "", err
	}

	return tokenString, nil
}

func isTokenValid(tokenString string) (bool, error) {
	//sprawdza czy podany token jest pawidlowy
	if tokenString != "" { //JWT Parse przyjmuje stroinga, bo header norlanie jest lista, a ja mu tu wyluskane stringa dam
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { //sprawdzanie tokena
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error ")
			}
			//jesli token jest ok
			return mySigningKey, nil //zwraca token, do zmiennej 'token'
		}) //koniec dunkcji zwracajacej token

		if err != nil {
			fmt.Println(err.Error())
			return false, err
		}

		if token.Valid {
			//jesli token prawidlowy, to zwraca true
			return true, nil
		} else {
			fmt.Println("[isTokenValid] Nieprawidłowy token")
			return false, nil

		}

	}
	return false, nil

}
