package main

import (
	"fmt"
)

func main() {
	var arr0 [4]int // declare an array that holds 4 ints
	fmt.Println("Empty Array 0: ", arr0)

	arr0[3] = 469
	arr0[1] = 569
	fmt.Println("Array 0 with contents", arr0)

	arr1 := [4]int{6, 3, 1, 5} // declaring and initilaizing values at the same time
	fmt.Println("Array 1 contents", arr1)

	// two dimensional arrays
	var twoDimensional [3][2]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			twoDimensional[i][j] = (j - i) * 2
		}
	}

	fmt.Println("Two dimensional Array: ", twoDimensional)

}
