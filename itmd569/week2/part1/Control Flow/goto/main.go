package main

import "fmt"

func main() {
	// This program prints out only the even numbers from 1-100

	var i int = 1

LOOPAGAIN:
	for i <= 100 {
		if i%2 != 0 {
			i++
			goto LOOPAGAIN // execution jumps back to the line with LOOPAGAIN does not print out the value
		}
		fmt.Printf("value of a: %d\n", i)
		i++
	}
}
