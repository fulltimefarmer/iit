package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	request := make([]byte, 100)
	_, err := conn.Read(request)
	if err != nil {
		fmt.Println("Error reading request:", err)
	}

	response := []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\n")
	response = append(response, request...)

	_, err = conn.Write(response)
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
	conn.Close()
}


func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection concurrently using a Goroutine
		go handleConnection(conn)
	}
}
