package main

import "fmt"

func main() {

	var name string

	fmt.Println("Enter name: ")
	fmt.Scanln(&name)
	fmt.Println("You entered name: ", name)
}
