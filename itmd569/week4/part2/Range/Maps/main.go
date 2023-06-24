package main

import "fmt"

func main() {
	users := map[int]string{0: "user0", 1: "user1", 2: "user2", 3: "user3"}

	fmt.Println("Range and Maps")
	for key, value := range users {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}
}
