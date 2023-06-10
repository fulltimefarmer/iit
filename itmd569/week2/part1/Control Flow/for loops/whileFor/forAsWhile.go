package main

import (
	"fmt"
)

func main() {
	// go doesnt have while loops
	// but we can use for loops as a while loop

	//for as a while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		if i >= 5 {
			break //break statement is used to break the control if cond occurs
		}

		i++
	}

}
