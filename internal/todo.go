package internal

import "fmt"

var tasks []string

func AddTask() {
	fmt.Println("Enter task: ")
	var task string
	fmt.Scanln(&task)
	tasks = append(tasks, task)
	fmt.Println("Task added.")
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet.")
		return
	}
	for i, t := range tasks {
		fmt.Printf("%d. %s\n", i+1, t)
	}
}

func RemoveTask() {
	if len(tasks) == 0 {
		fmt.Println("No tasks to remove.")
		return
	}
	ListTasks()
	fmt.Println("Enter task number to remove: ")
	var num int
	fmt.Scanln(&num)
	if num < 1 || num > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks = append(tasks[:num-1], tasks[num:]...)
	fmt.Println("Task removed.")
}

func TodoMenu() {
	for {
		var choice int
		fmt.Println("-- To-Do List --")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Back")
		fmt.Println("Choose: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			AddTask()
		} else if choice == 2 {
			ListTasks()
		} else if choice == 3 {
			RemoveTask()
		} else if choice == 4 {
			return
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}
