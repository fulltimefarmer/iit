package main

import "fmt"

func main() {
	users := make([]string, 0)
	users = append(users, "user1", "user2", "user3", "user4", "user5", "user6")

	fmt.Println("Range Slices")
	for index, user := range users {
		fmt.Println("index:", index)
		fmt.Println("users:", user)
	}
}
