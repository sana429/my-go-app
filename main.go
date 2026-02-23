package main

import "fmt"

func main() {

	var choice int

	fmt.Println("1. Greet")
	fmt.Println("2. Exit")
	fmt.Println("Choose: ")

	fmt.Scanln(&choice)

	if choice == 1 {
		var name string
		fmt.Println("Enter your name: ")
		fmt.Scanln(&name)
		fmt.Println("Hello:", name)
	} else {
		fmt.Println("Goodbye")
	}
}
