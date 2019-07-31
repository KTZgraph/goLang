// Użycie ORM dla bazy sqlite3
// w innych bazach na podobnej zasadzie
// operowanie raczej na metodach db niż modeli
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct { // ~= models.py
	gorm.Model //modeled somehow in the database
	Name       string
	Email      string
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{}) //migracja modelu
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	//widok
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var users []User //lista obiektow
	db.Find(&users)
	json.NewEncoder(w).Encode(users) //parsowanie lisy na jsona

}

func NewUser(w http.ResponseWriter, r *http.Request) {
	//widok
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to database")
	}
	defer db.Close()

	// pobieranie danych z url'a
	vars := mux.Vars(r) //wyłuskanie danych z obiekt request
	name := vars["name"]
	email := vars["email"]

	// tworzenie obiektu Name - atrybut struktury, name wartosc zmiennej z url'a
	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New user successfully created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//widok
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)  //wyłuskanie danych z  obiekt request
	name := vars["name"] //wyłuskanie danych ze zmiennej z url'a
	//[!] USUWA WSZYSTKO  db.Delete(&User{Name: name}) //USUWA WSZYSTKO [!]

	var user User
	//znajdz obiekt user, gdzie atrybut name jest równy zmiennej name
	db.Where("name = ?", name).Find(&user)
	//znajduje tylko jeden obiekt jesli dane sa identyczne w baze i tylko jeden usuwa
	db.Delete(&user)
	fmt.Fprintf(w, "DELETE User")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//widok
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name =?", name).Find(&user)

	user.Email = email

	db.Save(&user)

	fmt.Fprintf(w, "UPDATE User")
}
