package main

import "fmt"

func main() {

	var firstName, lastName string

	fmt.Println("Enter first name: ")
	fmt.Scanln(&firstName)
	fmt.Println("Enter last name: ")
	fmt.Scanln(&lastName)
	fmt.Println("You entered name: ", firstName, lastName)
}
