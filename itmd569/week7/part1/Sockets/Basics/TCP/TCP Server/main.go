package main

import (
	"fmt"

	"log"
	"net"
)

func main() {
	connection, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	for {
		conn, err := connection.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	fmt.Println(c)
	c.Close()
}
