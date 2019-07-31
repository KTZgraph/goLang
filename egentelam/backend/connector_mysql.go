package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func connectMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/golang_user")

	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

func addUser(_login string, _password string, _userDescription string) error {
	//do rejestracji uzytkownika
	// Insert do bazy Mysql Nowego użytkownika

	password := []byte(_password) //zamiana stringa na bajty dla funckji hashujacej

	db, err := connectMysql()
	if err != nil {
		panic(err.Error())
		return err
	}

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	encryptedPassword := string(hashedPassword)

	queryInsert := fmt.Sprintf(`INSERT INTO users (
			login,
			password,
			userDescription,
			points
		)VALUES (
			"%s",
			"%s",
			"%s",
			"%d"
		)`, _login, encryptedPassword, _userDescription, 0) //przy rejestracji kzdy ma 0 punktow

	fmt.Printf(queryInsert)
	insert, err := db.Query(queryInsert)
	if err != nil {
		panic(err.Error())
		return err
	}
	fmt.Println(insert)
	return nil

}

func getAllUsers() []User {
	//pobiera wszystkich uzytkownikow - do wysiwetlenie na front

	allUsers := []User{}
	db, err := connectMysql()

	result, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var user User

		err = result.Scan(&user.UserId, &user.Login, &user.Password, &user.UserDescription, &user.Points, &user.IsLogged)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user)
		allUsers = append(allUsers, user)
	}

	// fmt.Println("All users: ", allUsers)
	// fmt.Println(reflect.TypeOf(allUsers))

	return allUsers
}

func checkLoginExists(_login string) (bool, error) {
	//sprawdza czy dany login jest juz w bazie, przy rejestracji przydatne
	db, err := connectMysql()

	querySelect := fmt.Sprintf(`SELECT login FROM users WHERE login='%s' ;`, _login)

	result, err := db.Query(querySelect)
	if err != nil {
		panic(err.Error())
		return false, err
	}

	for result.Next() {
		var userLogin string

		err = result.Scan(&userLogin)
		if err != nil {
			panic(err.Error())
			return false, err

		}

		if userLogin != "" {
			return true, nil
		}
	}

	return false, nil
}

func loginUser(_login string, _password string) (bool, error) {
	//do logowanie w bazie
	db, err := connectMysql()

	querySelect := fmt.Sprintf(`SELECT password FROM users WHERE login='%s' ;`, _login)

	result, err := db.Query(querySelect)
	if err != nil {
		panic(err.Error())
		return false, err
	}

	for result.Next() {
		var hashedPassword string

		err = result.Scan(&hashedPassword)
		if err != nil {
			panic(err.Error())
			return false, err
		}

		// Comparing the password with the hash
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(_password))
		if err == nil { // nil means it is a match
			return true, nil
		}
	}

	return false, nil
}

func updateUserPoints(_login string) error {
	db, err := connectMysql()
	var userPoints int

	querySelectPoints := fmt.Sprintf(`SELECT points FROM users WHERE login='%s' ;`, _login)

	result, err := db.Query(querySelectPoints)
	if err != nil {
		panic(err.Error())
		return err
	}

	for result.Next() { //pobieranie aktualnej liczby punktow uzytkownika

		err = result.Scan(&userPoints)
		if err != nil {
			panic(err.Error())
			return err
		}
	}

	queryUpdatePoints := fmt.Sprintf(`UPDATE users SET points = %d WHERE login ='%s' ;`, userPoints+1, _login)
	result, err = db.Query(queryUpdatePoints)
	if err != nil {
		panic(err.Error())
		return err
	}

	return nil //jak bez bledow wszystko
}

func getAllLoggedUsers() ([]string, error) {
	//pobieranie listy obiektow zalogowanych uzytkownikow
	db, err := connectMysql()
	loggedUsers := make([]string, 0)

	result, err := db.Query("SELECT * FROM users WHERE isLogged = true;")
	if err != nil {
		panic(err.Error())
		fmt.Println("[getAllLoggedUsers] [Error] %s", err)

		return nil, err // zwraca nula jak nie uda sie pobrac listy loginow
	}

	for result.Next() {
		var user User

		err = result.Scan(&user.UserId, &user.Login, &user.Password, &user.UserDescription, &user.Points, &user.IsLogged)
		if err != nil {
			panic(err.Error())
			return nil, err // zwraca nula jak nie uda sie pobrac listy loginow
			fmt.Println("[getAllLoggedUsers] [Error] %s", err)
		}

		fmt.Println(user)
		loggedUsers = append(loggedUsers, user.Login)
	}

	fmt.Println("Wszyscy zalogowani: ", loggedUsers)
	return loggedUsers, nil // zwraca nula jak nie uda sie pobrac listy loginow

}

func updateLoggedUser(_login string) error {
	//aktualizuje stan uzytkownika na zalogowanego
	db, err := connectMysql()
	if err != nil {
		panic(err.Error())
		fmt.Println("[updateLoggedUser] [Error] %s", err)
		return err
	}

	updateLogin := fmt.Sprintf(`UPDATE users SET isLogged=true WHERE login='%s' ;`, _login)
	_, err = db.Query(updateLogin)
	if err != nil {
		panic(err.Error())
		fmt.Println("[getAllLoggedUsers] [Error] %s", err)
		return err
	}

	return nil
}

func updateLogoutUser(_login string) error {
	//aktualizuje stan na false jak sie uzytkownik wyloguje
	db, err := connectMysql()
	if err != nil {
		panic(err.Error())
		fmt.Println("[updateLoggedUser] [Error] Nie udało polaczy sie z baza")
		return err
	}

	updateLogout := fmt.Sprintf(`UPDATE users SET isLogged=false WHERE login='%s' ;`, _login)
	result, err := db.Query(updateLogout)
	fmt.Println("RESULT: ", result)
	if err != nil {
		panic(err.Error())
		fmt.Println("[getAllLoggedUsers] [Error] Nie udalo sie zaktulizowac wylogowanego uzytkownika")
		return err
	}

	return nil

}
