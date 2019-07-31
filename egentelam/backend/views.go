package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// go get "golang.org/x/crypto/bcrypt"
)

type TokenAPI struct {
	Token string
}

func getUsersView(c *gin.Context) {
	// dodanie nowej karty do bzy
	fmt.Println("-----------------------------WYPISANIE TOKENU ------------------------------")

	var token TokenAPI
	c.Bind(&token)
	fmt.Println("MOJ TOKEN %s", token.Token)

	// c.Header("Content-Type", "application/json")
	fmt.Println("Pobieranei listy użytkowników")

	var userList []map[string]interface{}

	allUsers := getAllUsers()
	for _, arg := range allUsers {
		tmp := make(map[string]interface{})
		tmp["login"] = arg.Login
		tmp["userDescription"] = arg.UserDescription
		tmp["points"] = arg.Points
		userList = append(userList, tmp)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"allUsers": userList, // cast it to string before showing
	})
}

func addNewUserView(c *gin.Context) {
	//Dane z frontu - json z danymi uzytkownika, bedzie jeden
	//json walidacja na froncie - tu nie musze walidowac i elo
	// dobra, jednak waliduje
	c.Header("Content-Type", "application/json")
	var newUser User
	c.Bind(&newUser)

	fmt.Println(newUser.Login)
	_login := newUser.Login
	_password := newUser.Password
	_userDescription := newUser.UserDescription

	if _login == "" {
		c.JSON(http.StatusOK, "[addNewUserView][Error] Nie podano loginu")
		return
	}

	if _password == "" {
		c.JSON(http.StatusOK, "[addNewUserView][Error] Nie podano hasła")
		return
	}

	isExists, err := checkLoginExists(_login)

	if err != nil {
		c.JSON(http.StatusOK, "[addNewUserView][Error] Nie mozna dodac do bazy")
		return
	}

	if isExists {
		c.JSON(http.StatusOK, "Login zajęty")
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, "[addNewUserView][Error] Nie mozna zaszyfrowac hasla")
		return
	}

	err = addUser(_login, _password, _userDescription)

	if err != nil {
		c.JSON(http.StatusOK, "[addNewUserView][Error] Nie mozna dodac do bazy")
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, "[addNewUserView] Dodano uzytkownika do bazy")
}

func loginUserView(c *gin.Context) {
	//logowanie - czy jest w bazie
	fmt.Println("HEARES--------------------------------")
	fmt.Println(c.Header)
	fmt.Println(c.Request)
	c.Header("Content-Type", "application/json")

	var checkUser User
	c.Bind(&checkUser)
	_login := checkUser.Login
	_password := checkUser.Password
	_token := checkUser.Token

	fmt.Println("_login: ", _login)
	fmt.Println("_pasowrd: ", _password)
	fmt.Println("_token: ", _token)

	if _login == "" {
		// c.JSON(http.StatusOK, "[loginUserView][Error] Nie podano loginu")
		c.JSON(http.StatusOK, false)

	}

	if _password == "" {
		// c.JSON(http.StatusOK, "[loginUserView][Error] Nie podano hasła")
		c.JSON(http.StatusOK, false)

	}

	validLoginData, err := loginUser(_login, _password)
	if err != nil {
		// c.JSON(http.StatusOK, "[loginUserView][Error] Nie mozna zalogowac")
		c.JSON(http.StatusOK, false)

	}

	if validLoginData { //jezeli zwortic true - poprawne dane to generuje tokena
		tokenString, err := GenerateJWT(_login) //generowanie tokenu dla dane uzytkownika
		if err != nil {                         //jezeli jest bład
			// c.JSON(http.StatusOK, "[loginUserView][Error] Nie mozna wygenerowac tokenu")
			c.JSON(http.StatusOK, false)

		}
		c.JSON(http.StatusOK, tokenString) //zwraca token jak zalogowano
	} else {
		c.JSON(http.StatusOK, false) //zwraca token jak zalogowano
	}
}

func updateUserPointsView(c *gin.Context) {
	//dodawanie punktu +1 dla uzytkownika
	c.Header("Content-Type", "application/json")

	var checkUser User
	c.Bind(&checkUser)
	_login := checkUser.Login

	if _login == "" {
		c.JSON(http.StatusOK, "[updateUserPointsView][Error] Nie podano loginu")
		return
	}

	isExists, err := checkLoginExists(_login)
	if err != nil {
		c.JSON(http.StatusOK, "[updateUserPointsView][Error] Nie mozna dodac punktow")
		return
	}
	if !isExists {
		c.JSON(http.StatusOK, "[updateUserPointsView][Error] Brak uzytkownika w bazie")
		return
	}

	//TODO dopisac update punktow na bazie
	err = updateUserPoints(_login)
	if err != nil {
		c.JSON(http.StatusOK, "[updateUserPointsView][Error] Nie mozna dodac punktow")
		return
	}

	c.JSON(http.StatusOK, "[updateUserPointsView] Dodano punkt")
}

func addNewCardView(c *gin.Context) {
	//dodawanie karty na backendzie, odczytywanie wszystkich na froncie bo to nie sa zadne dane wrazliwe
	//ale dodawanei tutaj bo i tak mozna to zrobic w tej samem domenie co couchdb
	//PATRZ:  documentation/couchdb/local.ini; documentation/couchdb/corsy_ustawienie.txt
	c.Header("Content-Type", "application/json")
	fmt.Println("--------------------        ADD NEW CARD    ----------------------------------------------")
	var newCard Card
	c.Bind(&newCard)
	_isQuestion := newCard.IsQuestion
	_blank := newCard.Blank
	_text := newCard.Text
	fmt.Println("_isQuestion: ", _isQuestion)
	fmt.Println("_blank: ", _blank)
	fmt.Println("_text: ", _text)

	if _text == "" {
		c.JSON(http.StatusOK, "[updateUserPointsView][Error] Nie podano 'text'")
		return
	}

	err := addNewCard(_isQuestion, _blank, _text)
	if err != nil {
		fmt.Println("Err: ")
		c.JSON(http.StatusOK, "[addNewCardView][Error] Nie można oddac nowej karty do [couchdb]")
		return
	}

	c.JSON(http.StatusOK, "[addNewCardView] Dodano nową kartę")

}

func getAllLoggedUsersView(c *gin.Context) {
	fmt.Println("Zwracanie wszystkich zalogowanych uzytkownikow")

	allLoggedUsers, err := getAllLoggedUsers()
	if err != nil {
		fmt.Println("[getAllLoggedUsersView] [Error] %s", err)
		c.JSON(http.StatusOK, gin.H{
			"code":           http.StatusOK,
			"allLoggedUsers": allLoggedUsers, // cast it to string before showing
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":           http.StatusOK,
			"allLoggedUsers": nil, // cast it to string before showing
		})
	}
}

func updateLoggedUserView(c *gin.Context) {
	var loggedUser User //login zalogowane uzytkownika
	c.Bind(&loggedUser)
	_login := loggedUser.Login
	err := updateLoggedUser(_login)
	if err != nil {
		c.JSON(http.StatusOK, "[updateLoggedUserView] Nie udalo sie zaktualizowac zalogowanego uzytkownika")
		return
	}

	c.JSON(http.StatusOK, true) //true jak sie udalo zaktualizowac
}

func updateLogoutUserView(c *gin.Context) {
	var logoutUser User //login zalogowane uzytkownika
	c.Bind(&logoutUser)
	_login := logoutUser.Login
	fmt.Println("Wylogowano uzytkownika: %s", _login)

	err := updateLogoutUser(_login)

	if err != nil {
		c.JSON(http.StatusOK, "[updateLoggedUserView] Nie udalo sie zaktualizowac wylogowanego uzytkownika")
		return
	}

	c.JSON(http.StatusOK, true) //true jak sie udalo zaktualizowac
}

func getGentelman(c *gin.Context) {
	//zwraca login aktualnego gentelamana
}

func getAllCardsView(c *gin.Context) {
	fmt.Println("Wszystkie karty z  couchdb")
	var cardList []map[string]interface{}

	allCards := getAllCards()
	for _, arg := range allCards {
		tmp := make(map[string]interface{})
		tmp["Text"] = arg.Text
		tmp["Timestamp"] = arg.Timestamp
		tmp["IsQuestion"] = arg.IsQuestion
		tmp["Id"] = arg.Id
		tmp["Blank"] = arg.Blank
		cardList = append(cardList, tmp)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"allCards": cardList, // cast it to string before showing
	})

}
