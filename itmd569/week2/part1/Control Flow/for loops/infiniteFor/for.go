package main

import (
	"fmt"
)

func main() {
	// for loops without a condition can be used to loop until a condition is met
	condition := 0
	for {
		fmt.Println(condition)
		if condition >= 10 {
			break
		}
		condition++
	}

	// this runs forever and is legal in go
	for {
		fmt.Println("Inside an unconditional for!!!!")
	}

}
