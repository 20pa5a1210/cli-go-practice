package commands

import (
	"fmt"
	"os"
)

func Read_command(cfg *TodoList, args ...string) error {
	// read tasks from file
	// read all tasks from ReadTasksFromFile()
	// write code
	// read exisiting tasks from ReadTasksFromFile()
	content, err := os.ReadFile("todos.txt")
	if err != nil {
		return err
	}
	for _, curr := range content {
		if curr == '\n' {
			fmt.Println()
		} else {
			fmt.Printf("%c", curr)
		}
	}
	return nil
}
