package main

import (
	"fmt"
	"strings"
)

func main() {

	var bestLanguage string
	fmt.Println("Enter the best programming language")
	fmt.Scanln(&bestLanguage)
	strings.ToLower(bestLanguage)

	switch bestLanguage {
	case "go":
		fmt.Println("You are correct!")

	case "python":
		fmt.Println("pretty good, but not quite")
		break

	case "c":
		fmt.Println("You're a daring one, aren't you?")

	case "c++":
		fmt.Println("ugh")

	case "c#":
		fmt.Println("Basically microsoft java")

	case "java":
		fmt.Println("No.")

	case "swift":
		fmt.Println("good, but not quite")

	case "html":
		fmt.Println("Nice meme")

	default:
		fmt.Println("I dont know that language")
	}
}
