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
    if  err != nil {
        return err
    }
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
	fmt.Printf("%s %-10s %-20s\n", "No", "Task Status", "Task")
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
