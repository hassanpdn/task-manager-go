package usecases

import "fmt"

func CreateTask(name, dueDate, category string) {
	fmt.Println(name, category, dueDate)
}

func CreateCategory(title, color string) {
	fmt.Println(title, color)
}
