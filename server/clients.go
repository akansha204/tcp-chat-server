package server

import (
	"bufio"
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer func() {
		mu.Lock()
		ClientId := clients[conn]
		delete(clients, conn) // safely remove client from map
		mu.Unlock()
		conn.Close()
		fmt.Println(ClientId, "got disconnected")
	}()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		msg := fmt.Sprintf("[%s]: %s", clients[conn], message)
		channel <- Message{From: conn, Text: msg}
	}
}
