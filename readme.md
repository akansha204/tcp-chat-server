# Go TCP Chat Server

A simple TCP-based chat server built in Go to understand concurrency concepts such as goroutines, channels, and mutex locks.

---

## Overview

This project demonstrates how multiple clients can connect to a single TCP server and exchange messages in real time.  
Whenever a client sends a message, the server receives it and broadcasts it to all connected clients.

---

## Key Concepts

- **Goroutines** – Lightweight threads in Go used to handle multiple clients concurrently.
- **Channels** – Used for communication between goroutines; here, messages are sent through a channel to be broadcast.
- **Mutex (sync.Mutex)** – Ensures that shared data structures like the `clients` map are accessed safely by one goroutine at a time.
- **TCP Connections** – Each client establishes a persistent TCP connection with the server.

---

## Folder Structure

tcp-chat-server/
│
├── go.mod
├── server/
│ ├── server.go # Starts the server, handles new clients, and broadcasts messages
│ └── client.go # Handles each client's messages and communication
└── readme.md

---

## How It Works

1. The server starts on a given port (e.g., `:8080`) and listens for new client connections.
2. When a client connects, it is added to a shared map (`clients`).
3. Each client runs in its own goroutine using `handleClient()`.
4. All messages are sent to a shared channel.
5. Another goroutine (`handleChannel()`) continuously reads messages from the channel and broadcasts them to all clients.
6. Mutex locks ensure that the shared `clients` map is accessed safely.

---

## How to Run

1. Run the server 
 ```go run server/server.go ```
2. Run in another terminal this command to connect clients using nc(netcat)
```nc localhost 8080```
