package commands

import "fmt"

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
