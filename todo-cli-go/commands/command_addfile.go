package commands

import "os"

func Add_todos_to_file(cfg *TodoList, args ...string) error {
	content, err := os.ReadFile("todos.txt")
	if err != nil {
		return err
	}
	// append new task in new line
	for _, curr := range cfg.Tasks {
		content = append(content, []byte("\n")...)
		content = append(content, []byte(curr.Task)...)
	}
	err = os.WriteFile("todos.txt", content, 0644)
	if err != nil {
		return err
	}
	return nil
}
