package main

import (
	"fmt"
)

func main() {
	users := [5]string{"user1", "user2", "user3", "user4", "user5"}

	fmt.Println("Range and Arrays")
	for index, user := range users {
		fmt.Println("index:", index)
		fmt.Println("users:", user)
	}
}
