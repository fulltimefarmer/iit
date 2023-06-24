package main

import (
	"fmt"
)

func main() {

	slice0 := make([]int, 4)
	fmt.Println("Empty Slice0: ", slice0)

	slice0[0] = 6
	slice0[1] = 3
	slice0[2] = 1
	slice0[3] = 5

	fmt.Println("Slice0 with values: ", slice0)

	// appending values to a slice
	slice0 = append(slice0, 469)
	slice0 = append(slice0, 569)
	slice0 = append(slice0, 123, 456, 789)
	fmt.Println("Slice0 with more appended values: ", slice0)

	// making a copy of slice0 into a new slice
	slice0Copy := make([]int, len(slice0))
	copy(slice0Copy, slice0)
	fmt.Println("Copied Slice0: ", slice0Copy)

	// declaring new slices using the slice operator on existing slices
	partialSlice := slice0[:2]
	fmt.Println("Partial slice: ", partialSlice)

	// two dimensional slices
	twoDimensional := make([][]int, 4)
	for i := 0; i < 3; i++ {
		twoDimensional[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			twoDimensional[i][j] = (j - i) * 2
		}
		fmt.Println("Two Dimensional Slice: ", twoDimensional)
	}
}
