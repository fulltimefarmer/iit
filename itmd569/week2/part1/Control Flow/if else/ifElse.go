package main

import (
	"fmt"
)

func main() {

	if true {
		fmt.Println("You will see me!")
	}
	if false {
		fmt.Println("You will not see me :( ")
	}
	if !true {
		fmt.Println("You will see me!")
	}
	if !false {
		fmt.Println("You will not see me :( ")
	}

	// if else
	n := 45
	if n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	number := 51
	// shorthand
	if computer := "Dell"; number > 50 {
		fmt.Println(computer)
	}
}
