package commands

import (
	"fmt"
	"strings"
)

func Edit_Command(cfg *TodoList, args ...string) error {
	taskNum := int(args[0][0] - '0')
	if taskNum > len(cfg.Tasks) {
		fmt.Println("> Task number does not exist")
		return nil
	}
	// convert array of strings to a single string
	task := strings.Join(args[1:], " ")
	cfg.Tasks[taskNum].Task = task
	fmt.Printf("Edited task %d %s\n", taskNum, args[1])
	return nil
}
