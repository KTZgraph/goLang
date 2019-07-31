package main

import (
	"fmt"
)

func GenerateJWT() (string, error) {
	fmt.Println("Generowanie JWT tokena dla klienta")
}

func ValidateToken() (string, error) {
	fmt.Println("Walidacja Tokena dla tresci dostepnych tylko dla zalogowanych użytkowników")
}
