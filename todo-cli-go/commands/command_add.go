package commands

import (
	"fmt"
	"os"
	"strings"
)

func Add_command(cfg *TodoList, args ...string) error {
	if len(args) == 0 {
		return nil
	}
	// convert array of strings to a single string
	task := strings.Join(args, " ")
	cfg.Tasks = append(cfg.Tasks, Todo{Task: task, Done: false})
	file, err := os.ReadFile("todos.txt")
	if err != nil {
		return err
	}
	file = append(file, []byte("\n")...)
	file = append(file, []byte(task)...)
	err = os.WriteFile("todos.txt", file, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Added task: %s\n", task)

	return nil
}
