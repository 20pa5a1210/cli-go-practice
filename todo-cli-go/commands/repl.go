package commands

import (
	"strings"
)

type CliCommands struct {
	Title       string
	Description string
	CallBack    func(...string) error
}

func GetCommands() map[string]CliCommands {
	return map[string]CliCommands{
		"help": {
			Title:       "help",
			Description: "Prints Help Menu",
			CallBack:    Help_command,
		},
	}

}

func Cleaned(str string) []string {
	lowStr := strings.ToLower(str)
	words := strings.Fields(lowStr)
	return words
}
