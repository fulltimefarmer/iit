package main

import (
	"fmt"
	"strconv"
)

func main() {
	// nest for loops
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println("Outer loop: " + strconv.Itoa(i))
			fmt.Println("Inner Loop: " + strconv.Itoa(j))
			fmt.Println("\n")
		}
	}

}
