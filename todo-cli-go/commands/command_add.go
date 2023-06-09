package commands

import (
	"fmt"
	"strings"
)

func Add_command(cfg *TodoList, args ...string) error {
	if len(args) == 0 {
		return nil
	}
	// convert array of strings to a single string
	task := strings.Join(args, " ")
	cfg.Tasks = append(cfg.Tasks, Todo{Task: task, Done: false})
	fmt.Printf("Added task: %s\n", task)

	return nil
}

func List_command(cfg *TodoList, args ...string) error {
	if len(cfg.Tasks) == 0 {
		fmt.Println("> No tasks to list")
		return nil
	}
	fmt.Println("> Tasks:")
	fmt.Println("----------")
	//i make a table of Tasks include number of task and task and done or not
	// write code
    // task number , done or not , tasks
    fmt.Printf("%s %-10s %-20s\n", "No", "Task Status",          "Task")
    fmt.Printf("%s %-10s %-20s\n", "---", "------------", "------------------------------")


	for i, curr := range cfg.Tasks {
		DoneFlag := ""
		if curr.Done {
			DoneFlag = "Completed"
		} else {
			DoneFlag = "uncompleted"
		}
        // print task number and task and done or not
        // print in table format
        fmt.Printf("%d.  %-10s %-20s\n", i, DoneFlag, curr.Task)
	}
	return nil
}

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
