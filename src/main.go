package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

// upgrader config
var upgrader = websocket.Upgrader{}

// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {

	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// start listening to incoming chat messages
	go handleMessages()

	// start the server on localhost port 8080
	log.Println("http server started on port 8080")
	err := http.ListenAndServe(":8000", nil)
	// log errors
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	// upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// close connection when function returns
	defer ws.Close()

	// register new client
	clients[ws] = true

	for {
		var msg Message

		// read in a new message as a JSON and map it to a Message Object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)

		}

		// send the newly grabbed message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// grabs the next message from broadcast channel
		msg := <-broadcast

		// send the message to every client currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error %v", err)
				client.Close()
				delete(clients, client)
			}
		}

	}
}
