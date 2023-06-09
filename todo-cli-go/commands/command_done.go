package commands

import "fmt"

func Done_command(cfg *TodoList, args ...string) error {
	if len(cfg.Tasks) == 0 {
		fmt.Println("> No tasks to mark done")
		return nil
	}
	if len(args) == 0 {
		fmt.Println("> Please specify a task number to mark done")
		return nil
	}
	taskNum := int(args[0][0] - '0')
	cfg.Tasks[taskNum].Done = true
	fmt.Printf("Marked task %d as done\n", taskNum)
	return nil
}
