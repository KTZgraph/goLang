//websocket zamiast rest api
// problem tylko z wysiwietlaniem danych z przegladarki
//Firefox:
//Zablokowano żądanie do zasobu innego pochodzenia: zasady „Same Origin Policy”
//nie pozwalają wczytywać zdalnych zasobów z „http://localhost:5000/socket.io/?EIO=3&transport=polling&t=MSKrtCr”
// (brakujący nagłówek CORS „Access-Control-Allow-Origin”).[Więcej informacji]

//Chrome: to samo co ^
//
//
//

package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io" //alias do biblioteki importu
)

func main() {
	fmt.Println("Hello World")

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("New Connection")
		//handle chat message to and front
		so.Join("chat")
		so.On("chat message", func(msg string) {
			log.Println("Message Received From Client: " + msg)
			so.BroadcastTo("chat", "chat message", msg)
		})

	})

	//file system handler
	fs := http.FileServer(http.Dir("static")) //do folderu
	http.Handle("/", fs)

	http.Handle("/socket.io/", server) //native html socket
	log.Fatal(http.ListenAndServe(":5000", nil))
}
