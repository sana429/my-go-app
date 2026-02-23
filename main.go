package main

import (
	"fmt"
	"my-go-app/internal"
)

func main() {

	var choice int

	fmt.Println("1. Greet")
	fmt.Println("2. Show Date/Time")
	fmt.Println("3. Exit")
	fmt.Println("Choose: ")

	fmt.Scanln(&choice)

	if choice == 1 {
		internal.GreetUser()
	} else if choice == 2 {
		internal.ShowDateTime()
	} else {
		fmt.Println("Goodbye")
	}
}
