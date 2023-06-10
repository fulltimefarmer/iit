package main

import (
	"fmt"
)

func main() {

	fmt.Println("Chat Program")

	for {
		fmt.Println("How are you doing today?")
		fmt.Println("Type !options, to see how you can talk to me")
		fmt.Println("Enter your response: ")
		var response string
		fmt.Scanln(&response)

		switch response {

		case "Good":
			fmt.Println("That great! Me too!")
			continue

		case "Bad":
			fmt.Println("I'm sorry to hear that, maybe programming in Go will cheer you up!")
			continue

		case "Tired":
			fmt.Println("zzzzzz")
			continue

		case "Average":
			fmt.Println("Average is ok! I'm sure tomorrow will be great")
			continue

		case "!options":
			fmt.Println("You can respond with")
			fmt.Println("Good", "Bad", "Average", "Amazing")
			continue

		case "!quit":
			break

		default:
			fmt.Println("Hmm, I don't know what that means")
			continue
		}
		break

	}

}
