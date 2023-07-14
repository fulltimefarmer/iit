// example reference: https://shaneutt.com/blog/golang-ca-and-signed-cert-go/
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(caPrivKey)

}
