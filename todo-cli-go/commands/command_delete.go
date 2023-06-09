package commands

import "fmt"

func Delete_command(todos *TodoList, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Please provide a task to delete")
		return nil
	}
	// convert taskNum to int
	taskNum := int(args[0][0] - '0')
	for i, task := range todos.Tasks {
		if taskNum == i {
			todos.Tasks = append(todos.Tasks[:i], todos.Tasks[i+1:]...)
			fmt.Println("Task deleted")
			fmt.Println(task)
			return nil
		}
	}
	return nil
}
