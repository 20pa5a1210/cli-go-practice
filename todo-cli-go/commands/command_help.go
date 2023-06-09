package commands

import (
	"fmt"
)

func Help_command(cfg *TodoList, args ...string) error {
    fmt.Println("                                              |       WELOCOME TO TODO APP              |");
    fmt.Println("                                              |       --------------------              |");
	commands := GetCommands()
	fmt.Printf("%-10s %s\n", "Title", "Description")
	fmt.Printf("%-10s %s\n", "-----", "-----------")
	for _, curr := range commands {
		println(" ")
		// display title and description in table format
		// add the -- around them it like table
		// add headers to the table
		fmt.Printf("%-10s %s\n", curr.Title, curr.Description)

	}
	return nil
}
