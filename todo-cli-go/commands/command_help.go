package commands

import (
	"fmt"
)

func Help_command(args ...string) error {
	fmt.Println("Welcome to todo cli application")
	fmt.Println("List of availble commands")

	commands := GetCommands()
	for _, curr := range commands {
		println(" ")
		fmt.Printf(" %s %s ", curr.Title, curr.Description)
	}
	return nil
}
