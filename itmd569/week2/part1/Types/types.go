package main

import "fmt"

func main() {

	// basic common types
	// showcase default values for types

	var myString string
	fmt.Println(myString)
	myString = "veryCoolString"
	fmt.Println(myString)

	var myInt int
	fmt.Println(myInt)
	myInt = 10
	fmt.Println(myInt)

	var myFloat32 float32
	fmt.Println(myFloat32)
	myFloat32 = 10.00000001 // rounded to 10.0 since its 32 bit, cannot store this many digits
	fmt.Println(myFloat32)

	var myFloat64 float64 // takes up twice as much memory, but is more accurate
	fmt.Println(myFloat64)
	myFloat64 = 10.000000001 // will not be rounded since its 64 bit
	fmt.Println(myFloat64)

	var myBool bool
	fmt.Println(myBool)
	myBool = true
	fmt.Println(myBool)
}
