package main

import "fmt"

const pi float64 = 3.14159265

func main() {
	var circleArea float64
	radius := 2.5

	circleArea = pi * (radius * radius)
	fmt.Println(pi)
	fmt.Println(radius)
	fmt.Println(circleArea)

	// cannot modify pi since its constant!
	// pi = 3.15 throws compilier error

}
