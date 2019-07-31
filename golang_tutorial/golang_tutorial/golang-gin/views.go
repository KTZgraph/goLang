package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllCards(c *gin.Context) {
	fmt.Println("Pobieranie wszystkich kart z bazy - pytań i odpowiedzi")
	fmt.Println("Długość listy: ", len(cards))

	c.JSON(http.StatusOK, &cards)
}

func addNewCard() {
	fmt.Println("Dodanie do couchDB nowej karty pytania lub odowiedzi")
}

func addNewUser() {
	fmt.Println("Rejestracja nowego użytkownika")
}

func loginUser() {
	fmt.Println("Logowanie użytkownika")
}
