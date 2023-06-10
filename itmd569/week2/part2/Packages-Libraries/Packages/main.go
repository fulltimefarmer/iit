package main

import (
	"fmt"

	"github.com/tsmith-41/go-examples/Week2/Day2/Packages-Libraries/Packages/math"
)

func main() {

	num1 := 469.0
	num2 := 569.0

	result := math.Add(num1, num2)
	fmt.Println(result)
}

// How would we make num1 and num2 configurable by the user?
