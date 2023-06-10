// https://gobyexample.com/variables

package main

import "fmt"

var a = 2 // var can be used inside or outside of functions

func main() {

	fmt.Println(a)

	// var vs :=
	var b = 2 // declares and assigns b
	c := 2    // shorthand for declare c and assign it a value, can only be used inside of functions
	fmt.Println(b, c)

	// use var to declare variables without assigning a value
	var d int
	d = 2
	fmt.Println(d)

	// use var to declare multiple variables of the same type at once
	var e, f int
	e, f = 3, 4
	fmt.Println(e, f)

	// use var to declare then assign variables of the same type at once
	var g, h int = 1, 2
	fmt.Println(g, h)

	// can do the same thing with short declaration, even with mixed types
	i, j, k := 2, 4.0, "string!"
	fmt.Println(i, j, k)

}
