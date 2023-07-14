package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	c := &http.Client{
		Timeout: 15 * time.Second,
	}
	_, err := c.Get("https://jsonplaceholder.typicode.com/users/1")

	if err != nil {
		log.Fatal()
	}
}
