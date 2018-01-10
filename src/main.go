package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var client = make(map[*websocket.Conn]bool)
var braodcast = make(chan Message)

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

	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("http server started on port 8080")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}
