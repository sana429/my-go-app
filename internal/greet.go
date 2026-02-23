package internal

import "fmt"

func GreetUser() {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello, " name)
}