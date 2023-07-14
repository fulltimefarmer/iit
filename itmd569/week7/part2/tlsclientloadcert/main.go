package main

import (
	"crypto/rand"
	"crypto/tls"
	"log"
)

func main() {
	// load client pem and key
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	// create a tls config to store the cert
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	// start a connection to the server
	conn, err := tls.Dial("tcp", "127.0.0.1:8000", &config)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
}
