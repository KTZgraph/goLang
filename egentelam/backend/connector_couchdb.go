package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"

	// "github.com/mikebell-org/go-couchdb"
	"github.com/zemirco/couchdb"
)

// create your own document
type cardDocument struct {
	couchdb.Document
	IsQuestion bool   `json:"isquestion"`
	Blank      int    `json:"blank"`
	Text       string `json:"text`
	Timestamp  int64  `json:"timestamp` //na froncie jest
}

func connectCouchdb() couchdb.DatabaseService {
	url_couchdb := "http://root:password@127.0.0.1:5984/"

	u, err := url.Parse(url_couchdb)
	if err != nil {
		panic(err)
	}

	client, err := couchdb.NewClient(u)

	if err != nil {
		panic(err)
	}

	db := client.Use("golang_cards")
	return db
}

func addNewCard(_isQuestion bool, _blank int, _text string) error {
	if !_isQuestion { //jezeli nie jest karta pytanie to zero pol pustych
		_blank = 0
	}

	db := connectCouchdb()

	now := time.Now()
	unixNano := now.UnixNano()
	umillisec := unixNano / 1000000

	doc := &cardDocument{ //struktura karty
		IsQuestion: _isQuestion,
		Blank:      _blank,
		Text:       _text,
		Timestamp:  umillisec,
	}

	result, err := db.Post(doc)
	if err != nil {
		panic(err)
		// return err
	}
	fmt.Println(reflect.TypeOf(result))

	// if err := db.Get(doc, result.ID); err != nil { // get id and current revision.
	// 	panic(err)
	// 	return err //nie matakiego dodane pliku - cos sie nie powiodlo
	// }

	fmt.Printf(result.ID)
	return nil
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getAllCards() []Card {

	//przez zapytanie bardziej optymalne i tak karza robic
	//
	resp, err := http.Get("http://localhost:5984/golang_cards/_all_docs?include_docs=true")
	if err != nil {
		log.Fatal(err)
	}
	var generic map[string]interface{}
	//fmt.Println(resp.Body) //&{0xc000040740 {0 0} false <nil> 0x693ad0 0x693a50}

	err = json.NewDecoder(resp.Body).Decode(&generic)
	if err != nil {
		log.Fatal(err)
	}
	keys := reflect.ValueOf(generic).MapKeys()
	fmt.Println("KEYS: ", keys)

	allCards := []Card{}

	for k, v := range generic {
		if k == "rows" {
			switch val := v.(type) {
			case []interface{}:
				fmt.Println(k, "is an array")
				for _, v := range val {
					var tmpCard Card

					mResult := v.(map[string]interface{})
					mResult2 := mResult["doc"].(map[string]interface{})
					// fmt.Println(mResult2["Text"])
					// fmt.Println(reflect.TypeOf(mResult2["Text"]))
					tmpCard.Text = mResult2["Text"].(string)
					// fmt.Println(mResult2["Timestamp"])
					tmpCard.Timestamp = mResult2["Timestamp"].(float64)

					// fmt.Println(mResult2["_id"])
					tmpCard.Id = mResult2["_id"].(string)

					// fmt.Println(mResult2["blank"])
					tmpCard.Blank = int(mResult2["blank"].(float64))

					// fmt.Println(mResult2["isquestion"])
					tmpCard.IsQuestion = mResult2["isquestion"].(bool)
					allCards = append(allCards, tmpCard)

				}
			}
		}
	}
	return allCards

}
