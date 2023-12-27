package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var addr = "ws://localhost:8080/ws"

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Print("Enter your username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()

	go func() {
		for {
			var msg Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Println("Error reading JSON:", err)
				return
			}
			fmt.Printf("[%s] %s\n", msg.Username, msg.Content)
		}
	}()

	for {
		message := readInput()

		msg := Message{
			Username: username,
			Content:  message,
		}

		err := conn.WriteJSON(msg)
		if err != nil {
			log.Println("Error writing JSON:", err)
			return
		}

		time.Sleep(100 * time.Millisecond)
	}
}
