// TCP client example
// Can be used to visit a webservice though a network port
// A client can send a TCP request to the network port and get its response
// servers need to supply a service to that network port and give the client a response
// net TCPConn is what handles these client server interactions

// writes text to a remote address, and reads the response
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// use ResolveTCPAddr to create address to connect to
	raddr, err := net.ResolveTCPAddr("tcp", "remote address")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Use DialTCP to create a connection to the remote address
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		fmt.Println("failed to connect to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	text := "test"
	// send text to server
	_, err = conn.Write([]byte(text))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// read response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("failed reading response:", err)
		os.Exit(1)
	}
	fmt.Println(string(buf[:n]))
}
