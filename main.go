package main

import (
	"fmt"
	"my-go-app/internal"
)

func main() {
	for {
		var choice int

		fmt.Println("1. Greet")
		fmt.Println("2. Show Date/Time")
		fmt.Println("3. To-Do List")
		fmt.Println("4. Exit")
		fmt.Println("Choose: ")

		fmt.Scanln(&choice)

		if choice == 1 {
			internal.GreetUser()
		} else if choice == 2 {
			internal.ShowDateTime()
		} else if choice == 3 {
			internal.TodoMenu()
		} else if choice == 4 {
			fmt.Println("Goodbye")
			return
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}
