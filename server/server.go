package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Message struct {
	From net.Conn
	Text string
}

var clients = make(map[net.Conn]string)

var channel = make(chan Message)

var mu sync.Mutex

func Start() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	fmt.Println("Chat server started on :8080")

	// This goroutine runs in background continuously reads messages from channel and sends to all clients
	go handleChannel()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}

		mu.Lock()
		id := len(clients) + 1
		label := fmt.Sprintf("client%d-> %s", id, conn.RemoteAddr().String())

		clients[conn] = label
		mu.Unlock()

		fmt.Println("New client connected:", clients[conn])

		go handleClient(conn)
	}
}

// responsible to listen on the channel and broadcast messages to all connected clients
func handleChannel() {
	for {
		msg := <-channel
		mu.Lock()
		for conn := range clients {
			if conn == msg.From {
				continue
			}
			fmt.Fprint(conn, msg.Text)
		}
		mu.Unlock()
	}
}
