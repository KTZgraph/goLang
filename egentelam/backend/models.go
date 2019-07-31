package main

import "github.com/zemirco/couchdb"

//struktura użytkownika
/*
CREATE TABLE users (
	userId INT NOT NULL AUTO_INCREMENT,
	login varchar(50)  NOT NULL UNIQUE,
	password varchar(30)  NOT NULL,
	userDescription varchar(300),
	points int,

	PRIMARY KEY (userId)
);
*/
type User struct {
	UserId          int    `json:"userId"` // do pobierania z GET
	Login           string `json:"login"`
	Password        string `json:"password"`
	UserDescription string `json:"userDescription"`
	Points          int    `json:points`
	Token           string
	IsLogged        bool
}

// Structura karty
/*
 */
type Card struct {
	couchdb.Document
	Id         string `json:"id" binding:"required"`
	IsQuestion bool   `json:"isQuestion"`               //0 karta pytanie, 1 karta odpowiedź
	Blank      int    `json:"blank" binding:"required"` //ile kart odpowiedzi na pytanie
	Text       string `json:"text" binding:"required"`  // podłoga to luka
	Timestamp  float64
}
