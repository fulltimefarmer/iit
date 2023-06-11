package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println("How are you doing today?")
	fmt.Println("Tips: !commands for available commands")

	for {
		var input string
		fmt.Scanln(&input)

		switch strings.ToLower(input) {
			case "good":
				fmt.Println("Cool~")
			case "bad":
				fmt.Println("I'm sorry to hear that. :-(")
			case "tired":
				fmt.Println("You should get some rest.")
			case "average":
				fmt.Println("Cheer~")
			case "!commands":
				fmt.Println("Available commands:")
				fmt.Println("good")
				fmt.Println("bad")
				fmt.Println("tired")
				fmt.Println("average")
				fmt.Println("!commands")
				fmt.Println("!stop")
			case "!stop":
				fmt.Println("Bye!")
				os.Exit(1)
			default:
				fmt.Println("Sorry, I didn't understand what you said.")
				fmt.Println("Type !commands to see the available options.")
		}
	}

}