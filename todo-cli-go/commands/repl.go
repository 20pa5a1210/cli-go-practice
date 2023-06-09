package commands

import (
	"strings"
)

type Todo struct {
	Task string
	Done bool
}

type TodoList struct {
	Tasks []Todo
}
type CliCommands struct {
	Title       string
	Description string
	CallBack    func(*TodoList, ...string) error
}

func GetCommands() map[string]CliCommands {
	return map[string]CliCommands{
		"help": {
			Title:       "help",
			Description: "Prints Help Menu",
			CallBack:    Help_command,
		},
		"exit": {
			Title:       "exit",
			Description: "Exits the program",
			CallBack:    Exit_command,
		},
		"add": {
			Title:       "add",
			Description: "Adds a task to the list",
			CallBack:    Add_command,
		},
		"list": {
			Title:       "list",
			Description: "Lists all tasks",
			CallBack:    List_command,
		},
		"done": {
			Title:       "done",
			Description: "Marks a task as done",
			CallBack:    Done_command,
		},
		"read": {
			Title:       "read",
			Description: "Reads a task from the list",
			CallBack:    Read_command,
		},
		"addfile": {
			Title:       "addfile",
			Description: "Adds a task to the list from file",
			CallBack:    Add_todos_to_file,
		},
	}

}

func Cleaned(str string) []string {
	lowStr := strings.ToLower(str)
	words := strings.Fields(lowStr)
	return words
}
