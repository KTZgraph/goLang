package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

type User2 struct {
	Name  string `json:"name8"`
	Name2 string `json:"name8"`
}

func main() {
	fmt.Println("Go Mysql Tutorial")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Succesfully Connected to MySQL database")
	defer db.Close() //keep the connection open for the remainder of this function

	// insert, err := db.Query("INSERT INTO users VALUES('ELLIOT3')")
	insert, err := db.Query("INSERT INTO users2 VALUES('a1', 'b1')")

	if err != nil {
		panic(err.Error)
	}

	defer insert.Close()
	fmt.Println("Succes Insert")

	results, err := db.Query("SELECT * FROM users2")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user2 User2

		err = results.Scan(&user2.Name, &user2.Name2)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user2.Name, user2.Name2)
	}

	fmt.Println(*results)

}
