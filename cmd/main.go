package main

import (
	"fmt"
	"github.com/hassanpdn/task-manager-go/internal/usecases"
	"github.com/hassanpdn/task-manager-go/internal/interfaces"
	"github.com/hassanpdn/task-manager-go/internal/infrastructure"
)

func main() {
	// get inputs from user
	email := interfaces.GetInput("Enter your email")
	password := interfaces.GetInput("Enter your password")

	// register user
	user := usecases.Register(email, password)

	// Save user
	infrastructure.SaveUser(user)

	// Show users list
	fmt.Println("User List:", userStorage)

	// create a new task
	name := interfaces.GetInput("Enter task name")
	dueDate := interfaces.GetInput("Enter task due date")
	category := interfaces.GetInput("Enter task category")
	usecases.CreateTask(name, dueDate, category)
}
